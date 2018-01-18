package main

import (
	"net/http"
)

type checkUpdateReq struct {
}

type checkUpdateResp struct {
	Error      *string    `json:"error"`
	UpdateInfo updateInfo `json:"update_info"`
}

func checkUpdateHandler(w http.ResponseWriter, r *http.Request) {
	resp := checkUpdateResp{
		Error: nil,
		UpdateInfo: updateInfo{
			LatestVersion: [4]uint{1, 1, 1708, 16290},
			UpdateURL:     "https://localhost/qianliyun-v1.0.0.0.msi",
			Changelog:     "Bug fixes and performance improvements.",
		},
	}
	err := encodeResponse(w, resp)
	if err != nil {
		reportError(w, err, "login", "failed to generate response")
		return
	}
}
