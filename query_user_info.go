package main

import (
	"net/http"
)

type queryUserInfoResp struct {
	Error         *string `json:"error"`
	ID            uint64  `json:"id"`
	Username      string  `json:"username"`
	Alias         string  `json:"alias"`
	ResellerAlias string  `json:"reseller_alias"`
	AuthMax       string  `json:"auth_max"`
	AuthLeft      string  `json:"auth_left"`
	DeauthLeft    string  `json:"deauth_left"`
	Reseller      uint64  `json:"reseller"`
}

func queryUserInfoHandler(w http.ResponseWriter, req *http.Request) {
}
