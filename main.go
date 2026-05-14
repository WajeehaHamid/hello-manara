package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const greeting = "v2: real GitOps round-trip works — edited locally, pushed, built by Actions, synced by Argo CD"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	startedAt := time.Now().UTC().Format(time.RFC3339)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w,
			"%s\n"+
				"pod:        %s\n"+
				"started_at: %s\n"+
				"served_at:  %s\n"+
				"path:       %s\n",
			greeting, host, startedAt, time.Now().UTC().Format(time.RFC3339), r.URL.Path,
		)
	})

	log.Printf("hello-manara listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
