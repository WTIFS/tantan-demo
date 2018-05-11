package dao

import (
	"github.com/go-pg/pg"
	"github.com/WTIFS/tantan-demo/config"
)

func GetConn() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr: config.PG_ADDRESS,
		User: config.PG_USER,
		Password: config.PG_PASSWORD,
		Database: config.PG_DATABASE,
	})
	return conn
}