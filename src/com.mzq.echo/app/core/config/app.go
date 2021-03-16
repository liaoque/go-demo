package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	AppKey string `json:"app_key"`
}
var app AppConfig

func App () (*AppConfig){
	return &app;
}

func init()  {
	file, err := os.Open("./config/com.mzq.echo/app.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err3 := decoder.Decode(&app)
	if err != nil {
		fmt.Println(err3)
		panic(err3)
	}
}


