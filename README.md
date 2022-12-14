# Go cryptowallet API client

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/whagency/go-cryptowallet/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/whagency/go-cryptowallet?v2)](https://goreportcard.com/report/github.com/whagency/go-cryptowallet)
[![GoDoc](https://godoc.org/github.com/whagency/go-cryptowallet?status.svg)](https://godoc.org/github.com/whagency/go-cryptowallet)

## Install

```
go get -u github.com/whagency/go-cryptowallet
```

## Getting started

```go
import (
    cw "github.com/whagency/go-cryptowallet/v1"
)

client := cw.NewClient(&cw.Config{
    ApiKey: "API_KEY",
    ApiUrl: "https://my-wallet-url.com",
    Test: false,
    Logging: cw.LoggerStdout,
})
```

###### Available loggers

| Logger        | Destination                  |
|---------------|------------------------------|
| LoggerStdout  | Log to Stdout                |
| LoggerFile    | Log to file ./log/wallet.log |
| LoggerOff     | Logger disable               |

## Examples

###### Get currencies list

```go
clientData, err := client.GetCurrencies()
if err != nil {
    panic(err)
}

for _, item := range clientData {
    fmt.Println(item)
}
```

###### Get tokens list

```go
clientData, err := client.GetTokens("ETH_TOKEN")
if err != nil {
    panic(err)
}

for _, item := range clientData {
    fmt.Println(item)
}
```

###### Add new coin address

```go
clientData, err := client.AddCoinAddress("BTC")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Add new token address

```go
clientData, err := client.AddTokenAddress("ETH_TOKEN", "USDT", "0xdac17f958d2ee523a2206206994597c13d831ec7")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get balance by coin address

```go
clientData, err := client.GetBalanceByCoinAddress("BTC", "3QvvGMcdr942zruPxM4mWiuQW4wkCQG4UG")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get balance by token address

```go
clientData, err := client.GetBalanceByTokenAddress("ETH_TOKEN", "0x99e80ef931487b08f12e7c249bf2ce4e9177819c", "USDT", "0xdac17f958d2ee523a2206206994597c13d831ec7")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get transactions list

*Arguments* : `page, fromTime, operationType`

```go
clientData, err := client.GetTransactions(0, 1654690672035, cw.OperationTypeIn)
if err != nil {
    panic(err)
}

fmt.Println(clientData.TotalPages, clientData.TotalElements, clientData.Size, clientData.Page)

for _, item := range clientData.Transactions {
    fmt.Println(item)
}
```
