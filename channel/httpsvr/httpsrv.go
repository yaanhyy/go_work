package httpsvr

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)


//basicAuth
func basicAuth(h httprouter.Handle, requiredUser string, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		if requiredUser == "" || requiredPassword == "" {
			h(w, r, ps)
			return
		}
		user, password, hasAuth := r.BasicAuth()
		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// StartHTTP 在一个新的goroutine里面启动一个HTTPserver
func StartHTTP(endpoint string, user string, pwd string, allowedOrigins []string) {

	c := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
		MaxAge:         600,
		AllowedHeaders: []string{"*"},
	})

	router := httprouter.New()
	router.GET("/", basicAuth(Root,user, pwd))
	go http.ListenAndServe(endpoint, c.Handler(router))// c.Handler(router)
}