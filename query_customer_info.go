package main

import (
	"context"
	"database/sql"
	"net/http"
)

type queryCustomerInfoReq struct {
	customerInfo
}

type queryCustomerInfoResp struct {
	Error        *string        `json:"error"`
	CustomerInfo []customerInfo `json:"customer_info"`
}

func queryCustomerInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "query_customer_info", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "query_customer_info", "unauthorized operation")
		return
	}
	var req queryCustomerInfoReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "query_customer_info", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		reportError(w, err, "query_customer_info", "database error")
		return
	}
	defer tx.Commit()
	rows, err := tx.Query(
		"SELECT * FROM customerInfo WHERE "+
			"IFNULL(ID = ?, 1) AND "+
			"UserID = ? AND "+
			"IFNULL(CustomerName = ?, 1) AND "+
			"IFNULL(Mobile = ?, 1) AND "+
			"IFNULL(Status = ?, 1) AND "+
			"IFNULL(Tags = ?, 1);",
		req.ID, userID, req.CustomerName, req.Mobile, req.Status, encodeList(req.Tags),
	)
	if err != nil {
		reportError(w, err, "query_customer_info", "database error")
		return
	}
	defer rows.Close()
	resp := queryCustomerInfoResp{}
	for rows.Next() {
		var rec customerInfo
		var tags *string
		err = rows.Scan(&rec.ID, &rec.UserID, &rec.CustomerName, &rec.Mobile, &rec.Status, &tags)
		if err != nil {
			reportError(w, err, "query_customer_info", "database error")
			return
		}
		rec.Tags = decodeList(tags)
		resp.CustomerInfo = append(resp.CustomerInfo, rec)
	}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "query_customer_info", "failed to generate response")
		return
	}
}
