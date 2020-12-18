package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

//データベースを操作するための変数を定義
var DBInstance *sql.DB

func init() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	//os.Getenvで指定したキーの環境変数が取得できる
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	//user:password@tcp(host:port)/dbname
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbuser, dbpassword, dbhost, dbname)

	var err error
	//open関数でDBに接続
	DBInstance, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	log.Printf("データベースへの接続に成功しました;%s\n", dataSource)
}
