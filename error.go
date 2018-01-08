package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type errorResp struct {
	Error string `json:"error"`
}

func reportError(w http.ResponseWriter, err error, module, message string) {
	debug.PrintStack()
	if err != nil {
		log.Printf("[%s] %s\n", module, err.Error())
	} else {
		log.Printf("[%s] %s\n", module, message)
	}
	msg := errorResp{message}
	buf, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Apache/2.2.34'<!--")
	w.Header().Set("X-AspNet-Version", "2.0.50727'<!--")
	w.Header().Set("X-Powered-By", "PHP/5.5.38'<!--")
	w.WriteHeader(503)
	w.Write(buf)
}

func reportInvalidArgument(w http.ResponseWriter, module, arg string) {
	reportError(w, nil, module, fmt.Sprintf("Invalid argument: %q", arg))
}
