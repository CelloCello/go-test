package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "mycasdb"
	// cluster.Consistency = gocql.Quorum
	// //设置连接池的数量,默认是2个（针对每一个host,都建立起NumConns个连接）
	// cluster.NumConns = 3

	session, _ := cluster.CreateSession()
	time.Sleep(1 * time.Second) // Sleep so the fillPool can complete.
	// fmt.Println(session.Pool.Size())
	defer session.Close()

	query := session.Query("SELECT * FROM user").Iter()
	// if query != nil {
	// 	log.Fatal(query)
	// }

	var id int
	var name string
	fmt.Println("=== LIST ===")
	for query.Scan(&id, &name) {
		fmt.Println("User:", id, name)
	}
}
