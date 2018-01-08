package main

import (
	"context"
	"database/sql"
	"net/http"
)

type updateCustomerInfoReq struct {
	customerInfo
	LiveID *int64 `json:"live_id"`
}

type updateCustomerInfoResp struct {
	Error *string `json:"error"`
	ID    int64   `json:"id"`
}

func updateCustomerInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		reportError(w, nil, "update_customer_info", "failed to retrieve log in state")
		return
	}
	if userID == nil {
		reportError(w, nil, "update_customer_info", "unauthorized operation")
		return
	}
	var req updateCustomerInfoReq
	err = decodeRequest(r, &req)
	if err != nil {
		reportError(w, nil, "update_customer_info", "invalid request")
		return
	}
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		reportError(w, err, "update_customer_info", "database error")
		return
	}
	var resp updateCustomerInfoResp
	if req.ID != nil {
		resp.ID = *req.ID
	} else if req.CustomerName == nil {
		reportError(w, nil, "update_customer_info", "empty query criteria")
		return
	} else {
		err = tx.QueryRow(
			"SELECT ID FROM customerInfo WHERE "+
				"UserID = ? AND "+
				"CustomerName = ? "+
				"LIMIT 1;",
			userID, req.CustomerName,
		).Scan(&resp.ID)
		if err == sql.ErrNoRows {
			res, err := tx.Exec(
				"INSERT INTO customerInfo "+
					"(UserID, CustomerName)"+
					"VALUES (?, ?);",
				userID, req.CustomerName,
			)
			if err != nil {
				reportError(w, err, "update_customer_info", "database error")
				return
			}
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				reportError(w, err, "update_customer_info", "database error")
				return
			}
			resp.ID = lastInsertID
		}
	}
	res, err := tx.Exec(
		"UPDATE customerInfo SET "+
			"CustomerName = IFNULL(?, CustomerName), "+
			"Mobile = IFNULL(?, Mobile), "+
			"Status = IFNULL(?, Status), "+
			"Tags = IFNULL(?, Tags) "+
			"WHERE ID = ? AND UserID = ?;",
		req.CustomerName, req.Mobile, req.Status, encodeList(req.Tags), resp.ID, userID,
	)
	if err != nil {
		reportError(w, err, "update_customer_info", "database error")
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		reportError(w, err, "update_customer_info", "database error")
		return
	}
	if rowsAffected == 0 {
		reportError(w, err, "update_customer_info", "no record affected")
		return
	}
	resp.ID = *req.ID
	if req.LiveID != nil {
		var liveViewerRecordCount int
		err = tx.QueryRow(
			"SELECT COUNT(*) FROM liveViewer WHILE "+
				"UserID = ? AND "+
				"LiveID = ? AND "+
				"CustomerID = ?;",
			userID, req.LiveID, resp.ID,
		).Scan(&liveViewerRecordCount)
		if err != nil {
			reportError(w, err, "update_customer_info", "database error")
			return
		}
		_, err = tx.Exec(
			"INSERT INTO liveViewer (UserID, LiveID, CustomerID) VALUES (?, ?, ?);",
			userID, req.LiveID, resp.ID,
		)
		if err != nil {
			reportError(w, err, "update_customer_info", "database error")
			return
		}
	}
	tx.Commit()
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, nil, "update_customer_info", "failed to generate response")
		return
	}
}
