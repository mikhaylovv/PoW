package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bwesterb/go-pow"
)

func main() {
	getResp, err := http.Get("http://localhost:8080/wisdom")
	if err != nil {
		fmt.Println("can not send get request to server", err)
		return
	}
	if getResp.StatusCode != http.StatusOK {
		fmt.Println("bad get response status", getResp.StatusCode)
		return
	}

	getBody, err := io.ReadAll(getResp.Body)
	_, err = getResp.Body.Read(getBody)
	if err != nil && err != io.EOF {
		fmt.Println("can not read get response body", err)
		return
	}
	bodyString := string(getBody)
	fmt.Println("DEBUG: proof ", bodyString)

	proof, err := pow.Fulfil(bodyString, []byte("wisdom"))
	if err != nil {
		fmt.Println("can not fulfil proof of work", err)
		return
	}

	proofResp, err := http.Post("http://localhost:8080/proof", "string", strings.NewReader(proof))
	if err != nil {
		fmt.Println("can not post proof", err)
		return
	}
	if proofResp.StatusCode != http.StatusOK {
		fmt.Println("bad post response status", getResp.StatusCode)
		return
	}
	proofBody, err := io.ReadAll(proofResp.Body)
	if err != nil && err != io.EOF {
		fmt.Println("can not read proof request body", err)
		return
	}

	fmt.Println(string(proofBody))
}
