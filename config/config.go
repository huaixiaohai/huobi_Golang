package config

type DBConfig struct {
	Name            string
	URL             string
	ConnMaxLifeTime int64
	MaxOpenConns    int64
	MaxIdleConns    int64
}

var DB = &DBConfig{
	Name:            "mysql",
	URL:             "root:2015.ami@tcp(localhost:3306)/hb?charset=utf8&parseTime=true",
	ConnMaxLifeTime: 100,
	MaxOpenConns:    100,
	MaxIdleConns:    10,
}

var Host = "api.huobi.pro"
var AccessKey = "hrf5gdfghe-38992eb9-52319e1f-1065d"
var AccountId = "14794244"
var SubUid int64 = 153730843
var SecretKey = "5b7aa865-6306a844-61c52dc7-defd7"
