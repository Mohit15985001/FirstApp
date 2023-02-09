package database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBConnnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

type Query struct {
	Query       string
	Projections []interface{}
	ProcessRow  func()
}

var connection DBConnnection

func SetupDBConnection() {
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "test_db"
	connection.session, _ = connection.cluster.CreateSession()

}

func ExecuteInsertQuery(query string, values ...interface{}) {
	if err := connection.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal((err))
	}

}
func ExecuteUpdateQuery(query string, values ...interface{}) {
	if err := connection.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal((err))
	}

}

func (q *Query) ExecuteSelectQuery() {

	iter := connection.session.Query(q.Query).Iter()
	for iter.Scan(q.Projections...) {
		q.ProcessRow()
	}
	iter.Close()

}
