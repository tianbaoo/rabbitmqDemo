## Go Rabbitmq Demo
### 第一步 docker-compose安装rabbitmq
提前安装好docker、docker-compose(开发调试利器)  
在当前docker-compose.yaml目录下执行以下指令
```bash
docker-compose up -d
```
### 第二步 测试
开启第一个shell窗口
```bash
go run worker.go
```
开启第二个shell窗口
```bash
go run worker.go
```
开启第三个shell窗口
```bash
go run task.go
```

