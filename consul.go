package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

var ConsulClient *api.Client

type Consul struct {
	Host string
	Port int
}

func NewConsul(host string, port int) *Consul {

	return &Consul{
		Host: host,
		Port: port,
	}
}

// RegisterConsulWithCheck 注册服务并带健康检查
func (c *Consul) RegisterConsulWithCheck(name, address string, port int, tags []string, check *api.AgentServiceCheck) error {
	if tags == nil {
		tags = []string{}
	}

	registration := &api.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    name,
		Address: address,
		Port:    port,
		Tags:    tags,
		Check:   check,
	}

	err := ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}
	return nil
}

// 服务过滤 Filtration
func (c *Consul) FilterConsul(name string) (map[string]*api.AgentService, error) {

	sprintf := fmt.Sprintf(`Service == "%s"`, name)
	filter, err := ConsulClient.Agent().ServicesWithFilter(sprintf)
	if err != nil {
		return nil, err
	}
	return filter, err
}

// 获取consul服务列表
func (c *Consul) GetConsulServices() (map[string]*api.AgentService, error) {
	services, err := ConsulClient.Agent().Services()
	if err != nil {
		return nil, err
	}
	return services, nil
}

// consul注销
func (c *Consul) ServiceDeregister(serviceID string) error {
	err := ConsulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		return err
	}
	return nil
}
