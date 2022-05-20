package main

import (
	"atp"

	"github.com/gin-gonic/gin"
	//"go.elastic.co/apm/module/apmgin"
)

func main() {
	r := gin.Default()
	//r.Use(apmgin.Middleware(r))
	r.GET("/api/ranking", atp.Ranking)

	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
