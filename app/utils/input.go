package utils

import (
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"strconv"
	"time"
)

func getCurrentTime() time.Time {
	currentTime := time.Now()
	return currentTime
}

func getExpiryTime(expiryArg string) string {
	expiryMs, _ := strconv.ParseInt(expiryArg, 10, 64)
	currentTime := getCurrentTime()
	expiryTime := currentTime.Add(time.Duration(expiryMs) * time.Millisecond)
	return expiryTime.Format(time.RFC3339Nano)
}

func InputParser(data string) string {
	inputMap := GetInstance()

	respObj := Resp{}
	tokenList := respObj.DecodeData(data)

	switch tokenList[2] {
	case "ping":
		return "+PONG\r\n"
	case "echo":
		echo := []string{tokenList[4]}
		encodedEcho := respObj.EncodeData(echo)
		return encodedEcho
	case "set":
		ParsedInputObj := ParsedInput{args: tokenList[0], key: tokenList[4], val: tokenList[6]}
		if ParsedInputObj.args == "*5" {
			ParsedInputObj.expiryFlag = tokenList[8]
			ParsedInputObj.expiryTime = getExpiryTime(tokenList[10])
		}
		response := inputMap.InsertData(ParsedInputObj)
		return response
	case "get":
		key := tokenList[4]
		ParsedInputObj := inputMap.GetData(key)
		if ParsedInputObj.nullBulk != "" {
			return ParsedInputObj.nullBulk
		}
		return respObj.EncodeData(ParsedInputObj.toGetArr())
	case "config":
		config := config.GetConfig()
		switch tokenList[4] {
		case "get":
			switch tokenList[6] {
			case "dir":
				configDir := []string{"dir", config.Dir}
				encodedConfigDir := respObj.EncodeData(configDir)
				return encodedConfigDir
			case "dbfilename":
				configDB := []string{"dbfilename", config.DBFilename}
				encodedConfigDB := respObj.EncodeData(configDB)
				return encodedConfigDB
			}
		}
	}

	return string("-1")
}
