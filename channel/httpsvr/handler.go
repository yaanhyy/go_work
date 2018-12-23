package httpsvr

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "just a test")
}

