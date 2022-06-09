# consulTool
Tool to register and unregister services with consul. 

将服务注册到consul和将服务从consul中注销的工具。

	Register(consulAddress string, serviceID string, serviceName string, serviceTags []string, serviceIP string, servicePort int)
Register the service to the specified consul.

将服务注册到指定consul中。


	Unregister(consulAddress string, serviceID string)
Unregister the specified service from the consul.

在consul中注销指定的服务。
	
