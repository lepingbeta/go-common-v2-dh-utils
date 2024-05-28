/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 08:09:45
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-28 17:01:04
 * @FilePath     : /v2/go-common-v2-dh-utils/utils.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsElementInSlice[T comparable](element T, slice []T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Struct2BsonD(doc interface{}) (bson.D, error) {
	// 将结构体编码为BSON字节序列
	data, err := bson.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// 将BSON字节序列解码为bson.D
	var bsonDoc bson.D
	err = bson.Unmarshal(data, &bsonDoc)
	if err != nil {
		return nil, err
	}

	return bsonDoc, nil
}

func Struct2BsonM(doc interface{}) (bson.M, error) {
	// 将结构体编码为BSON字节序列
	data, err := bson.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// 将BSON字节序列解码为bson.D
	var bsonDoc bson.M
	err = bson.Unmarshal(data, &bsonDoc)
	if err != nil {
		return nil, err
	}

	return bsonDoc, nil
}

func ObjectIDFromHex(s string) primitive.ObjectID {
	objId, _ := primitive.ObjectIDFromHex(s)
	return objId
}

// FilterBsonM 函数接受原始 bson.M 数据和要保留的字段列表，
// 返回一个新的 bson.M 只包含指定的字段。
// 示例	keepFields := []string{"name", "email"}
func FilterBsonM(data bson.M, keepFields []string) bson.M {
	filteredData := bson.M{}
	for _, key := range keepFields {
		if value, ok := data[key]; ok {
			filteredData[key] = value
		}
	}
	return filteredData
}
