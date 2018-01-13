package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type telemetryReq struct {
	Data string `json:"data"`
}

type telemetryResp struct {
	Error *string `json:"error"`
	UUID  string  `json:"uuid"`
}

var telemetryLogger *log.Logger

func init() {
	telemetryLogger = log.New(&lumberjack.Logger{
		Filename:   "./logs/telemetry.log",
		MaxSize:    64,
		MaxBackups: 52,
		MaxAge:     7,
		Compress:   true,
	}, "", log.LstdFlags)
}

func telemetryHandler(w http.ResponseWriter, r *http.Request) {
	var req telemetryReq
	err := decodeRequest(r, &req)
	if err != nil {
		reportError(w, err, "login", "invalid request")
		return
	}
	resp := &telemetryResp{
		Error: nil,
		UUID:  uuid.New().String(),
	}
	telemetryLogger.Printf("==== BEGIN ENTRY %s ====\n", resp.UUID)
	telemetryLogger.Println(req.Data)
	telemetryLogger.Printf("==== END ENTRY %s ====\n", resp.UUID)
	err = encodeResponse(w, resp)
	if err != nil {
		reportError(w, err, "login", "failed to generate response")
		return
	}
}
