package main

import (
	"flag"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const DefaultConsumerCount int = 100

var confFile string

func init() {
	flag.StringVar(&confFile, "f", "conf.yml", "global config file")
}

func UnmarshalConfig() (*Client, error) {
	var client Client
	f, err := ioutil.ReadFile(confFile)
	if err != nil {
		logger.Error("yamlFile.Get err   #%v \"", zap.Error(err))
		return nil, err
	}
	err = yaml.Unmarshal(f, &client)
	if err != nil {
		logger.Error("yamlFile.Unmarshal err   #%v \", err", zap.Error(err))
		return nil, err
	}
	logger.Info("config", zap.Any("client", &client))
	return &client, nil
}

type ConsumerGroup struct {
	EnableCG          bool
	ConsumerGroupName string
	ConsumerName      string
	InOrder           bool
	CursorPosition    string
	CursorStartTime   int64
	ConsumerCount     int
}

type Consumer struct {
	EnableC         bool
	CursorPosition  string
	ConsumerCount   int
	CursorStartTime int64
}

type Client struct {
	sls.Client
	ProjectName  string
	LogStoreName string
	ConsumerGroup
	Consumer
}
