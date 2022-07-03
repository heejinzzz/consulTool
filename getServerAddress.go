package consulTool

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

// GetServerAddress 根据serviceID，去指定的consul服务器上获取该服务的服务端地址
func GetServerAddress(consulAddress string, serviceID string) string {
	// 指定服务所注册到的consul的地址，之后会从该consul中取得提供其所需服务的服务端的地址
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulAddress}
	})

	// 获取提供该服务的所有服务端地址
	serverList, err := consulReg.GetService(serviceID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 设置选择器，选择一个服务端，将其地址返回
	next := selector.RoundRobin(serverList)
	serverNode, err := next()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return serverNode.Address
}
