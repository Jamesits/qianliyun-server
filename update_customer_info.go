package main

import (
	"net/http"
)

type updateCustomerInfoReq struct {
	ID       uint64    `json:"id"`
	Username *string   `json:"username"`
	Mobile   *string   `json:"mobile"`
	Status   *string   `json:"status"`
	Tags     *[]string `json:"tags"`
}

type updateCustomerInfoResp struct {
	Error    *string  `json:"error"`
	ID       uint64   `json:"id"`
	Username string   `json:"username"`
	Mobile   string   `json:"mobile"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
}

func updateCustomerInfoHandler(w http.ResponseWriter, req *http.Request) {
}
