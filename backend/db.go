package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//"gorm.io/driver/sqlite"

	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/sqlite"
)

var db *gorm.DB

type User struct {
	ID       uint
	FullName string
	Point    int
}

func init() {
	var err error
	db, err = apmgorm.Open("sqlite3", "tennisATP.db")
	//db, err = gorm.Open("sqlite3", "tennisATP.db")
	if err != nil {
		log.Fatalln("EXIT PROBLEM DB", err)
	}

	//db.AutoMigrate(&User{})

}

func getDB(ctx *gin.Context) *gorm.DB {

	//return db
	return apmgorm.WithContext(ctx.Request.Context(), db)
}
