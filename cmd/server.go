package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/bwesterb/go-pow"
)

func main() {
	quotes := [][]byte{
		[]byte("Want to be healthy? Don't get sick. (c) Jason Stathem"),
		[]byte("The secret of success is simple, but it's a secret. (c) Jason Stathem"),
	}

	http.HandleFunc("/wisdom", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.Header().Set("Allow", http.MethodGet)
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Printf("DEBUG - wisdom: request Host %s\n", request.RemoteAddr)

		req := pow.NewRequest(5, []byte("wisdom for"+request.RemoteAddr))
		fmt.Printf("DEBUG - wisdom: request: %s\n", req)
		_, err := writer.Write([]byte(req))
		if err != nil {
			fmt.Println("can not write /wisdom response", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/proof", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.Header().Set("Allow", http.MethodPost)
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Printf("DEBUG - proof: request URI %s\n", request.RequestURI)

		proof, err := io.ReadAll(request.Body)
		if err != nil && err != io.EOF {
			fmt.Println("can not read /proof body", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		req := pow.NewRequest(5, []byte("wisdom for"+request.RemoteAddr))
		ok, err := pow.Check(req, string(proof), []byte("wisdom"))
		if err != nil {
			fmt.Println("error check /proof", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		if ok {
			_, err := writer.Write(quotes[rand.Int()%2])
			if err != nil {
				fmt.Println("can not write quote", err)
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
