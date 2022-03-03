package main

import (
	"encoding/json"
	"flag"
	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var ConfigPath string
var ZookeeperHost string
var CONSUMER_PATH = "/transform/consumers"

type (
	Config struct {
		Topic    string
		Consumer string
		Offsets  map[int64]int64
	}
	ConfigEx1 struct {
		Topic    string
		Consumer string
		Offsets  []map[string]int64
	}
	ConfigEx2 struct {
		Topic    string
		Consumer string
		Offsets  []int64
	}
	ZkNodeData struct {
		Topic       string `json:"topic"`
		PartitionId int64  `json:"partitionId"`
		Offset      int64  `json:"offset"`
	}
)

func (c *ConfigEx1) convertToConfig() *Config {
	if c.Topic == "" || c.Offsets == nil || len(c.Offsets) == 0 {
		return nil
	}

	conf := &Config{
		Topic:    c.Topic,
		Consumer: c.Consumer,
		Offsets:  make(map[int64]int64, 0),
	}

	for _, m := range c.Offsets {
		idx := int64(0)
		if v, ok := m["partition"]; !ok {
			log.Fatalf("config invalid: %+v", m)
		} else {
			idx = v
		}

		if v, ok := m["offset"]; !ok {
			log.Fatalf("config invlaid: %+v", m)
		} else {
			conf.Offsets[idx] = v
		}
	}

	return conf
}

func (c *ConfigEx2) convertToConfig() *Config {
	if c.Topic == "" || c.Offsets == nil || len(c.Offsets) == 0 {
		return nil
	}

	conf := &Config{
		Topic:    c.Topic,
		Consumer: c.Consumer,
		Offsets:  make(map[int64]int64, 0),
	}

	for idx, v := range c.Offsets {
		conf.Offsets[int64(idx)] = v
	}

	return conf
}

func parseYamlFromLocal(configPath string) *Config {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("read config file failed")
		return nil
	}
log.Printf(string(data))
	var configEx1 ConfigEx1
	if err := yaml.Unmarshal(data, &configEx1); err == nil {
	log.Printf("%v",configEx1)
		return configEx1.convertToConfig()
	}
log.Printf(string(data))
	var configEx2 ConfigEx2
	if err := yaml.Unmarshal(data, &configEx2); err == nil {
		log.Printf("%v",configEx2)
		return configEx2.convertToConfig()
	}

	log.Fatal("parse config failed")
	return nil
}

func init() {
	flag.StringVar(&ConfigPath, "c", "config.yml", "please set config path with -c.")
	flag.StringVar(&ZookeeperHost, "h", "127.0.0.1:2181", "please set zookeeper host with -h.")
}

func main() {
	flag.Parse()
	// 读取配置
	c := parseYamlFromLocal(ConfigPath)
	log.Printf("config: %+v\n", c)

	// 连接zk
	conn, _, err := zk.Connect([]string{ZookeeperHost}, time.Second*5)
	defer conn.Close()

	time.Sleep(time.Second * time.Duration(2))

	if err != nil {
		log.Fatal(err)
	}

	// 扫描目录
	partitions, _, err := conn.Children(CONSUMER_PATH + "/" + c.Consumer)
	if err != nil {
		log.Fatalf("fetch partitions childrens error: %+v", err)
	}

	if len(partitions) > 0 {
		for _, partition := range partitions {
			partitionPath := CONSUMER_PATH + "/" + c.Consumer + "/" + partition
			data, stat, err := conn.Get(partitionPath)
			if err != nil {
				log.Fatalf("fetch data error: %+v", err)
			}

			var zkNodeData ZkNodeData
			err = json.Unmarshal(data, &zkNodeData)
			if err != nil {
				log.Printf("node data invalid: %+v\n", string(data))
				continue
			}

			if zkNodeData.Topic == c.Topic {
				for k, v := range c.Offsets {
					if zkNodeData.PartitionId == k {
						preV := zkNodeData.Offset
						zkNodeData.Offset = v
						b, _ := json.Marshal(zkNodeData)
						_, err = conn.Set(partitionPath, b, stat.Version)
						if err != nil {
							log.Fatalf("set partition failed: %+v", err)
						}

						log.Printf("modify success with %s, offset from %d -> %d\n", partitionPath, preV, v)
					}
				}
			}
		}
	}

	log.Println("task finish!")
}

