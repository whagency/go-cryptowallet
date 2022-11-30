# Go cryptowallet API client

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/whagency/go-cryptowallet/master/LICENSE.md)
[![Build Status](https://travis-ci.org/whagency/go-cryptowallet.svg?branch=master)](https://travis-ci.org/whagency/go-cryptowallet)
[![Go Report Card](https://goreportcard.com/badge/github.com/whagency/go-cryptowallet?)](https://goreportcard.com/report/github.com/whagency/go-cryptowallet)
[![GoDoc](https://godoc.org/github.com/whagency/go-cryptowallet?status.svg)](https://godoc.org/github.com/whagency/go-cryptowallet)

## Install

```
go get -u github.com/whagency/go-cryptowallet
```

## Getting started

```
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

```
clientData, err := client.GetCurrencies()
if err != nil {
    panic(err)
}

for _, item := range clientData {
    fmt.Println(item)
}
```

###### Get tokens list

```
clientData, err := client.GetTokens("ETH_TOKEN")
if err != nil {
    panic(err)
}

for _, item := range clientData {
    fmt.Println(item)
}
```

###### Add new coin address

```
clientData, err := client.AddCoinAddress("BTC")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Add new token address

```
clientData, err := client.AddTokenAddress("ETH_TOKEN", "USDT", "0xdac17f958d2ee523a2206206994597c13d831ec7")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get balance by coin address

```
clientData, err := client.GetBalanceByCoinAddress("BTC", "3QvvGMcdr942zruPxM4mWiuQW4wkCQG4UG")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get balance by token address

```
clientData, err := client.GetBalanceByTokenAddress("ETH_TOKEN", "0x99e80ef931487b08f12e7c249bf2ce4e9177819c", "USDT", "0xdac17f958d2ee523a2206206994597c13d831ec7")
if err != nil {
    panic(err)
}

fmt.Println(clientData)
```

###### Get balance by token address

*Arguments* : `page, fromTime, operationType`

```
clientData, err := client.GetTransactions(0, 1654690672035, cw.OperationTypeIn)
if err != nil {
    panic(err)
}

fmt.Println(clientData.TotalPages, clientData.TotalElements, clientData.Size, clientData.Page)

for _, item := range clientData.Transactions {
    fmt.Println(item)
}
```