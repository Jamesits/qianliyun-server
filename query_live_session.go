package main

import (
	"context"
	"database/sql"
	"net/http"
)

type queryLiveSessionReq struct {
	liveSession
}

type queryLiveSessionResp struct {
	Error       *string       `json:"error"`
	LiveSession []liveSession `json:"live_session"`
}

func queryLiveSessionHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "query_live_session", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "query_live_session", "unauthorized operation")
		return
	}
	var req queryLiveSessionReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "query_live_session", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		reportError(w, err, "query_live_session", "database error")
		return
	}
	defer tx.Commit()
	rows, err := tx.Query(
		"SELECT * FROM liveSession WHERE "+
			"IFNULL(ID = ?, 1) AND "+
			"UserID = ? AND "+
			"IFNULL(URL = ?, 1) AND "+
			"IFNULL(Title = ?, 1) AND "+
			"IFNULL(Host = ?, 1) AND "+
			"IFNULL(Comment = ?, 1) AND "+
			"IFNULL(Tags = ?, 1);",
		req.ID, userID, req.URL, req.Title, req.Host, req.Comment, encodeList(req.Tags),
	)
	if err != nil {
		reportError(w, err, "query_live_session", "database error")
		return
	}
	defer rows.Close()
	resp := queryLiveSessionResp{}
	resp.LiveSession = []liveSession{}
	for rows.Next() {
		var rec liveSession
		var tags *string
		err = rows.Scan(&rec.ID, &rec.UserID, &rec.URL, &rec.Title, &rec.Host, &rec.Comment, &rec.Begin, &rec.End, &tags)
		if err != nil {
			reportError(w, err, "query_live_session", "database error")
			return
		}
		rec.Tags = decodeList(tags)
		resp.LiveSession = append(resp.LiveSession, rec)
	}
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "query_live_session", "failed to generate response")
		return
	}
}
