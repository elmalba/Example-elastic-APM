package db

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "go.elastic.co/apm/module/apmgorm/dialects/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint
	FullName string
	Point    int
}

func init() {
	var err error
	// DB, err = apmgorm.Open("sqlite3", "tennisATP.db")
	DB, err = gorm.Open(sqlite.Open("tennisATP.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("EXIT PROBLEM DB", err)
	}

	//db.AutoMigrate(&User{})
	fmt.Println("ACAA")

}

func GetDB(ctx *gin.Context) *gorm.DB {
	db := DB
	db = DB.WithContext(ctx)
	return db
}
