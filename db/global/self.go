package global

import (
	"github.com/Unknwon/goconfig"
)

//定义自身工程中的全局结构体
type SELFCONFIG struct {
	IP            string
	PORT          string
	RdbClient     *RDB
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
func GetRdbClient(c *goconfig.ConfigFile) *RDB {
	GetSectionKeyValue(c, "RDB", "host", "127.0.0.1")
	GetSectionKeyValue(c, "RDB", "port", "5432")
	return nil
}

func GetMongodbClient(c *goconfig.ConfigFile) *MONGODB {
	GetSectionKeyValue(c, "RDB", "host", "127.0.0.1")
	GetSectionKeyValue(c, "RDB", "port", "27017")
	return nil
}

func GetRedisClient(c *goconfig.ConfigFile) *REDIS {
	GetSectionKeyValue(c, "RDB", "host", "127.0.0.1")
	GetSectionKeyValue(c, "RDB", "port", "3306")
	return nil
}

func LoadSelfConfig(c *goconfig.ConfigFile) *SELFCONFIG {
	return &SELFCONFIG{
		IP:            GetSectionKeyValue(c, "SELF", "host", "127.0.0.1"),
		PORT:          GetSectionKeyValue(c, "SELF", "port", "5011"),
		RdbClient:     GetRdbClient(c),
		MongodbClient: GetMongodbClient(c),
		RedisClient:   GetRedisClient(c),
	}
}
