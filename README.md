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