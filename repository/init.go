package repository

import (
	"context"
	"fmt"
	"github.com/RezaOptic/gin-project-structure/config"
	"github.com/RezaOptic/gin-project-structure/constant"
	"github.com/couchbase/gocb/v2"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/scylladb/gocqlx/v2"
	"gitlab.com/RezaOptic/go-utils/logger"
	"time"

	"google.golang.org/grpc"
)

// dbs struct for managing database connections
type dbs struct {
	Redis    redis.UniversalClient
	Couch    *gocb.Cluster
	Dgraph   *dgo.Dgraph
	MinIO    *minio.Client
	ScyllaDB gocqlx.Session
}

var DBS dbs

// Init function for init databases
func Init() {
	redisConnection()
	couchBaseConnection()
	dgraphConnection()
	minIOConnection()
	scylladbConnection()
}

// redisConnection function for connecting to redis server
func redisConnection() {
	opt := redis.UniversalOptions{
		Addrs: config.Configs.Redis.Addresses,
	}
	DBS.Redis = redis.NewUniversalClient(&opt)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := DBS.Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis connected with result :%s \n", result)
}

// couchBaseConnection function for connection to couchbase server
func couchBaseConnection() {
	cluster, err := gocb.Connect(
		config.Configs.Couchbase.Addresses,
		gocb.ClusterOptions{
			Username: config.Configs.Couchbase.Username,
			Password: config.Configs.Couchbase.Password,
		})
	if err != nil {
		panic(err)
	}
	result, err := cluster.Bucket(constant.UsersBucket).Ping(&gocb.PingOptions{ServiceTypes: []gocb.ServiceType{gocb.ServiceTypeQuery}})
	if err != nil {
		panic(err)
	}
	logger.ZSLogger.Infow("couchbase connection ... ")
	logger.ZSLogger.Infof("%#v", result.Services)
	DBS.Couch = cluster
}

// GetUserCollection method for get couchbase user collections
func (d *dbs) GetUserCollection() *gocb.Collection {
	return DBS.Couch.Bucket(constant.UsersBucket).DefaultCollection()
}

// couchBaseConnection function for connection to couchbase server
func dgraphConnection() {
	var conns []api.DgraphClient
	for _, i := range config.Configs.Dgraph.Addresses {
		conn, err := grpc.Dial(i, grpc.WithInsecure())
		if err != nil {
			logger.ZSLogger.Panicf("error on connect to dgraph server with error :%s", err)
		}
		conns = append(conns, api.NewDgraphClient(conn))
	}
	DBS.Dgraph = dgo.NewDgraphClient(conns...)
}

// minIOConnection function for connecting to minIO server
func minIOConnection() {
	opt := minio.Options{
		Creds: credentials.NewStaticV4(config.Configs.MinIO.AccessKeyID, config.Configs.MinIO.SecretAccessKey, ""),
	}
	MinIO, err := minio.New(config.Configs.MinIO.Addresses, &opt)
	if err != nil {
		logger.ZSLogger.Panicf("error on connect to min io server with error :%s", err)
	}
	DBS.MinIO = MinIO
}

// scylladbConnection function for connection to scylladb server
func scylladbConnection() {
	cluster := gocql.NewCluster(config.Configs.Scylladb.Addresses...)
	cluster.Keyspace = config.Configs.Scylladb.KeySpace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logger.ZSLogger.Panicf("error on connect to scylladb server with error :%s", err)
	}
	DBS.ScyllaDB = session
}
