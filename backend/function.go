package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Player struct {
	ID       int `json:"key"`
	FullName string
	Point    int
}

func ranking(ctx *gin.Context) {
	var players []Player
	var sql, query string
	query = ctx.Request.URL.Query().Get("q")
	if query == "" {
		sql = `select id, full_name,point from users`
	} else {
		sql = fmt.Sprintf(`select id, full_name,point from users where full_name like '%%%s%%'`, query)

	}

	getDB(ctx).Raw(sql).Scan(&players)
	ctx.JSON(200, players)
}
