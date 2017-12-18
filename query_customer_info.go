package main

import (
	"net/http"
)

type queryCustomerInfoReq struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type queryCustomerInfoResp struct {
	Error    *string  `json:"error"`
	ID       uint64   `json:"id"`
	Username string   `json:"username"`
	Mobile   string   `json:"mobile"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
}

func queryCustomerInfoHandler(w http.ResponseWriter, req *http.Request) {
}
