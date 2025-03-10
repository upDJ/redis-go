package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	inputMap map[string]map[string]string
	once     sync.Once
)

// {'foo': {"length": "","val": "", "expiryFlag": "", "expiry": ""}}

// getInputMap initializes inputMap only once and returns it.
func getInputMap() map[string]map[string]string {
	once.Do(func() {
		inputMap = make(map[string]map[string]string)
	})
	return inputMap
}

func getCurrentTime() time.Time {
	currentTime := time.Now()
	return currentTime
}

func InputParser(data string) string {
	var ping = string("ping")
	var echo = string("echo")
	var set = string("set")
	var get = string("get")
	var delim = string("\r\n")
	var nullBulk = string("$-1\r\n")

	var dataArr = strings.Split(data, delim)
	args := dataArr[0]
	fmt.Println(dataArr, data)
	inputMap := getInputMap()

	if strings.Contains(strings.ToLower(data), ping) {
		return string("+PONG\r\n")

	} else if strings.Contains(strings.ToLower(data), echo) {
		return string("+" + dataArr[4] + delim)

	} else if strings.Contains(strings.ToLower(data), set) {
		key := dataArr[4]
		if _, exists := inputMap[key]; !exists {
			inputMap[key] = make(map[string]string)
		}
		valLength := dataArr[5]
		val := dataArr[6]

		if args == "*5" {
			expiryFlag := dataArr[8]
			expiryMs, _ := strconv.ParseInt(dataArr[10], 10, 64)
			currentTime := getCurrentTime()
			expiryTime := currentTime.Add(time.Duration(expiryMs) * time.Millisecond)
			expiryTimeStr := expiryTime.Format(time.RFC3339Nano)
			inputMap[key]["expiryFlag"] = expiryFlag
			inputMap[key]["expiry"] = expiryTimeStr
		}

		inputMap[key]["valLength"] = valLength
		inputMap[key]["val"] = val

		return string("+OK" + delim)

	} else if strings.Contains(strings.ToLower(data), get) {
		key := dataArr[4]
		valDict := inputMap[key]
		val := valDict["val"]
    valLength := valDict["valLength"]

		if expiry, exists := valDict["expiry"]; exists {
			currentTime := getCurrentTime()
			if parsedExpiryTime, err := time.Parse(time.RFC3339Nano, expiry); err == nil {
				if parsedExpiryTime.Before(currentTime) {
          fmt.Println(parsedExpiryTime, currentTime)
          fmt.Println("Return Null Bulk", nullBulk)
					return nullBulk
				}
			}
		}
		return valLength + delim + val + delim
	}

	return string("")
}
