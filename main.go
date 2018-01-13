package main

import (
	"crypto/rand"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/natefinch/lumberjack.v2"
)

var store *sessions.CookieStore

func main() {
	err := os.MkdirAll("./logs", 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(os.Stderr, &lumberjack.Logger{
		Filename:   "./logs/qianliyun.log",
		MaxSize:    64,
		MaxBackups: 52,
		MaxAge:     7,
		Compress:   true,
	}))

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
	servemux.HandleFunc("/", indexHandler)
	servemux.HandleFunc("/api/check_update.php", checkUpdateHandler)
	servemux.HandleFunc("/api/login.php", loginHandler)
	servemux.HandleFunc("/api/query_customer_info.php", queryCustomerInfoHandler)
	servemux.HandleFunc("/api/query_live_activity.php", queryLiveActivityHandler)
	servemux.HandleFunc("/api/query_live_session.php", queryLiveSessionHandler)
	servemux.HandleFunc("/api/query_user_info.php", queryUserInfoHandler)
	servemux.HandleFunc("/api/telemetry.php", telemetryHandler)
	servemux.HandleFunc("/api/update_customer_info.php", updateCustomerInfoHandler)
	servemux.HandleFunc("/api/update_live_activity.php", updateLiveActivityHandler)
	servemux.HandleFunc("/api/update_live_session.php", updateLiveSessionHandler)
	servemux.HandleFunc("/cgi-bin/luci", luciHandler)
	servemux.PathPrefix("/cgi-bin/luci/").HandlerFunc(luciHandler)
	servemux.PathPrefix("/luci-static/").Handler(http.FileServer(http.Dir("./static")))
	log.Fatalln(http.ListenAndServe("[::1]:6452", handlers.CombinedLoggingHandler(io.MultiWriter(os.Stderr, &lumberjack.Logger{
		Filename:   "./logs/access.log",
		MaxSize:    64,
		MaxBackups: 52,
		MaxAge:     7,
		Compress:   true,
	}), servemux)))
}
