/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 08:09:45
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-04 23:20:22
 * @FilePath     : /v2/go-common-v2-dh-utils/utils_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package utils

import (
	"encoding/base64"
	"strings"
	"testing"
)

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// 测试GenerateAccessID函数
func TestGenerateAccessID(t *testing.T) {
	const testIDLength = 16
	expectedLength := testIDLength
	actualID, err := GenerateAccessID(expectedLength)
	if err != nil {
		t.Errorf("GenerateAccessID returned an error: %v", err)
	}
	if len(actualID) != expectedLength {
		t.Errorf("Expected ID length of %d, but got %d", expectedLength, len(actualID))
	}
	// 检查生成的Access ID是否只包含字母数字字符
	for _, char := range actualID {
		if !strings.ContainsRune(alphanum, char) {
			t.Errorf("GenerateAccessID contains invalid characters: %c", char)
		}
	}
}

// 测试GenerateAccessSecret函数
func TestGenerateAccessSecret(t *testing.T) {
	actualSecret, err := GenerateAccessSecret()
	if err != nil {
		t.Errorf("GenerateAccessSecret returned an error: %v", err)
	}
	// 检查编码后的字符串是否是base64格式
	// Base64编码的字符串长度是4的倍数
	if len(actualSecret)%4 != 0 {
		t.Errorf("Expected base64 encoded string length to be a multiple of 4, but got %d", len(actualSecret))
	}
	// 检查解码后的字节长度是否为32
	decodedBytes, err := base64.URLEncoding.DecodeString(actualSecret)
	if err != nil {
		t.Errorf("Failed to decode base64 string: %v", err)
	}
	if len(decodedBytes) != 32 {
		t.Errorf("Expected decoded byte length of 32, but got %d", len(decodedBytes))
	}
}

// BenchmarkGenerateAccessID 基准测试GenerateAccessID函数
func BenchmarkGenerateAccessID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateAccessID(16)
		if err != nil {
			b.Errorf("GenerateAccessID error: %v", err)
		}
	}
}

// BenchmarkGenerateAccessSecret 基准测试GenerateAccessSecret函数
func BenchmarkGenerateAccessSecret(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateAccessSecret()
		if err != nil {
			b.Errorf("GenerateAccessSecret error: %v", err)
		}
	}
}
