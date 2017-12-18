package main

import (
	"net/http"
)

type newLiveSessionResp struct {
	Error *string `json:"error"`
	ID    uint64  `json:"id"`
}

func newLiveSessionHandler(w http.ResponseWriter, req *http.Request) {
}
