package utils

import "fmt"

type DsnConfig struct {
	Port     string
	User     string
	Pass     string
	Host     string
	Database string
}

func MakeMysqlDsn(c DsnConfig) string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Database,
	)
}
