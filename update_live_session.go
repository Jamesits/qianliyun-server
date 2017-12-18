package main

import (
	"net/http"
)

type updateLiveSessionReq struct {
	ID      uint64   `json:"id"`
	URL     *string  `json:"url"`
	Title   *string  `json:"title"`
	Host    *string  `json:"host"`
	Comment *string  `json:"comment"`
	Begin   *float64 `json:"begin"`
	End     *float64 `json:"end"`
}

type updateLiveSessionResp struct {
	Error   *string `json:"error"`
	ID      uint64  `json:"id"`
	URL     string  `json:"url"`
	Title   string  `json:"title"`
	Host    string  `json:"host"`
	Comment string  `json:"comment"`
	Begin   float64 `json:"begin"`
	End     float64 `json:"end"`
}

func updateLiveSessionHandler(w http.ResponseWriter, req *http.Request) {
}
