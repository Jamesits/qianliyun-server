package main

import (
	"net/http"
)

type queryLiveSessionReq struct {
	ID uint64 `json:"id"`
}

type queryLiveSessionResp struct {
	Error   *string `json:"error"`
	ID      uint64  `json:"id"`
	URL     string  `json:"url"`
	Title   string  `json:"title"`
	Host    string  `json:"host"`
	Comment string  `json:"comment"`
	Begin   float64 `json:"begin"`
	End     float64 `json:"end"`
}

func queryLiveSessionHandler(w http.ResponseWriter, req *http.Request) {
}
