package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DbConfig struct {
	Host string
	User string
	Pwd string
	Db string
}
var configDb DbConfig

func Db () (*DbConfig){
	return &configDb;
}

func init()  {
	file, err := os.Open("./config/com.mzq.echo/db.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err3 := decoder.Decode(&configDb)
	if err != nil {
		fmt.Println(err3)
		panic(err3)
	}
}


