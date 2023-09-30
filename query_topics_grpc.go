package main

import (
	context "context"
	"encoding/json"
	"log"

	SQLConn "github.com/bofen97/sqlc"
)

type QueryTopics struct {
	sqlc *SQLConn.SQLConn
}

func (queryg *QueryTopics) QueryCustom(ctx context.Context, in *QueryCustomArg) (*QueryCustomRets, error) {
	topic := in.GetTopic()
	log.Printf("Got query [%s]\n", topic)
	data, err := queryg.sqlc.QueryCustomTopicFromArxiv(topic)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var jdatas []*QueryCustomRet
	json.Unmarshal(data, &jdatas)

	return &QueryCustomRets{Querys: jdatas}, nil

}
func (queryg *QueryTopics) QueryTopic(ctx context.Context, in *QueryTopicArg) (*QueryCustomRets, error) {
	topic := in.GetTopic()
	date := in.GetDate()
	log.Printf("Got query [%s]  date [%s] \n", topic, date)

	data, err := queryg.sqlc.QueryTitleAuthorsSummaryId(topic, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var jdatas []*QueryCustomRet
	json.Unmarshal(data, &jdatas)

	return &QueryCustomRets{Querys: jdatas}, nil

}

func (*QueryTopics) mustEmbedUnimplementedQueryServer() {}
