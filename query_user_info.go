package main

import (
	"context"
	"database/sql"
	"net/http"
)

type queryUserInfoReq struct {
}

type queryUserInfoResp struct {
	Error    *string  `json:"error"`
	UserInfo userInfo `json:"user_info"`
}

func queryUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "query_user_info", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "query_user_info", "unauthorized operation")
		return
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		reportError(w, err, "query_user_info", "database error")
		return
	}
	defer tx.Commit()
	resp := queryUserInfoResp{}
	var rec userInfo
	rec.ResellerInfo = new(resellerInfo)
	var password, salt *string
	err = tx.QueryRow(
		"SELECT * FROM userInfo INNER JOIN resellerInfo ON resellerInfo.ID = userInfo.ResellerID WHERE userInfo.ID = ?;",
		userID,
	).Scan(&rec.ID, &rec.Username, &password, &salt, &rec.Alias, &rec.ResellerAlias, &rec.AuthMax, &rec.AuthLeft, &rec.DeauthLeft, &rec.ResellerID,
		&rec.ResellerInfo.ID, &rec.ResellerInfo.Alias, &rec.ResellerInfo.AppTitle, &rec.ResellerInfo.AppStatus, &rec.ResellerInfo.AppCopyright)
	if err != nil {
		reportError(w, err, "query_user_info", "database error")
		return
	}
	resp.UserInfo = rec
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "query_user_info", "failed to generate response")
		return
	}
}
