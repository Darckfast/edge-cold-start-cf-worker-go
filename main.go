//go:build js && wasm

package main

import (
	"net/http"
	"strconv"
	"time"

	"codeberg.org/darckfast/workers-go/platform/cloudflare/fetch"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.HandlerFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"time":` + strconv.FormatInt(time.Now().UnixMicro(), 10) + `}`))
	})

	fetch.ServeNonBlock(mux)

	<-make(chan struct{}) // required
}
