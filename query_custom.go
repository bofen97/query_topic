package main

import (
	"fmt"
	"net/http"

	SQLConn "github.com/bofen97/sqlc"
)

type QueryCustomTopic struct {
	sqlc *SQLConn.SQLConn
}

func (qt *QueryCustomTopic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		topic := r.URL.Query().Get("topic")
		//arxiv query .
		data, err := qt.sqlc.QueryCustomTopicFromArxiv(topic)
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
