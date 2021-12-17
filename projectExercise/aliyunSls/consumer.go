package main

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"strconv"
)

func (client *Client) PullDataByConsumer() {
	if !client.Consumer.EnableC {
		return
	}
	var beginCursor, endCursor string
	var err error
	var ok bool
	shards, err := client.ListShards(client.ProjectName, client.LogStoreName)
	if err != nil {
		logger.Sugar().Info("ListShards err :", err)
	}
	logger.Sugar().Info("shards :", shards)
	if client.Consumer.ConsumerCount == 0 {
		client.Consumer.ConsumerCount = DefaultConsumerCount
	}
	logger.Sugar().Info(client.Consumer.CursorStartTime)

	if client.Consumer.CursorPosition == "" {
		logger.Sugar().Info(client.Consumer.CursorStartTime)
		if client.Consumer.CursorStartTime > 0 {
			beginCursor = strconv.FormatInt(client.Consumer.CursorStartTime, 10)
			logger.Sugar().Info(beginCursor)
		} else {
			beginCursor = "begin"
		}
	} else {
		beginCursor = client.Consumer.CursorPosition
	}
	shardCursor := make(map[string]string)
	for client.Consumer.ConsumerCount > 0 {
		for _, shard := range shards {
			if beginCursor, ok = shardCursor[strconv.Itoa(shard.ShardID)]; !ok {
				logger.Sugar().Info(beginCursor)
				if beginCursor == "" {
					beginCursor = strconv.FormatInt(client.Consumer.CursorStartTime, 10)
					logger.Sugar().Info(beginCursor)
				}
				logger.Sugar().Info(beginCursor)
				if beginCursor, err = client.GetCursor(client.ProjectName, client.LogStoreName, shard.ShardID, beginCursor); err != nil {
					logger.Sugar().Error("get cursor error: ", err)
					return
				}

			}
			logger.Sugar().Info("cursor :", shard.ShardID, beginCursor)
			client.getTimeOfCursor(shard.ShardID, beginCursor)
			if endCursor, err = client.GetCursor(client.ProjectName, client.LogStoreName, shard.ShardID, "end"); err != nil {
				logger.Sugar().Error("get cursor error: ", err)
				return
			}
			client.getTimeOfCursor(shard.ShardID, endCursor)

			logger.Sugar().Info("cursor :", shard.ShardID, endCursor)
			logger.Sugar().Info("start to pull data ")
			logGroupList, nextCursor, err := client.PullLogs(client.ProjectName, client.LogStoreName, shard.ShardID, beginCursor, endCursor, 10)

			if err != nil {
				logger.Sugar().Errorf("pull shard logs failed, %v", err)
				return
			}
			logger.Sugar().Info("pull data successfully")
			shardCursor[strconv.Itoa(shard.ShardID)] = nextCursor
			client.Consumer.ConsumerCount -= len(logGroupList.LogGroups)
			client.printfLog(shard.ShardID, logGroupList)
		}
	}
}

func (client *Client) printfLog(shardId int, logGroupList *sls.LogGroupList) string {
	for _, loggroup := range logGroupList.LogGroups {
		for _, l := range loggroup.Logs {
			length := 0
			data := make(map[string]interface{})
			for _, content := range l.Contents {
				data[content.GetKey()] = content.GetValue()
				length += len(content.GetKey()) + len(content.GetValue())
			}
			for _, tag := range loggroup.LogTags {
				data[tag.GetKey()] = tag.GetValue()
				length += len(tag.GetKey()) + len(tag.GetValue())
			}
			logger.Sugar().Info("shard id ", shardId, data)
		}
	}
	return ""
}

func (client *Client) getTimeOfCursor(shardId int, cursor string) {
	time, err := client.GetCursorTime(client.ProjectName, client.LogStoreName, shardId, cursor)
	if err != nil {
		logger.Sugar().Info("get cursor time err:", err)
	}
	logger.Sugar().Info("get cursor time :", time)
}
