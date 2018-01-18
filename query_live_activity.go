package main

import (
	"context"
	"database/sql"
	"net/http"
)

type queryLiveActivityReq struct {
	liveActivity
}

type queryLiveActivityResp struct {
	Error        *string        `json:"error"`
	LiveActivity []liveActivity `json:"live_activity"`
}

func queryLiveActivityHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "query_live_activity", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "query_live_activity", "unauthorized operation")
		return
	}
	var req queryLiveActivityReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "query_live_activity", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		reportError(w, err, "query_live_activity", "database error")
		return
	}
	defer tx.Commit()
	rows, err := tx.Query(
		"SELECT * FROM liveActivity INNER JOIN customerInfo ON "+
			"customerInfo.ID = liveActivity.CustomerID AND "+
			"customerInfo.UserID = liveActivity.UserID WHERE "+
			"IFNULL(liveActivity.ID = ?, 1) AND "+
			"liveActivity.UserID = ? AND "+
			"IFNULL(liveActivity.LiveID = ?, 1) AND "+
			"IFNULL(liveActivity.Time = ?, 1) AND "+
			"IFNULL(liveActivity.CustomerID = ?, 1) AND "+
			"IFNULL(liveActivity.Activity = ?, 1);",
		req.ID, userID, req.LiveID, req.Time, req.CustomerID, req.Activity,
	)
	if err != nil {
		reportError(w, err, "query_live_activity", "database error")
		return
	}
	defer rows.Close()
	resp := queryLiveActivityResp{}
	for rows.Next() {
		var rec liveActivity
		var tags *string
		rec.CustomerInfo = new(customerInfo)
		err = rows.Scan(&rec.ID, &rec.UserID, &rec.LiveID, &rec.Time, &rec.CustomerID, &rec.Activity,
			&rec.CustomerInfo.ID, &rec.CustomerInfo.UserID, &rec.CustomerInfo.CustomerName, &rec.CustomerInfo.Mobile, &rec.CustomerInfo.Status, &tags)
		if err != nil {
			reportError(w, err, "query_live_activity", "database error")
			return
		}
		rec.CustomerInfo.Tags = decodeList(tags)
		resp.LiveActivity = append(resp.LiveActivity, rec)
	}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "query_live_activity", "failed to generate response")
		return
	}
}
