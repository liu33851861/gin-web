package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 结构体转为json
func Struct2Json(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}

// json转为结构体
func Json2Struct(str string, obj interface{}) {
	// 将json转为结构体
	_ = json.Unmarshal([]byte(str), obj)
}

// json interface转为结构体
func JsonI2Struct(str interface{}, obj interface{}) {
	// 将json interface转为string
	jsonStr, _ := str.(string)
	Json2Struct(jsonStr, obj)
}

// 结构体转结构体, json为中间桥梁, struct2必须以指针方式传递, 否则可能获取到空数据
func Struct2StructByJson(struct1 interface{}, struct2 interface{}) {
	// 转换为响应结构体, 隐藏部分字段
	jsonStr := Struct2Json(struct1)
	Json2Struct(jsonStr, struct2)
}

// 两结构体比对不同的字段, 不同时将取struct1中的字段返回, json为中间桥梁, struct3必须以指针方式传递, 否则可能获取到空数据
func CompareDifferenceStructByJson(struct1 interface{}, struct2 interface{}, struct3 interface{}) {
	// 通过json先将其转为map集合
	m1 := make(gin.H, 0)
	m2 := make(gin.H, 0)
	m3 := make(gin.H, 0)
	Json2Struct(Struct2Json(struct1), &m1)
	Json2Struct(Struct2Json(struct2), &m2)
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			// key相同, 值不同
			if k1 == k2 && v1 != v2 {
				m3[k1] = v1
			}
		}
	}
	Json2Struct(Struct2Json(m3), struct3)
}