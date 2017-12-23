package main

import (
	"net/http"
)

type loginReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	MachineKey string `json:"machine_key"`
}

type loginResp struct {
	Error *string `json:"error"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	err := decodeRequest(r, &req)
	if err != nil {
		reportError(w, err, "login", "invalid request")
		return
	}
	err = setUserID(w, r, "root", 1)
	if err != nil {
		reportError(w, err, "login", "failed to log in")
		return
	}
	resp := loginResp{Error: nil}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, err, "login", "failed to generate response")
		return
	}
}
