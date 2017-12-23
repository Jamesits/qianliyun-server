package main

import (
	"crypto/rand"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func main() {
	var (
		authkey [32]byte
		enckey  [16]byte
	)
	_, _ = rand.Read(authkey[:])
	_, _ = rand.Read(enckey[:])
	store = sessions.NewCookieStore(authkey[:], enckey[:])
	store.Options.Path = "/api"
	store.Options.HttpOnly = true
	initDatabase()
	servemux := mux.NewRouter()
	servemux.HandleFunc("/api/login.php", loginHandler)
	servemux.HandleFunc("/api/query_user_info.php", queryUserInfoHandler)
	servemux.HandleFunc("/api/update_live_session.php", updateLiveSessionHandler)
	servemux.HandleFunc("/api/query_live_session.php", queryLiveSessionHandler)
	servemux.HandleFunc("/api/update_customer_info.php", updateCustomerInfoHandler)
	servemux.HandleFunc("/api/query_customer_info.php", queryCustomerInfoHandler)
	servemux.HandleFunc("/api/list_customer_info.php", listCustomerInfoHandler)
	log.Fatalln(http.ListenAndServe("[::1]:6452", handlers.CombinedLoggingHandler(os.Stdout, servemux)))
}
