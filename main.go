//go:build js && wasm

package main

import (
	"main/types"
	"net/http"
	"time"

	"codeberg.org/darckfast/workers-go/platform/cloudflare/fetch"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
)

func main() {
	mux := httprouter.New()
	mux.HandlerFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		easyjson.MarshalToWriter(types.Payload{
			Time: time.Now().UnixMilli(),
		}, w)
	})

	fetch.ServeNonBlock(mux)

	<-make(chan struct{}) // required
}
