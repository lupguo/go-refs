## 执行

### 编译
```shell
cd kafkahttpd
go install ./cmd/...
```

### 生产者
执行curl请求httpd服务，通过生产者写入消息到MQ

```shell
curl --location --request POST '0.0.0.0:3000/api/v1/comments' \
--header 'Content-Type: application/json' \
--data-raw '{ "text":"nice boy" }'

curl --location --request POST '0.0.0.0:3000/api/v1/comments' \
--header 'Content-Type: application/json' \
--data-raw '{ "text":"keep up the good work" }'
```

httpd服务接收消息：

```shell
$ kproducer-httpd
...
Message is stored in topic(comments)/partition(1)/offset(70)
Message is stored in topic(comments)/partition(1)/offset(71)
Message is stored in topic(comments)/partition(0)/offset(73)
Message is stored in topic(comments)/partition(0)/offset(74)
```

### 消费者
```shell
$ kconsumer-comments
2022/01/10 15:51:19 Partition(2) | Topic(comments) | Message({"text":"nice boy1"}) | Offset(48)
2022/01/10 15:51:19 Partition(2) | Topic(comments) | Message({"text":"nice boy2"}) | Offset(49)
2022/01/10 15:51:19 Partition(2) | Topic(comments) | Message({"text":"nice boy3"}) | Offset(50)
2022/01/10 15:51:19 Partition(0) | Topic(comments) | Message({"text":"nice boy4"}) | Offset(66)
```

