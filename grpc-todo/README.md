### 来源
* 【详细】https://esakat.github.io/esakat-blog/posts/vue-go-grpcweb/

* 原作者也是参考: https://qiita.com/otanu/items/98d553d4b685a8419952#docker

* **代码跑不动了，但还是可以参考用**

* 浏览器不能直接调用 gRPC，因为 gRPC 基于 HTTP/2 流式传输，而浏览器不支持裸 gRPC。 所以需要一个中间代理：gRPC-Web。
	* **浏览器不能直接访问 :50051 的 gRPC 服务，需要一个代理转换 gRPC-Web 到 gRPC**。 使用 **`envoy`** 或轻量级的 `grpcwebproxy`
	```shell
	[Web Browser] 
	    ↓ (gRPC-Web)
	[Envoy Proxy 或 grpcwebproxy] 
	    ↓ (gRPC)
	[Go 服务端]
	```
