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

	mux := http.NewServeMux()
	mux.Handle("/query", query)
	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

//http://export.arxiv.org/api/query? {parameters}

/*
parameters

	search_query=all:electron
	search_query=all:electron+AND+all:proton
*/

/*
http://export.arxiv.org/api/query?search_query=all:electron&start=0&max_results=10 (1)
http://export.arxiv.org/api/query?search_query=all:electron&start=10&max_results=10 (2)
http://export.arxiv.org/api/query?search_query=all:electron&start=20&max_results=10 (3)
*/

/*
For example to retrieve matches 6001-8000: http://export.arxiv.org/api/query?search_query=all:electron&start=6000&max_results=8000
*/

/*
https://export.arxiv.org/api/query?search_query=ti:%22electron%20thermal%20conductivity%22&sortBy=lastUpdatedDate&sortOrder=descending
*/

// source && pdf

//http://export.arxiv.org/api/query?search_query=au:del_maestro

//http://export.arxiv.org/api/query?search_query=au:del_maestro+AND+ti:checkerboard
//http://export.arxiv.org/api/query?search_query=au:del_maestro+ANDNOT+ti:checkerboard
