package utils

import "go.mongodb.org/mongo-driver/bson"

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
