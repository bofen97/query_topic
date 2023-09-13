package main

import (
	"fmt"
	"net/http"

	SQLConn "github.com/bofen97/sqlc"
)

type QueryTopic struct {
	sqlc *SQLConn.SQLConn
}

func (qt *QueryTopic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		topic := r.URL.Query().Get("topic")
		date := r.URL.Query().Get("date")

		data, err := qt.sqlc.QueryTitleAuthorsSummaryId(topic, date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "[ERROR] URL ERROR %v \n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", data)
		return

	}
	w.WriteHeader(http.StatusBadRequest)

}
