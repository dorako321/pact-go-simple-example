# 使い方


consumerの実行方法
契約（pact file）を作成する

```
cd pact-consumer/
go test -v
```

providerの実行方法
契約（pact file）に基づいてテストを実行する

```
cd pact-provider/
go test -v
```