package main

import (
	"github.com/MarkTBSS/096_Coin_Adding/config"
	"github.com/MarkTBSS/096_Coin_Adding/databases"
	"github.com/MarkTBSS/096_Coin_Adding/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
