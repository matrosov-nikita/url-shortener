// Package main provides implementation for CLI client.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	encode := flag.String("encode", "", "URL to encode")
	decode := flag.String("decode", "", "URL to decode")
	apiURL := flag.String("apiURL", "http://localhost:8080", "API URL in host:port format ")
	flag.Parse()
	bothEmpty := len(*encode) == 0 && len(*decode) == 0
	bothFull := len(*encode) > 0 && len(*decode) > 0
	if bothEmpty || bothFull {
		log.Fatalf("encode or decode should be specified")
	}

	var val string
	handlerPath := *apiURL + "/urlshortener/URLHandler/"

	if len(*encode) > 0 {
		handlerPath += "Encode"
		val = *encode
	} else {
		handlerPath += "Decode"
		val = *decode
	}

	res, err := request(handlerPath, val)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	io.Copy(os.Stdout, res)
}

func request(URL, payload string) (io.ReadCloser, error) {
	data, err := json.Marshal(map[string]string{"url": payload})
	if err != nil {
		return nil, fmt.Errorf("could not marshal url: %v", err)
	}

	res, err := http.Post(URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("could not send post request: %v", err)
	}

	return res.Body, nil
}
