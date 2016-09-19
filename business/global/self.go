package global

import (
	"github.com/Unknwon/goconfig"
	"net/rpc"
)

//定义自身工程中的全局结构体
type SELFCONFIG struct {
	IP            string
	PORT          string
	RdbClient     *rpc.Client
	MongodbClient *MONGODB
	RedisClient   *REDIS
}
type RDB struct {
	ip       string
	username string
	password string
	dbname   string
}

type MONGODB struct {
	ip       string
	username string
	password string
	dbname   string
}

type REDIS struct {
	ip       string
	username string
	password string
	dbname   string
}

//定义自身工程中的全局变量
var SelfConfig *SELFCONFIG

//定义自身工程中的全局函数
func GetRpcRdbClient(c *goconfig.ConfigFile) *rpc.Client {
	if SelfConfig != nil && SelfConfig.RdbClient != nil {
		return SelfConfig.RdbClient
	}
	rdb_host := GetSectionKeyValue(c, "RDB", "host", "127.0.0.1")
	rdb_port := GetSectionKeyValue(c, "RDB", "port", "5021")
	rpcC, _ := rpc.DialHTTP(TCP_PROTOCOL_L, rdb_host+":"+rdb_port)
	return rpcC
}

func LoadSelfConfig(c *goconfig.ConfigFile) *SELFCONFIG {
	return &SELFCONFIG{
		IP:            GetSectionKeyValue(c, "SELF", "host", "127.0.0.1"),
		PORT:          GetSectionKeyValue(c, "SELF", "port", "5011"),
		RdbClient:     GetRpcRdbClient(c),
		MongodbClient: nil,
		RedisClient:   nil,
	}
}
