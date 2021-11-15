package main

import (
	"net/http"

	"github.com/jackc/pgx/v4"
)

func handlerGood(conn *pgx.Conn, req *http.Request) {
	q := "SELECT ITEM,PRICE FROM PRODUCT WHERE ITEM_CATEGORY='?' ORDER BY PRICE"
	conn.Query(q, req.URL.Query()["category"])
}
