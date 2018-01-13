package main

import (
	"context"
	"net/http"
)

type updateLiveActivityReq struct {
	liveActivity
}

type updateLiveActivityResp struct {
	Error *string `json:"error"`
	ID    int64   `json:"id"`
}

func updateLiveActivityHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "update_live_activity", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "update_live_activity", "unauthorized operation")
		return
	}
	var req updateLiveActivityReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "update_live_activity", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		reportError(w, err, "update_live_activity", "database error")
		return
	}
	var resp updateLiveActivityResp
	if req.ID != nil {
		res, err := tx.Exec(
			"UPDATE LiveActivity SET "+
				"LiveID = IFNULL(?, LiveID), "+
				"Time = IFNULL(?, Time), "+
				"CustomerID = IFNULL(?, CustomerID), "+
				"Activity = IFNULL(?, Activity) "+
				"WHERE ID = ? AND UserID = ?;",
			req.LiveID, req.Time, req.CustomerID, req.Activity, req.ID, userID,
		)
		if err != nil {
			reportError(w, err, "update_live_activity", "database error")
			return
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			reportError(w, err, "update_live_activity", "database error")
			return
		}
		if rowsAffected == 0 {
			reportError(w, nil, "update_live_activity", "no record found")
			return
		}
		resp.ID = *req.ID
	} else {
		res, err := tx.Exec(
			"INSERT INTO LiveActivity "+
				"(UserID, LiveID, Time, CustomerID, Activity)"+
				"VALUES (?, ?, ?, ?, ?);",
			userID, req.LiveID, req.Time, req.CustomerID, req.Activity,
		)
		if err != nil {
			reportError(w, err, "update_live_activity", "database error")
			return
		}
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			reportError(w, err, "update_live_activity", "database error")
			return
		}
		resp.ID = lastInsertID
	}
	tx.Commit()
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "update_live_activity", "failed to generate response")
		return
	}
}
