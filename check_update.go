package main

import (
	"net/http"
)

type checkUpdateReq struct {
}

type checkUpdateResp struct {
	Error         *string `json:"error"`
	LatestVersion [4]uint `json:"latest_version"`
	UpdateURL     string  `json:"update_url"`
}

func checkUpdateHandler(w http.ResponseWriter, r *http.Request) {
	resp := checkUpdateResp{
		Error:         nil,
		LatestVersion: [4]uint{1, 0, 0, 0},
		UpdateURL:     "https://localhost/qianliyun-v1.0.0.0.msi",
	}
	err := encodeResponse(w, resp)
	if err != nil {
		reportError(w, err, "login", "failed to generate response")
		return
	}
}
