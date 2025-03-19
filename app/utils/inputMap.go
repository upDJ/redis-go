package utils

import (
	"strconv"
	"sync"
	"time"
)

var (
	instance *InputMapSingleton
	once     sync.Once
)

func GetInstance() *InputMapSingleton {
	once.Do(func() {
		instance = &InputMapSingleton{
			data: make(map[string]map[string]string),
		}
	})
	return instance
}

// {'foo': {"length": "","val": "", "expiryFlag": "", "expiry": ""}}
type InputMapSingleton struct {
	data map[string]map[string]string
}

func (i *InputMapSingleton) InsertData(p ParsedInput) string {
	if _, exists := i.data[p.key]; !exists {
		i.data[p.key] = make(map[string]string)
	}

	if p.args == "*5" {
		i.data[p.key]["expiryFlag"] = p.expiryFlag
		i.data[p.key]["expiry"] = p.expiryTime
	}

	valLength := strconv.Itoa(len(p.val))
	i.data[p.key]["valLength"] = valLength
	i.data[p.key]["val"] = p.val
	return "+OK\r\n"
}

func (i *InputMapSingleton) GetData(key string) ParsedInput {
	var p ParsedInput
	valDict := i.data[key]

	if expiry, exists := valDict["expiry"]; exists {
		currentTime := getCurrentTime()
		if parsedExpiryTime, err := time.Parse(time.RFC3339Nano, expiry); err == nil {
			if currentTime.After(parsedExpiryTime) {
				p = ParsedInput{nullBulk: "$-1\r\n"}
				return p
			}
		}
	}

	p = ParsedInput{val: valDict["val"], valLength: valDict["valLength"]}
	return p
}
