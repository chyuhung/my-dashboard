# my-dashboard

## 项目结构说明

```stylus
my-dashboard/
├── api/
│   ├── client.go
│   ├── flavors.go
│   ├── images.go
│   ├── instances.go
│   └── volumes.go
├── handlers/
│   ├── flavors_handler.go
│   ├── images_handler.go
│   ├── instances_handler.go
│   └── volumes_handler.go
├── models/
│   ├── flavor.go
│   ├── image.go
│   ├── instance.go
│   └── volume.go
├── templates/
│   ├── flavors.html
│   ├── images.html
│   ├── instances.html
│   └── volumes.html
├── main.go
├── config.go
└── README.md
```
api/ 目录包含与OpenStack API交互的逻辑，例如创建Gophercloud客户端和调用各种API的方法。

handlers/ 目录包含用于处理HTTP请求的处理程序，从api包获取数据并将其渲染到模板中。

models/ 目录包含表示OpenStack资源（如flavor、image、instance、volume）的模型结构。

templates/ 目录包含HTML模板文件，用于渲染不同的页面和资源。

main.go 是项目的入口文件，用于初始化路由和服务器的设置。

config.go 包含项目的配置信息，例如OpenStack认证参数和其他相关配置。

README.md 是项目的说明文档，可以包含有关项目的详细信息和使用说明。

## 项目运行

### 安装依赖

```bash
go get github.com/gophercloud/gophercloud
```
### 运行项目

```bash
go run main.go
```
## 项目说明

### 项目功能

本项目主要实现OpenStack Dashboard的flavors、images、instances、volumes功能。