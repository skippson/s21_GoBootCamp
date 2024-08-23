package main

/*
#include "../cow/cow.c"
*/
import "C"

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lizrice/secure-connections/utils"
)

type OrderJSON struct {
	Money      int64  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int64  `json:"candyCount"`
}

type answerJSON struct {
	Change int64  `json:"change,omitempty"`
	Thanks string `json:"thanks,omitempty"`
}

func main() {
	kType := flag.String("k", "", "candyType")
	cCount := flag.Int64("c", 0, "candyCount")
	mMoney := flag.Int64("m", 0, "money")
	flag.Parse()

	order := OrderJSON{
		CandyType:  *kType,
		Money:      *mMoney,
		CandyCount: *cCount,
	}

	var data bytes.Buffer
	err := json.NewEncoder(&data).Encode(order)
	if err != nil {
		log.Fatal(err)
	}
	client := makeClient()
	req, err := client.Post("https://127.0.0.1:3333/v1/buy_candy", "application/json", bytes.NewBuffer(data.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	var cb answerJSON
	if req.StatusCode == 201 {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &cb)
		if err != nil {
			log.Fatal(err)
		}

		cb.Thanks = C.GoString(C.ask_cow(C.CString("Thank you!")))
		fmt.Printf("Change:%d\n%s",cb.Change, cb.Thanks)
		defer req.Body.Close()
	} else {
		_, err = io.Copy(os.Stdout, req.Body)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func makeClient() *http.Client {
	data, err := os.ReadFile("ca/minica.pem")
	if err != nil {
		log.Println(err)
	}

	cert, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}
	cert.AppendCertsFromPEM(data)

	config := &tls.Config{
		InsecureSkipVerify:    true,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		RootCAs:               cert,
		GetCertificate:        utils.CertReqFunc("ca/candy.tld/cert.pem", "ca/candy.tld/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	return client
}
