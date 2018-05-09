package dao

import "github.com/go-pg/pg"

func GetConn() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr: "localhost:5432",
		User: "postgres",
		Password: "www",
		Database: "putong",
	})
	return conn
}