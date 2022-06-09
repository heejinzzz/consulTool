package consulTool

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// Unregister 将服务从consul中注销
func Unregister(consulAddress string, serviceID string) {
	// 设置配置，作为客户端连接consul
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	agent := client.Agent()

	// 在consul中注销指定服务
	err = agent.ServiceDeregister(serviceID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("已将服务：" + serviceID + " 从consul中取消注册！")
}
