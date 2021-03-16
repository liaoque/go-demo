package core

import (
	"com.mzq.echo/app/core/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/olivere/elastic/v7"
	"sync"
)
var appEcho *echo.Echo
var db *gorm.DB
var esClient *elastic.Client

func init()  {
	appEcho = newEcho();
}


func newEcho() *echo.Echo {
	e := echo.New()
	return e
}

func App() *echo.Echo {
	return appEcho
}

func Db() *gorm.DB  {
	if db != nil {
		return db
	}
	mutex := &sync.Mutex{}
	mutex.Lock()
	if db != nil {
		mutex.Unlock()
		return db
	}

	//"root:123456@(192.168.3.5:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	conf := config.Db().User + ":" + config.Db().Pwd + "@(" + config.Db().Host + ":3306)/" +
		config.Db().Db + "?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "root:123456@(192.168.3.5:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", conf)
	if err!= nil{
		mutex.Unlock()
		panic(err)
	}
	//设置 空闲连接池
	db.DB().SetMaxIdleConns(10)
	//设置 最大连接池
	db.DB().SetMaxOpenConns(100)
	mutex.Unlock()
	return db
}

func Elastic() *elastic.Client  {
	if esClient != nil {
		return esClient
	}
	mutex := &sync.Mutex{}
	mutex.Lock()
	if esClient != nil {
		mutex.Unlock()
		return esClient
	}

	esClient, err := elastic.NewClient(elastic.SetURL("http://192.168.3.5:9200"))
	if err != nil {
		mutex.Unlock()
		panic(err)
	}
	mutex.Unlock()
	return esClient
}