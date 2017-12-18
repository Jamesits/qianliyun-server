package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResp struct {
	Error string `json:"error"`
}

func reportError(w http.ResponseWriter, err string) {
	msg := errorResp{err}
	buf, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(503)
	w.Write(buf)
}

func reportInvalidArgument(w http.ResponseWriter, arg string) {
	reportError(w, fmt.Sprintf("Invalid argument: %q", arg))
}
