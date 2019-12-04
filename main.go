package main

import (
    "fmt"
//     "strconv"
//     "log"
    "flag"

    "github.com/welltok/go-commons/errutils"
    "github.com/welltok/metrics-store/c7a"
   // "github.com/gocql/gocql"
)

// var Session *gocql.Session

// keyspace = data replication on nodes
var (

// Specifies a simple replication factor for the cluster.
    c7aHosts = flag.String(
    		"c7a_hosts",
    		fmt.Sprintf("%s:%s", os.Getenv("MIP"),
    			os.Getenv("CPORT")),
    		"Cassandra URL host1[:port1],host2[:port2]",
    	)
    	c7aStrategy = flag.String(
    		"c7a_strategy",
    		"simple:1",
    		"Cassandra strategy",
    	)
    	c7aDebug = flag.Bool(
    		"c7a_debug",
    		false,
    		"Cassandra debug mode",
    	)
    	c7aConsistency = flag.String(
    		"c7a_consistency",
    		"local_quorum",
    		"Cassandra consistency level",
    	)
    	c7aKeyspace = flag.String(
    		"c7a_keyspace",
    		"oddjob",
    		"Cassandra keyspace name",
    	)
    	c7aUser = flag.String(
    		"c7a_user",
    		os.Getenv("C7A_USER"),
    		"Cassandra username",
    	)
    	c7aPwd = flag.String(
    		"c7a_pwd",
    		os.Getenv("C7A_PWD"),
    		"Cassandra password",
    	)
    	c7aCertPath = flag.String(
    		"c7a_cert_path",
    		os.Getenv("C7A_CERT_PATH"),
    		"Cassandra SSL Certificate Path",
    	)
    	c7aKeyPath = flag.String(
    		"c7a_key_path",
    		os.Getenv("C7A_KEY_PATH"),
    		"Cassandra SSL Key Path",
    	)
    	c7aCAPath = flag.String(
    		"c7a_ca_path",
    		os.Getenv("C7A_CA_PATH"),
    		"Cassandra SSL CA Path",
    	)
    	c7aCQLVersion = flag.String(
    		"c7a_cql_version",
    		os.Getenv("C7A_CQL_VERSION"),
    		"Cassandra CQL version",
    	)
    	c7aProtoVersion = flag.String(
    		"c7a_proto_version",
    		os.Getenv("C7A_PROTO_VERSION"),
    		"Cassandra Proto version",
    	)
    	debug = flag.Bool(
    		"debug",
    		false,
    		"Toggle service debug flag",
    	)
)

func initCassandra() (*c7a.Client, error) {
	config := c7a.Config{
        Hosts:        *c7aHosts,
		Debug:        *c7aDebug,
		Consistency:  *c7aConsistency,
		Strategy:     *c7aStrategy,
		Keyspace:     *c7aKeyspace,
		Username:     *c7aUser,
		Password:     *c7aPwd,
		CQLVersion:   *c7aCQLVersion,
		CertPath:     *c7aCertPath,
		KeyPath:      *c7aKeyPath,
		CAPath:       *c7aCAPath,
		ProtoVersion: 4,
	}
	return c7a.NewClient(config)
}

func getProtoVersion() int {
	pv := 4
	return pv
}

func main() {
    	client, err := initCassandra()
    	fmt.Println("Test 123")
    	errutils.Panic(err)
    	errutils.Panic(client.EnsureKeyspace(false))
    	errutils.Panic(client.EnsureSchema(*drop))
      //  getResults()
}

// func getResults() {
//     var storeResults []string
//        if err := Session.Query(`SELECT * FROM oddjob.fitness\n\tWHERE validic_id='cw__f3de4039-b2ac-4b34-80d2-e0a11fe236bf'\n\tAND source='manual'\n\tAND time >= '2019-07-16 00:00:00+0000' AND time <= '2019-07-17 23:59:59+0000'`,
//       storeResults ).Exec(); err != nil {
//             fmt.Println("Error")
//             fmt.Println(err)
//             }
//       fmt.Println(storeResults)
//
// }