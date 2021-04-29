# Zinx TCP服务器框架

## v0.2
### 简单的连接封装和业务绑定
> 连接的模块
- 方法
    - 启动连接
    - 停止连接
    - 获取当前连接的conn对象(socket)
    - 得到连接ID
    - 得到客户端连接的地址和端口
    - 发送数据的方法
- 属性
    - socket TCP套接字
    - 连接的ID
    - 当前连接的状态
    - 与当前连接绑定的处理业务方法
    - 等待连接被动退出的channel

## v0.3

> 基础router模块

- Ruquest请求封装(将链接和数据绑定在一起)
    - 属性
        - 连接IConnection
        - 请求数据
    - 方法
        - 得到当前连接
        - 得到当前数据
- Router模块
    - 抽象的IRouter
        -  处理业务之前的方法
        -  处理业务的方法
        -  处理业务之后的方法
    - 具体的BaseRouter(作为具体实现的基类)
        -  处理业务之前的方法
        -  处理业务的方法
        -  处理业务之后的方法
- zinx集成router模块
    - Iserver增添路由功能
    - Server类增添Router成员
    - Commection类绑定一个Router成员
    - 在Connection调用 已经注册的Router处理业务

## v0.4
> 增添全局配置

## v0.5

> 消息封装

- 定义一个消息的结构
    - 属性
        - 消息的ID
        - 消息长度
        - 消息的内容
    - 方法
        - Setter
        - Getter
- 将消息封装机制集成到Zinx框架中    
    - 将Message添加到Request中
    - 修改连接读取数据的机制 将之前的单纯读取byte改为拆包读取方式
    - 连接的发包机制 将发送的消息进行打包 再发送