package main

import (
	"context"
	"database/sql"
	"net/http"
)

type listCustomerInfoReq struct {
	LiveID int64 `json:"live_id"`
}

type listCustomerInfoResp struct {
	Error      *string `json:"error"`
	CustomerID []int64 `json:"customer_id"`
}

func listCustomerInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "list_customer_info", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "list_customer_info", "unauthorized operation")
		return
	}
	var req listCustomerInfoReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "list_customer_info", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		reportError(w, err, "list_customer_info", "database error")
		return
	}
	defer tx.Commit()
	rows, err := tx.Query(
		"SELECT CustomerID FROM liveViewer WHERE "+
			"UserID = ? AND "+
			"LiveID = ?;",
		userID, req.LiveID,
	)
	if err != nil {
		reportError(w, err, "list_customer_info", "database error")
		return
	}
	defer rows.Close()
	resp := listCustomerInfoResp{}
	for rows.Next() {
		var rec int64
		err = rows.Scan(&rec)
		if err != nil {
			reportError(w, err, "list_customer_info", "database error")
			return
		}
		resp.CustomerID = append(resp.CustomerID, rec)
	}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "list_customer_info", "failed to generate response")
		return
	}
}
