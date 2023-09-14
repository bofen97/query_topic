package main

import (
	"log"
	"net/http"
	"os"

	SQLConn "github.com/bofen97/sqlc"
)

// const mysqlUrl = "tyFeng:J0]nt4D_3-NbO>8|GgryV-ry.?G{@tcp(arxivinfo.cvheva0xliby.us-east-1.rds.amazonaws.com:3306)/arxivInfo?parseTime=true"

func main() {
	sqlurl := os.Getenv("sqlurl")
	if sqlurl == "" {
		log.Fatal("sqlurl is none")
		return
	}
	serverPort := os.Getenv("serverport")
	if serverPort == "" {
		log.Fatal("serverPort is none")
		return
	}
	query := new(QueryTopic)
	query.sqlc = new(SQLConn.SQLConn)
	if err := query.sqlc.Connect(sqlurl); err != nil {
		log.Fatal(err)
	}
	log.Print("connect !")

	query_custom := new(QueryCustomTopic)
	query_custom.sqlc = new(SQLConn.SQLConn)

	mux := http.NewServeMux()
	mux.Handle("/query", query)
	mux.Handle("/query_custom", query_custom)
	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
