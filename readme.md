# Task for [adjust.com](https://adjust.com) [![codecov](https://codecov.io/gh/gvencadze/adjust-task/branch/main/graph/badge.svg?token=YVMQC5DVIJ)](https://codecov.io/gh/gvencadze/adjust-task)

## General info
This is simple CLI to send multiple http requests and get md5 hashes of responses.

## Technologies
* Go 1.16.3

## Run
```
git clone github.com/gvencadze/adjust-task
cd /adjust-task/cmd/hash && go build main.go
go run ./main -p 5 vk.com/ac
```