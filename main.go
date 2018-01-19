package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", noDirListing(http.FileServer(http.Dir("public"))))
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8181", n)
}

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
