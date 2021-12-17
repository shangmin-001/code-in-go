package main

import (
	"flag"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	consumerLibrary "github.com/aliyun/aliyun-log-go-sdk/consumer"
	"go.uber.org/zap"
	"strings"
)

var (
	listConsumerGroup   bool
	createConsumerGroup string
	deleteConsumerGroup string
	updateConsumerGroup string
)

func init() {
	flag.BoolVar(&listConsumerGroup, "listConsumerGroup", false, "list consumer group")
	flag.StringVar(&createConsumerGroup, "createConsumerGroup", "", "create consumer group,split by ,")
	flag.StringVar(&deleteConsumerGroup, "deleteConsumerGroup", "", "delete consumer group,split by ,")
	flag.StringVar(&updateConsumerGroup, "updateConsumerGroup", "", "update consumer group, old1:new1,old2:old2")
}

func (client *Client) consumerGroupOperation() {
	if listConsumerGroup {
		client.listConsumerGroup()
	}
	if len(createConsumerGroup) != 0 {
		groups := strings.Split(createConsumerGroup, ",")
		for _, group := range groups {
			client.createConsumerGroup(group)
		}
	}
}

func (client *Client) listConsumerGroup() {
	cList, err := client.ListConsumerGroup(client.ProjectName, client.LogStoreName)
	if err != nil {
		logger.Error("list consumer group meet err : %v", zap.Error(err))
		return
	}
	groups := make([]string, 0)
	for _, group := range cList {
		groups = append(groups, group.ConsumerGroupName)
	}
	logger.Info("consumer group : ", zap.Any("group", groups))
}

func (client *Client) createConsumerGroup(group string) {
	cGroup := sls.ConsumerGroup{
		ConsumerGroupName: group,
		Timeout:           100,
	}
	err := client.CreateConsumerGroup(client.ProjectName, client.LogStoreName, cGroup)
	if err != nil {
		logger.Error("create consumer group meet err : ", zap.Error(err))
	}
}

func (client *Client) PullDataByConsumerGroup() {
	if !client.EnableCG {
		return
	}
	if client.ConsumerGroup.ConsumerGroupName == "" {
		logger.Sugar().Info("consumer group name should not be empty")
	}
	if client.ConsumerGroup.ConsumerName == "" {
		logger.Sugar().Info("consumer  name should not be empty")
	}
	option := consumerLibrary.LogHubConfig{
		Endpoint:          client.Endpoint,
		AccessKeyID:       client.AccessKeyID,
		AccessKeySecret:   client.AccessKeySecret,
		Project:           client.ProjectName,
		Logstore:          client.LogStoreName,
		ConsumerGroupName: client.ConsumerGroupName,
		ConsumerName:      client.ConsumerName,
		CursorStartTime:   client.ConsumerGroup.CursorStartTime,
		InOrder:           client.InOrder,
		CursorPosition:    client.ConsumerGroup.CursorPosition,
	}
	consumerWorker := consumerLibrary.InitConsumerWorker(option, client.printfLog)
	consumerWorker.Start()

}
