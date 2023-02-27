极简版抖音 v1.0.1

目前实现了core、interact、social所有模块
使用时从http请求中获取相应参数并生成rpc请求

## Quick Start

### 1.Setup Basic Dependence

```shell
docker-compose up
```

### 2.Run Interaction RPC Server

```shell
cd cmd/rpc/interaction
sh build.sh
sh output/bootstrap.sh
```

### 3.Run Message RPC Server

```shell
cd cmd/rpc/message
sh build.sh
sh output/bootstrap.sh
```

### 4.Run Sociality RPC Server

```shell
cd cmd/rpc/sociality
sh build.sh
sh output/bootstrap.sh
```

### 5.Run User RPC Server

```shell
cd cmd/rpc/user
sh build.sh
sh output/bootstrap.sh
```

### 6.Run Video RPC Server

```shell
cd cmd/rpc/video
sh build.sh
sh output/bootstrap.sh
```

### 7.Run API Server

```shell
cd cmd/api
sh run.sh
```