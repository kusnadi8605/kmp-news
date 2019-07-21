package main

import (
	"flag"
	"fmt"
	conf "kmp-news/config"
	hdlr "kmp-news/handler"
	mdw "kmp-news/middleware"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load config file
	configFile := flag.String("conf", "config/config.yml", "main configuration file")

	flag.Parse()
	conf.LoadConfigFromFile(configFile)
	conn, err := conf.New(conf.Param.DBType, conf.Param.DBUrl)
	conf.Logf("Load Database Conf: %s ", conf.Param.DBType)
	conf.Logf("running App on port: %s ", conf.Param.ListenPort)

	elasticConn := conf.ElasticConn(conf.Param.ElasticURL)

	if err != nil {
		conf.Logf("Load Database Conf: %s ", err)
		log.Fatal(err)
	}

	http.HandleFunc("/api/news", mdw.Chain(hdlr.GetNewsHandler(conn, elasticConn, conf.Param.ElasticIndex, conf.Param.ElasticPerpage), mdw.Method("GET")))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		conf.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
