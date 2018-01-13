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
		"SELECT * FROM liveActivity WHERE "+
			"IFNULL(ID = ?, 1) AND "+
			"UserID = ? AND "+
			"IFNULL(LiveID = ?, 1) AND "+
			"IFNULL(Time = ?, 1) AND "+
			"IFNULL(CustomerID = ?, 1) AND "+
			"IFNULL(Activity = ?, 1);",
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
		err = rows.Scan(&rec.ID, &rec.UserID, &rec.LiveID, &rec.Time, &rec.CustomerID, &rec.Activity)
		if err != nil {
			reportError(w, err, "query_live_activity", "database error")
			return
		}
		resp.LiveActivity = append(resp.LiveActivity, rec)
	}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "query_live_activity", "failed to generate response")
		return
	}
}
