package config

import "strconv"

type Db struct {
	User     string
	Password string
	DbName   string
	Ip       string
	Port     int
}

func GetDsn(d Db) string {
	dsn := d.User + ":" + d.Password + "@" + "tcp" + "(" + d.Ip + ":" + strconv.Itoa(d.Port) + ")" + "/" + d.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}
