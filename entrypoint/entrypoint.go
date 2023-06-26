package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var uri = flag.String("uri", "", "The URI for the payload generation server.")

func main() {
	flag.Parse()

	log.Printf("Program run with flag --uri=%s", *uri)

	if kubeHost := os.Getenv("KUBERNETES_SERVICE_HOST"); kubeHost == "" {
		log.Printf("☸ This binary may not be running on Kubernetes! WE LOVE KUBERNETES! ☸")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get(*uri)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong. Womp womp.\n")
			return
		}
		if res.StatusCode >= 300 {
			log.Println(res.Status)
			w.WriteHeader(res.StatusCode)
			fmt.Fprintf(w, "Something went wrong. Womp womp.\n")
			return
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong. Womp womp.\n")
			return
		}
		log.Println(string(body))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s generated this payload: %s\n", *uri, string(body))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
