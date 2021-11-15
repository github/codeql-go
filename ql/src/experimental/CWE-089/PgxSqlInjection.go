package main

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func handler(conn *pgx.Conn, req *http.Request) {
	q := fmt.Sprintf("SELECT ITEM,PRICE FROM PRODUCT WHERE ITEM_CATEGORY='%s' ORDER BY PRICE",
		req.URL.Query()["category"])
	conn.Query(q)
}
