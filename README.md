极简版抖音 v1.0.0

目前实现了core和interact模块
使用时从http请求中获取相应参数并生成rpc请求

## Quick Start

### 1.Setup Basic Dependence

```shell
docker-compose up
```

### 2.Run Core RPC Server

```shell
cd rpc/core
sh build.sh
sh output/bootstrap.sh
```

### 3.Run Interact RPC Server

```shell
cd rpc/interact
sh build.sh
sh output/bootstrap.sh
```

### 4.Run API Server

```shell
cd api
./run.sh
```