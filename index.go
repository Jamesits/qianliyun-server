package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Apache/2.2.34'<!--")
	w.Header().Set("X-AspNet-Version", "2.0.50727'<!--")
	w.Header().Set("X-Powered-By", "PHP/5.5.38'<!--")
}
