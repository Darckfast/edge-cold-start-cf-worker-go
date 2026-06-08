//go:build js && wasm

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"codeberg.org/darckfast/workers-go/platform/cloudflare/fetch"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]int64{
			"time": time.Now().UnixMilli(),
		})
	})

	fetch.ServeNonBlock(mux)

	<-make(chan struct{}) // required
}
