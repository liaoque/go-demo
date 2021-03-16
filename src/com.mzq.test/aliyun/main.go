package main

import (
	"encoding/json"
	"flag"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/consumer"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var option consumerLibrary.LogHubConfig
var client sls.Client
var logStore *sls.LogStore

type AKeyVal struct {
	Key string `json:"key"`
	Value string`json:"value"`
}

type AContent struct {
	Contents []AKeyVal
}

type ALog struct {
	LogTags []map[string]string `json:"LogTags"`
	Logs []AContent `json:"Logs"`
	Source string `json:"Source"`
	Topic string `json:"Topic"`
}

type StroJson struct {
	path string
	name string
}
var s StroJson

func main() {


	flag.StringVar(&s.path, "d", "", "指定目录")
	flag.StringVar(&s.name, "n", "", "指定消费者")
	flag.Parse()
	if len(s.path) == 0 {
		log.Println("请指定目录：go_build_main_go -d /alidata/xxxx/")
		return
	}
	if !IsDir(s.path) {
		log.Println("目录不存在, 请检查目录：go_build_main_go -d /alidata/xxxx/")
		return
	}
	if len(s.name) == 0 {
		log.Println("请制定消费者名字 -n xxx")
		return
	}


	option := consumerLibrary.LogHubConfig{
		Endpoint:          "cn-beijing.log.aliyuncs.com",
		AccessKeyID:       "LmzqTAI4FwibmFX3tXhwjh14RiM",
		AccessKeySecret:   "4mzqjPIo1drbkGobb7cuxeRO2MpaaPb8S",
		Project:           "com-asyke-playground",
		Logstore:          "class_run",
		ConsumerGroupName: "go-com-asyke-playground-" + s.name,
		ConsumerName:      s.name,
		// This options is used for initialization, will be ignored once consumer group is created and each shard has been started to be consumed.
		// Could be "begin", "end", "specific time format in time stamp", it's log receiving time.
		//CursorStartTime: 1593237688,
		//CursorPosition: consumerLibrary.BEGIN_CURSOR,
		CursorPosition: consumerLibrary.END_CURSOR,
	}
	//println(option.ConsumerGroupName)

	consumerWorker := consumerLibrary.InitConsumerWorker(option, process)
	ch := make(chan os.Signal)
	signal.Notify(ch)
	consumerWorker.Start()
	if _, ok := <-ch; ok {
		//level.Info(consumerWorker.Logger).Log("msg", "get stop signal, start to stop consumer worker", "consumer worker name", option.ConsumerName)
		consumerWorker.StopAndWait()
	}
}

func process(shardId int, logGroupList *sls.LogGroupList) string {
	//fmt.Println(shardId, logGroupList)
	for _, logGroup := range logGroupList.LogGroups {
		marshal, err := json.Marshal(logGroup)
		if err == nil {
			str := decodeLog(marshal)
			jsonStr, _ := json.Marshal(str)
			courseClassUserRunStartId := str["course_class_user_run_start_id"]
			log.Println("获取跑步id: %v", courseClassUserRunStartId)
			fileName := s.path + "/" + courseClassUserRunStartId + ".json"
			exist := IsExist(fileName)
			if exist {
				log.Println("文件已存在: %v", fileName)
				break
			}
			ioutil.WriteFile(fileName, jsonStr, os.ModePerm)
			time.Sleep(1 * time.Second)
			resp, err := http.Get("http://app.asyke.com/api/v1/cache/ali-log?course_class_user_run_start_id=" + courseClassUserRunStartId)
			if err != nil{
				break;
			}
			resp.Body.Close()
		}
	}

	return ""
}

func decodeLog(b []byte) map[string]string {
	var someOne ALog;
	_ = json.Unmarshal(b, &someOne)
	var m = make(map[string]string, 200)
	for _,v := range someOne.Logs {
		for _, v2:= range v.Contents {
			m[v2.Key] =v2.Value
		}
	}
	return m
}

func IsExist(fileAddr string)(bool){
	// 读取文件信息，判断文件是否存在
	_,err := os.Stat(fileAddr)
	if err!=nil{
		if os.IsExist(err){  // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

func IsDir(fileAddr string)bool{
	s,err:=os.Stat(fileAddr)
	if err!=nil{
		return false
	}
	return s.IsDir()
}
