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
