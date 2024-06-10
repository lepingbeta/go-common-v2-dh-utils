/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 08:09:45
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-11 00:13:10
 * @FilePath     : /v2/go-common-v2-dh-utils/utils.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"runtime"
)

func IsElementInSlice[T comparable](element T, slice []T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

// 生成随机的Access ID，长度为n
func GenerateAccessID(n int) (string, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes), nil
}

// 生成安全的Access Secret
func GenerateAccessSecret() (string, error) {
	bytes := make([]byte, 32) // 32字节的随机数
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

var DebugMsgFlag = true

func DebugMsg(msg string) string {
	if DebugMsgFlag {
		pc, file, line, _ := runtime.Caller(1)

		s := "\n"

		funcName := fmt.Sprintf("Caller function: %s"+s, runtime.FuncForPC(pc).Name())
		fileName := fmt.Sprintf("Caller file: %s"+s, file)
		LineNum := fmt.Sprintf("Caller line: %d"+s, line)

		msg = fmt.Sprintf(s+"%s%s%s%s", funcName, fileName, LineNum, msg)
	}

	return msg
}
