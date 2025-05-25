package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (*Handler) Health(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "ok")
}
