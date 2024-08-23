package api

import (
	"day03/ex01/store"
	"day03/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
)

type site struct {
	Total   int           `json:"total"`
	Places  []types.Place `json:"places"`
	Name    string        `json:"name"`
	Next    int           `json:"next"`
	Previos int           `json:"previos"`
	Last    int           `json:"last"`
}
type Err struct {
	Err string `json:"error"`
}

type API types.Place

func (a API) ApiShow(es1 *elasticsearch.Client, totalHits int) {
	var store *store.ElasticsearchStore
	store = store.NewElasticsearchStore(es1, "places")

	var amountOfPages int
	pageSize := 10
	if (totalHits/pageSize) % 10 == 0 {
		amountOfPages = totalHits / pageSize
	} else {
		amountOfPages = totalHits/pageSize + 1
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageString := r.URL.Query().Get("page")
		pageNum, err := strconv.Atoi(pageString)
		if err != nil {
			return

		}
		if pageNum > amountOfPages {
			err_str := Err{
				Err: fmt.Sprintf("Invalid 'page' value: '%s'", pageString),
			}
			fmt.Println(err_str.Err)
			errJson, err := json.MarshalIndent(err_str, "", "\t")
			if err != nil {
				http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			w.Write(errJson)

		} else {
			var prev_page, next_page int
			if pageNum < amountOfPages {
				next_page = pageNum + 1
			}
			if pageNum > 1 {
				prev_page = pageNum - 1
			}
			places, _, err := store.GetPlaces(pageSize, (pageNum-1)*10)
			if err != nil {
				http.Error(w, "smt", http.StatusInternalServerError)
				return
			}
			data := site{
				Name:    "places",
				Total:   totalHits,
				Places:  places,
				Previos: prev_page,
				Next:    next_page,
				Last:    amountOfPages,
			}
			w.Header().Set("Content-Type", "application/json")
			dataJson, err := json.MarshalIndent(data, "", "\t")
			if err != nil {
				http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
				return
			}
			w.Write(dataJson)

		}

	})
	log.Println(http.ListenAndServe(":8888", nil))
}
