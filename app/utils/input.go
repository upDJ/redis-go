package utils

import (
	"strings"
	"sync"
)

var (
	inputMap map[string][]string
	once     sync.Once
)

// getInputMap initializes inputMap only once and returns it.
func getInputMap() map[string][]string {
	once.Do(func() {
		inputMap = make(map[string][]string)
	})
	return inputMap
}

func InputParser(data string) string {
	var ping = string("ping")
	var echo = string("echo")
	var set = string("set")
	var get = string("get")
	var delim = string("\r\n")

	var dataArr = strings.Split(data, delim)
  inputMap := getInputMap()

	if strings.Contains(strings.ToLower(data), ping) {
		return string("+PONG\r\n")

	} else if strings.Contains(strings.ToLower(data), echo) {
		return string("+" + dataArr[4] + delim)

	} else if strings.Contains(strings.ToLower(data), set) {
		key := dataArr[4]
		val := dataArr[5:7]
		inputMap[key] = val
		return string("+OK" + delim)

	} else if strings.Contains(strings.ToLower(data), get) {
		val := inputMap[dataArr[4]]
		return strings.Join(val, delim) + delim

	}

	return string("")
}
