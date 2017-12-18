package main

import (
	"net/http"
)

type loginReq struct {
	Username   *string `json:"username"`
	Password   *string `json:"password"`
	MachineKey *string `json:"machine_key"`
}

type loginResp struct {
	Error *string `json:"error"`
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
}
