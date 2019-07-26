# telegram-prometheus-bot 

## install

- go build main.go mess.go

## Running

```
prometheus -h

Usage: bot [-port 6666] [-url http://xxxxxx] -h help

Options:        
  -char_id string
        telegram chat id
  -h    print help
  -port string
        listen port
  -url string
        telegram sendMessage url
```

## exemple to test

```
./prometheus -char_id 1234 -port 8888 -url http://123.com
```

## prometheus config

```yaml
route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1m
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://localhost:12345/alert'
```


## exemple alert message

```$xslt
告警类型: resolved
事件类型: InstanceDown
告警主机: 192.168.0.1:8080
主机角色: apiserver
告警摘要: php-fpm: has been down
告警描述: 192.168.0.1:8080: job php-fpm has been down value==0
故障开始时间: 2019-07-26T18:17:13.352482605+08:00
故障恢复时间: 2019-07-26T18:48:13.352482605+08:00
```