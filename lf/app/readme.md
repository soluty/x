
游戏开发框架

网络层  tcp websocket udp 等等
模块化 
模块间通信  rpc
生命周期管理  kill信号
日志 -> 包括error发送
序列化和反序列化
cmd
配置
协议路由

// ecs 中的world为一个模块
组件  每一个模块由不同组件组成???
组件间通信  getComponent


棋牌游戏  
account
room
player
game   // 明确的round


domain














code organization

packages contain code that has a single purpose

库组织  接口 实现

包组织  库
1 一个package只有一个目的
2 同样目的的package在一起用父子，接口组合

包组织  app

Domain Types  领域对象
Services  operate with Domain

