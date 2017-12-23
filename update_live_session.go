package main

import (
	"context"
	"net/http"
)

type updateLiveSessionReq struct {
	liveSession
}

type updateLiveSessionResp struct {
	Error *string `json:"error"`
	ID    int64   `json:"id"`
}

func updateLiveSessionHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "update_live_session", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "update_live_session", "unauthorized operation")
		return
	}
	var req updateLiveSessionReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "update_live_session", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		reportError(w, err, "update_live_session", "database error")
		return
	}
	var resp updateLiveSessionResp
	if req.ID != nil {
		res, err := tx.Exec(
			"UPDATE liveSession SET "+
				"URL = IFNULL(?, URL), "+
				"Title = IFNULL(?, Title), "+
				"Host = IFNULL(?, Host), "+
				"Comment = IFNULL(?, Comment), "+
				"Begin = IFNULL(?, Begin), "+
				"End = IFNULL(?, End), "+
				"Tags = IFNULL(?, Tags) "+
				"WHERE ID = ? AND UserID = ?;",
			req.URL, req.Title, req.Host, req.Comment, req.Begin, req.End, encodeList(req.Tags), req.ID, userID,
		)
		if err != nil {
			reportError(w, err, "update_live_session", "database error")
			return
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			reportError(w, err, "update_live_session", "database error")
			return
		}
		if rowsAffected == 0 {
			reportError(w, nil, "update_live_session", "no record found")
			return
		}
		resp.ID = *req.ID
	} else {
		res, err := tx.Exec(
			"INSERT INTO liveSession "+
				"(UserID, URL, Title, Host, Comment, Begin, End, Tags)"+
				"VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
			userID, req.URL, req.Title, req.Host, req.Comment, req.Begin, req.End, encodeList(req.Tags),
		)
		if err != nil {
			reportError(w, err, "update_live_session", "database error")
			return
		}
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			reportError(w, err, "update_live_session", "database error")
			return
		}
		resp.ID = lastInsertID
	}
	tx.Commit()
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "update_live_session", "failed to generate response")
		return
	}
}
