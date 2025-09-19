// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package httpx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Context 自定义Context, 扩展 gin.Context
type Context struct {
	*gin.Context // 【继承】gin.Context
}

// 【在编译时进行类型断言检查】，确保 *Context 类型实现了 Go 标准库中的 context.Context 接口
// 等价于 var _ context.Context = (*Context)(nil), 【更推荐这么写】
var _ context.Context = &Context{}

// NewContext 创建自定义Context
func NewContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}
