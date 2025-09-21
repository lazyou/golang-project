### 快速开始
https://goframe.org.cn/quick/install

* 项目初始化:
	```shell
	go mod init quick 
	
	go get -u -v github.com/gogf/gf/v2

	go run main.go
	```

* 引入新包: `go mod tidy`


### 【重要】 - 项目【`gf`脚手架】  
* https://goframe.org.cn/quick/scaffold-index

* 脚手架作用: 工程脚手架(项目初始化)、代码自动生成、工具及框架更新等

* 安装脚手架、初始化项目： `go install github.com/gogf/gf/cmd/gf/v2@latest`
	* `gf init demo -u` 创建goframe最新版本的项目.
    
    * `cd demo && gf run main.go` 运行项目, **开发时修改代码热更新**.

* 脚手架项目的 - 工程目录: https://goframe.org.cn/docs/design/project-structure
	* 入口文件 `main.go` 调用 `internal/cmd` 对应**命令启动**web.
    
    * 核心业务逻辑都是放到了`internal`目录下, **对外隐藏可见性**.

* web 支持优雅关闭: 监听关闭信号(ctrl+c)、不再接受新请求、保证旧请求处理完


### 接口开发
* https://goframe.org.cn/quick/scaffold-api

#### 【重要】**生成dao-do-entity代码**:
* 数据规范-gen dao: https://goframe.org.cn/docs/cli/gen-dao

* 数据库配置 `demo/hack/config.yaml`
  
* 生成dao-do-entity代码: `gf gen dao` || `make dao`

* 很显然挺复杂，一个表生成4个go文件:
  * `/dao/internal/user.go` - DO NOT EDIT! 封装对数据表user的访问. 该文件自动生成了一些数据结构和方法，简化对数据表的CRUD操作. 【生成覆盖】. 
  * `/dao/user.go` - may modify! 对 `dao/internal/user.go` 的进一步封装，用于**供其他模块直接调用访问**. 【可随意修改，或者扩展dao的能力】. 【不会被覆盖】
  * `/model/do/user.go` - DO NOT EDIT! 数据转换模型的使用. 【用来写入或更新时赋值! 避免零值问题】
  * `/model/entity/user.go` - DO NOT EDIT! 数据结构定义与数据表字段. 
  * **但大多不推荐修改, 而是生成覆盖!**

* 【重要】生成的**do 结构体**, 在写入/更新参数、查询条件时会用到
  * 字段全为`any`类型，避免了(如 entity 的结构体)字段其它类型零值可能不写入数据的问题。

* 【重要】生成的**entity 结构体**, 在响应数据时会用到

#### 【路由、请求、响应、数据验证】数据结构定义
* https://goframe.org.cn/quick/scaffold-api-definition

* 控制器文件(**一个控制器一个目录啊！**): `demo/api/user/v1/user.go`

### 生成控制器代码
* 接口规范-gen ctrl: https://goframe.org.cn/docs/cli/gen-ctrl
  * `请求、响应的结构体`命名要规范才能生成控制器

* `gf gen ctrl` 生成文件如下:
	* `/api/user/user.go` - api接口抽象文件 - 保证控制器实现的接口完整性
    * `/controller/user/user.go`: 空的, 管理控制器的初始化.【只生成一次,后随意修改】
    * `/controller/user/user_new.go`: New返回控制器的接口, 管理控制器的初始化.【只生成一次,后随意修改】
    * `/controller/user/user_v1_create.go`: 【实现业务代码】      
    * `/controller/user/user_v1_delete.go`      
    * `/controller/user/user_v1_update.go`      
    * `/controller/user/user_v1_get_one.go`     
    * `/controller/user/user_v1_get_list.go`

* TODO: 也没指定控制器名、路由名或模型名就生成了？？

#### 剩下步骤
* 引入数据库驱动: `_ "github.com/gogf/gf/contrib/drivers/mysql/v2"`

* 配置数据库、日志级别、swagger文档: `manifest/config/config.yaml`
	*  http://127.0.0.1:8000/swagger/

* 路由注册: `user.NewV1()`

* 启动服务: `go run main.go`

* 接口测试: https://goframe.org.cn/quick/scaffold-api-run-and-test
  * curl 命令不记录在这里, 开文档看


### 开发工具 gf: https://goframe.org.cn/docs/cli/install
* 安装: `go install github.com/gogf/gf/cmd/gf/v2@latest`

* 查看: `gf -v`, 输出信息挺多

* 升级: `gf up`

* 【重要】创建项目: `gf init`, 支持大仓、小仓模式!

* 交叉编译: `gf build`, 支持系统架构(`386,amd64,arm`等)
	* 在一种操作系统或 CPU 架构的机器上, 编译出能在另一种操作系统或架构上运行的可执行程序 -- 交叉编译.

* 生成代码 `gf gen`:
	* 接口规范【重要】: `gf gen ctrl` 先写请求、响应结构体, 才能用命令生成控制器!
    * 数据规范【重要】: `gf gen dao` 【最常用】, 支持`分表`等!
	* 模块规范: 不推荐, 就不看.
    * 枚举维护:
    * 协议编译： `gf gen pb` 一般rpc才用!

* 【重要】自动编译: `gf run main.go` 监测代码变化, 开发实用.

* 资源打包: `gf pack`. 将任意的文件打包为资源文件或者 Go 代码文件.

* 【重要】镜像打包: `gf docker`. TODO: 由于Windows下没有docker, **此步骤未尝试**


### 核心组件 
#### 对象管理 - g.Xxx: https://goframe.org.cn/docs/core/g
* 见单测文件`g_test.go`, 运行某个单测 `go test -v -run Test_gMap`

* 数据类型、HTTP 客户端对象、Validator 校验对象、(单例) 日志管理对象、(单例) 配置管理对象、(单例) 数据库 ORM 对象、(单例) Redis 客户端对象 等等!

* 【重要】**理解 `context.WithTimeout` 超时上下文的在 curl 请求中的使用!**

* 其中单例的对象, 是按传入 `name` 来算单例!

#### 调试模式 - https://goframe.org.cn/docs/core/debugging
* 框架内部开发者在开发阶段使用: `g.SetDebug(true)`, 日志会相当的多! 

* 也支持`命令行启动参数 + 环境变量`运行: 如 `go run .\main.go --gf.debug=true`

#### 命令管理 - https://goframe.org.cn/docs/core/gcmd
* 对比 https://github.com/spf13/cobra
	* 优点: "支持链路跟踪，便于父子进程的链路信息传递"

* 【重要】**使用场景: 同一个入口, 封装不同的命令启动不同的服务(如 `api`, `queue`, `crontab`) 等! 方便容器打包启动不同的命令!**

* 单元测试: `gcmd_test.go`

* 【重要】基本概念: 
  * **参数(Argument): 按照顺序进行传递、没有名称标识的数据**.
  * **选项(Option): 以 - 或者 -- 字符串作为前缀, 无序, 可缩写`别名(Short)`**. 类似`标识(Flag)`.

  * `gcmd.GetOptWithEnv`: 如果该选项不存在时，则从`环境变量`中读取

* 命令解析器: `gcmd.Parser` 自定义解析选项，包括有哪些选项名称，每个选项是否带有数值

* 命令行对象: `gcmd.Command` 管理单个或多个命令, 支持嵌套

* 结构化参数: `g.Meta` 结构化`管理参数`

* 终端交互: `gcmd.Scan|Scanf`

* 跨进程的链路跟踪: "gproc.ShellRun(ctx, `go run sub.go`)"

