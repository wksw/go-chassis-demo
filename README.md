> go-chassis框架demo

## 目录结构

```shell
├── bin/                      // 二进制包归档目录
├── conf/                     // 主配置文件存放目录
├── docker-compose.yaml       // onebox环境部署文件
├── Dockerfile                // docker镜像打包文件
├── go.mod                    // 依赖关系
├── go.sum
├── log/                      // 日志存放目录
├── main.go                   // 程序入口
├── system.go                 // 系统初始化
├── version                   // 版本号 
├── Makefile                  // 打包、编译、部署脚本
├── manifests                 // 部署配置文件
│   └── templates/            // 分布式部署配置文件
│       ├── deployment.yaml
│       ├── namespace.yaml
│       ├── secret.yaml
│       └── virtualservice.yaml
```

## 运行

```shell
# 更新代码
make update
# 安装依赖(只需执行一次)
make dependices
# 编译
make build
# 运行
make onebox
# 卸载
make onebox_down
```

## 环境变量

```shell
# 项目名称
export PROJECT_NAME=wksw
# 服务名称
export SERVICE_NAME=go-chassis-demo
# dockerhub推送镜像地址
export REGISTRY_ADDRESS=dockerhub.contoso.com
# dockerhub拉取镜像地址
export REGISTRY_INTERNAL_ADDRESS=dockerhub.contoso.com
# 命名空间
export NAMESPACE=passport
# git仓库路径
export GIT_REPO=github.com
# 部署环境,跟配置文件后缀相关
export DEPLOY_ENV=a
# 编译镜像, 不包含domain的路径和tag
export BUILD_IMAGE=chassis_v2.0.2:build_latest
```