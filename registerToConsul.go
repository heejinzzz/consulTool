package consulTool

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
)

// Register 将服务注册到consul中
func Register(consulAddress string, serviceID string, serviceName string, serviceTags []string, serviceIP string, servicePort int) {
	// 设置配置，作为客户端连接consul
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	agent := client.Agent()

	// 写要注册的服务信息
	checkAddress := serviceIP + ":" + strconv.Itoa(servicePort)

	registration := api.AgentServiceRegistration{
		ID:      serviceID,      // 服务节点名称
		Name:    serviceName,    // 服务名称
		Tags:    serviceTags,    // 服务的标签
		Address: serviceIP, // 提供服务的IP地址
		Port:    servicePort,    // 提供服务的端口号
		// 对服务的健康检查
		Check: &api.AgentServiceCheck{
			Interval: "30s",
			Timeout:  "2s",
			TCP:      checkAddress,
		},
	}

	// 将该服务注册进consul
	err = agent.ServiceRegister(&registration)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("已将服务：" + serviceID + " 成功注册到consul中！")
}
