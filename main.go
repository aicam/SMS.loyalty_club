package main

import (
	"github.com/aicam/scotechShops/internal/database"
	"github.com/aicam/scotechShops/internal/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)
// database info
const DatabaseConnectionString  = "aicam:021021ali@tcp(127.0.0.1:3306)/shop_data?charset=utf8mb4&parseTime=True"


func main() {
	s := handler.NewServer()
	db := database.ConnectMysql(DatabaseConnectionString)
	database.SqlMigrate(db)
	s.Routes(db)
	err := http.ListenAndServe("0.0.0.0:4300", s.Router)
	if err != nil {
		log.Print(err)
	}
}
