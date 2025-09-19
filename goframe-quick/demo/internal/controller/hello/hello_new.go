// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"demo/api/hello"
)

type ControllerV1 struct{}

// 路由对象: 返回【接口】- TODO: 并非必须, 为了比较严谨的代码编写方式!【个人感觉太繁琐了】

func NewV1() hello.IHelloV1 {
	return &ControllerV1{}
}
