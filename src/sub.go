package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := "root"
  PASS     := "####"
  PROTOCOL := "tcp(##.###.##.###:3306)"
  DBNAME   := "##"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){
  db := gormConnect()
  defer db.Close()
}