# go-architecture-packages-and-commands-in-the-same-repository

## 概要
Goチームが紹介しているパッケージ構成のうちの一つである`packages-and-commands-in-the-same-repository`の実装例を紹介します([参考リンク](https://go.dev/doc/modules/layout))

## 実行方法
- 足し算の場合
```shell
cd cmd/add
go run main.go user pass 3 2
```
- 引き算の場合
```shell
cd cmd/subtract
go run main.go user pass 3 2
```