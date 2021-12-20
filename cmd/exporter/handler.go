package exporter

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Stp Exporter</h1>")
}
