package base

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"day03/types"

	"github.com/cenkalti/backoff/v4"
	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type Data types.Place

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
	es         *elasticsearch.Client
)

func init() {
	flag.StringVar(&indexName, "index", "places", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 16358, "Number of documents to generate")
	flag.Parse()
}

func (db Data) CreateESC(places []*types.Place) {

	var err error
	log.Printf(
		"\x1b[1mBulkIndexer\x1b[0m: documents [%s] workers [%d] flush [%s]",
		humanize.Comma(int64(numItems)), numWorkers, humanize.Bytes(uint64(flushBytes)))
	log.Println(strings.Repeat("▁", 65))
	retryBackoff := backoff.NewExponentialBackOff()

	es, err = elasticsearch.NewClient(elasticsearch.Config{
		RetryOnStatus: []int{502, 503, 504, 429},

		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},

		MaxRetries: 5,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	bi := db.CreateBI(es, indexName, numWorkers, flushBytes, places)
	db.CreateIndex(&bi, places, es, indexName)
	db.EncodePlaces(&bi, places)
	log.Printf("→ Read %s place", humanize.Comma(int64(len(places))))

}

func (db Data) CreateBI(es *elasticsearch.Client, indexName string, numWorkers int, flushBytes int, places []*types.Place) esutil.BulkIndexer {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,        // The default index name
		Client:        es,               // The Elasticsearch client
		NumWorkers:    numWorkers,       // The number of worker goroutines
		FlushBytes:    int(flushBytes),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	return bi
}

func (db Data) CreateIndex(bi *esutil.BulkIndexer, places []*types.Place, es *elasticsearch.Client, indexName string) {
	var (
		res *esapi.Response
		err error
	)

	if res, err = es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}

	res.Body.Close()
	res, err = es.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}
	res.Body.Close()

}

func (db Data) EncodePlaces(bi *esutil.BulkIndexer, places []*types.Place) {
	var countSuccessful uint64
	start := time.Now().UTC()

	for _, a := range places {
		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode place %d: %s", a.ID, err)
		}

		err = (*bi).Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",

				DocumentID: strconv.Itoa(a.ID),

				Body: bytes.NewReader(data),

				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}

	if err := (*bi).Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
	biStats := (*bi).Stats()

	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)
	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}

}