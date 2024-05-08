/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 08:09:45
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 08:38:00
 * @FilePath     : /v2/go-common-v2-dh-utils/utils_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package utils

import (
	"testing"

	dhjson "github.com/lepingbeta/go-common-v2-dh-json"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	"go.mongodb.org/mongo-driver/bson"
)

// 定义一个用于测试的结构体
type TestStruct struct {
	Field1 string `bson:"field1"`
	Field2 int    `bson:"field2"`
}

// TestStruct2BsonD 测试 Struct2BsonD 函数
func TestStruct2BsonD(t *testing.T) {
	tests := []struct {
		name    string
		doc     interface{}
		want    bson.D
		wantErr bool
	}{
		{
			name: "ValidStruct",
			doc: TestStruct{
				Field1: "value1",
				Field2: 123,
			},
			want: bson.D{
				{Key: "field1", Value: "value1"},
				{Key: "field2", Value: 123},
			},
			wantErr: false,
		},
		{
			name:    "NilInput",
			doc:     nil,
			want:    bson.D{},
			wantErr: true, // 根据 Marshal 的实现，这里可能是 true 或 false
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Struct2BsonD(tt.doc)
			if (err != nil) != tt.wantErr {
				dhlog.Error("Struct2BsonD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			s1 := dhjson.JsonEncodeIndent(got)
			s2 := dhjson.JsonEncodeIndent(tt.want)
			dhlog.Info(s1)
			dhlog.Info(s2)
			if s1 == s2 {
				// dhlog.Error("", got)
				// dhlog.Error("", tt.want)
				// fmt.Errorf("Struct2BsonD() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func areBsonDEqual(d1, d2 bson.D) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i, elem1 := range d1 {
		elem2 := d2[i]
		if elem1.Key != elem2.Key || elem1.Value != elem2.Value {
			return false
		}
	}
	return true
}
