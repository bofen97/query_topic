package main

import (
	"log"
	"net"
	"os"

	serviceRegister "github.com/bofen97/ServiceRegister"
	SQLConn "github.com/bofen97/sqlc"
	"google.golang.org/grpc"
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

	etcdServer := os.Getenv("etcdserver")
	if serverPort == "" {
		log.Fatal("etcdServer is none")
		return
	}

	sr, err := serviceRegister.NewRegisteService([]string{
		etcdServer,
	}, 5)
	if err != nil {
		log.Fatal(err)
	}
	go sr.ListenLaser()
	go sr.PutServiceAddr("query_topic", serviceRegister.GetCurrentIP()+serverPort)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcsqlc := new(SQLConn.SQLConn)
	if err := grpcsqlc.Connect(sqlurl); err != nil {
		log.Fatal(err)
	}
	log.Print("grpc sqlc connect !")

	RegisterQueryServer(s, &QueryTopics{sqlc: grpcsqlc})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
