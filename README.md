# scaling-couscous
goland consul封装
Consul 注册只是“告诉别人我存在”，能不能用，还得靠调用方主动做服务发现和负载均衡。



初始化consul服务
        func InitConsul() {
        config := api.DefaultConfig()
        config.Address = fmt.Sprintf("%s:%d", "you consul ip", 8500)
        client, err := api.NewClient(config)
        if err != nil {
            return
        }
        ConsulClient = client 
       }


1.	实例化工具
      consul := tools.NewConsul("ip", 8500)

   2. 服务过滤 
`      filterConsul, err := consul.FilterConsul("服务名称")
      if err != nil {
      fmt.Println("服务发现失败:", err)
      return
      }`
      3. 如果未找到则注册
           `` `if len(filterConsul) == 0 {
            fmt.Println("service not found consul register service:")

            // gRPC健康检查配置
            grpcCheck := &api.AgentServiceCheck{
            GRPC:                           "ip", // 修改为外部 IP 地址
            Interval:                       "15s",                 // 健康检查间隔
            Timeout:                        "5s",                  // 超时时间
            DeregisterCriticalServiceAfter: "30m",                 // 故障30分钟后注销服务
            }`

            // 注册到 Consul
       err = consul.RegisterConsulWithCheck("serviceName", "ip", port, []string{""}, grpcCheck)
      if err != nil {
            fmt.Println("consul注册失败")
       } else {
            fmt.Println("consul注册成功")
      }
      } else {
            fmt.Println("服务已经注册到consul")
      }``

              services, err := consul.GetConsulServices()
              if err != nil {
              fmt.Println("获取服务失败", err)
              return
              }
              // 打印服务地址和端口
              if len(services) == 0 {
              fmt.Println("没有找到注册的服务")
              } else {
              for _, service := range services {
              fmt.Printf("服务ID: %s, 地址: %s, 端口: %d\n", service.Service, service.Address, service.Port)
              }
              }

