package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {
	err := os.MkdirAll("./db", 0755)
	if err != nil {
		log.Fatalln(err)
	}

	db, err = sql.Open("sqlite3", "./db/qianliyun.db")
	if err != nil {
		log.Fatalln(err)
	}
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		log.Fatalln(err)
	}
	_, err = tx.Exec(
		"CREATE TABLE IF NOT EXISTS userInfo (" +
			"ID INTEGER PRIMARY KEY, " +
			"Username TEXT UNIQUE NOT NULL, " +
			"Password TEXT, " +
			"Salt TEXT, " +
			"Alias TEXT, " +
			"ResellerAlias TEXT, " +
			"AuthMax INTEGER, " +
			"AuthLeft INTEGER, " +
			"DeauthLeft INTEGER, " +
			"Reseller INTEGER" +
			");",
	)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = tx.Exec(
		"REPLACE INTO userInfo (ID, Username, AuthMax, AuthLeft, DeauthLeft) VALUES (1, 'root', 2147483647, 2147483647, 2147483647);",
	)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = tx.Exec(
		"CREATE TABLE IF NOT EXISTS liveSession (" +
			"ID INTEGER PRIMARY KEY, " +
			"UserID INTEGER NOT NULL, " +
			"URL TEXT, " +
			"Title TEXT, " +
			"Host TEXT, " +
			"Comment TEXT, " +
			"Begin REAL, " +
			"End REAL, " +
			"Tags TEXT" +
			");",
	)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = tx.Exec(
		"CREATE TABLE IF NOT EXISTS customerInfo (" +
			"ID INTEGER PRIMARY KEY, " +
			"UserID INTEGER NOT NULL, " +
			"CustomerName TEXT, " +
			"Mobile TEXT, " +
			"Status TEXT, " +
			"Tags TEXT" +
			");",
	)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = tx.Exec(
		"CREATE TABLE IF NOT EXISTS liveActivity (" +
			"ID INTEGER PRIMARY KEY, " +
			"UserID INTEGER NOT NULL, " +
			"LiveID INTEGER, " +
			"Time REAL, " +
			"CustomerID INTEGER, " +
			"Activity TEXT" +
			");",
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}
