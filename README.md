# consulTool
## introduction
Tool to register and unregister services with consul. 

将服务注册到consul和将服务从consul中注销的工具。

## requirement
If you import "github.com/heejinzzz/consulTool" in your code, you may need to execute this command first:

	go get go-micro.dev/v4
Then execute:

	go mod tidy

若你在你的代码中 import 了 "github.com/heejinzzz/consulTool" ，你可能需要先执行：

	go get go-micro.dev/v4
再执行：

	go mod tidy

## api
	Register(consulAddress string, serviceID string, serviceName string, serviceTags []string, serviceIP string, servicePort int)
Register the service to the specified consul.

将服务注册到指定consul中。

----------------


	Unregister(consulAddress string, serviceID string)
Unregister the specified service from the consul.

在consul中注销指定的服务。

----------------
	
	GetServerAddress(consulAddress string, serviceID string)
According to the serviceID, go to the specified consul server to obtain the server address of the service.

根据serviceID，到指定的consul服务器上获取该服务的服务端地址。
