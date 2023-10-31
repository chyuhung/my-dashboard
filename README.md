# 项目名称

my-dashboard

# 项目简介

my-dashboard 是一个使用 Gophercloud 实现类似 OpenStack Dashboard 的功能的仪表板应用。它提供了对 OpenStack 云环境的资源管理和操作。

# 功能特性

- 用户认证与授权
- 虚拟机实例的管理（创建、启动、停止、删除等）
- 虚拟机实例的监控信息展示
- 网络和存储资源的管理
...
# 技术栈

- Golang
- Gin
- Gophercloud
- Vue
...

# 快速开始
## 后端

克隆仓库：git clone https://github.com/chyuhung/my-dashboard.git

进入后端目录：cd my-dashboard/backend

安装依赖：go mod download

修改配置：在 config/config.go 文件中填入正确的 OpenStack 认证信息和其他配置项

启动后端服务：go run main.go

## 前端

克隆仓库：git clone https://github.com/chyuhung/my-dashboard.git

进入前端目录：cd my-dashboard/frontend

安装依赖：npm install

修改配置：在 .env.development 文件中填入正确的后端服务地址

启动前端服务：npm run serve
