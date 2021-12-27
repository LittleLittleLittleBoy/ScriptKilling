package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/littlelittlelittleboy/scriptkilling/config"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB(config config.MySQLConfig) error {
	var err error
	Engine, err = xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database))
	return err
}
