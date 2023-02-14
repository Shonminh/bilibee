# bilibee
 哔哩哔哩（Bilibili B站）字幕分析工具

## 介绍
打算先做成只可以获取B站指定UP主的视频信息，包括字幕，然后将其同步到ES中，使用ES进行数据分析，NLP等。

## 数据流图
大致思路

![](assets/pic/bilibee_data_flow_chart.png)

## 如何启动？

### 安装golang
> go >= 1.18

### 启动Mysql
```bash
mysql.server start
```

### 创建表结构
> sql的内容位于deploy/sqls/db.sql中，可以使用客户端工具执行。

### 配置环境变量
```bash
export DB_USERNAME=YOUR_USERNAME
export DB_PASSWORD=YOUR_PWD
```

### 生成wire_gen.go（go依赖注入）
```bash
bash wire_gen.sh
```

### 启动http服务
```bash
go run ./cmd/collect_server
```
对应的http服务会监听8080端口

### 创建一条cron task
```bash
curl --location --request POST '127.0.0.1:8080/api/cron_task/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mid": 333
}'
```
目前http服务还没有做前端页面，待开发，创建完成后会在数据库cron_task_tab看到记录
![img.png](assets/pic/img.png)

### 启动获取视频信息的定时任务
```bash
go run ./cmd/collect_task
```
执行的过程中，相应的日志位于`log/data.log`中，可以观察对应的处理过程，正常处理完成后，task_status会置为1，表示完成。
此时，可以看到video_info_tab中记录video info信息：
![img2.png](assets/pic/img2.png)


### 

## 进展
- 2023.02.04 封装了获取b站视频信息的接口。
- 2023.02.15 完成获取视频信息服务的开发 && 自测。
