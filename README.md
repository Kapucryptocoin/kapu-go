[![GitHub issues](https://img.shields.io/github/issues/kristjank/ark-net.svg)](https://github.com/kristjank/ark-go/issues)&nbsp;[![GitHub forks](https://img.shields.io/github/forks/kristjank/ark-net.svg)](https://github.com/kristjank/ark-go/network)&nbsp;[![GitHub stars](https://img.shields.io/github/stars/kristjank/ark-net.svg)](https://github.com/kristjank/ark-go/stargazers)&nbsp;[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kristjank/ark-go/master/LICENSE)

## Overview

A short intro about the different directories within this repository:

* [/arkcoin](/kapucoin) - Coin configuration and helper functions
* [/core](/core) - API to use the Kapu blockchain out of Golang
* [/cmd/arkgopool](/cmd/kapugopool) - Client for delegate profit sharing pools
* [/cmd/arkgoserver](/cmd/kapugoserver) - Server for delgates profit sharing pools
* [/raw](/raw) - Images and other raw files

You can find more information about the different functionalities in their folder.

## Why Kapu-GO
GoLang is an open source programming language developed by Google and designed for building fast, simple and reliable software. It is not about theoretical concepts such as monads and virtual inheritance, but more about **hands-on experience**.

Kapu-GO is the KAPU Archealogical Blockchain  library client implemented in GOLANG programming language. It implements all most relevant KAPU/ARK functionalities to help you  **develop efficient, fast and scalable GOLANG applications built upon kAPU platform**. It provides also low level access to KAPU so you can easily build your application on top of it.

## A library demo app: KAPUGO-GUI
A demo client app was developed to test the goark library package dependencies. More about the demo gui client: [/cmd/kapugopool](/cmd/kapugopool). It targets delegate and basic account functionalites, plus silent mode - to be able to run automated reward payments.

See the demo console app shell recording here:
https://asciinema.org/a/5yndxl794ncfpmjoqftuaiodm?t=8.

## How to install?
```
$> go get github.com/kapucoin/kapu-go
```

## How to get dependencies?
```
$> go get ./...
```

## How to get started?
All node services have available reponses have their struct representations. It's best to let the code do the speaking. Every class implementation has it's own test class. **So it's best to start learning by looking at actual test code**.

## Kapu-GO Client Usage
**First call should be network selection, so all settings can initialize from the peers before going into action.**  By default MAINNET is active.

### Init
[GoDoc documentation available on this link](https://godoc.org/github.com/kristjank/ark-go/core)
```go
import "github.com/kapucoin/kapu-go/core"
var KAPUClient = core.NewArkClient(nil)
```

### Communication
Queries to the blockchain are done with the Query struct parameters:

```go
params := TransactionQueryParams{Limit: 10, SenderID: senderID}
```
... and the results -  reponse is also parametrized.
```go
transResponse, _, err := kapuapi.ListTransaction(params)
if transResponse.Success {
		log.Println(t.Name(), "Success, returned", transResponse.Count, "transactions")
	} else {
		t.Error(err.Error())
	}
```

### Other call samples
```go
//usage samples
deleResp, _, _ := kapuclient.GetDelegate(params)

//switch networks
kapuclient = kapuclient.SetActiveConfiguration(core.KAPU) //or core.MAINNET
//create and send tx
kapuclient := NewArkClient(nil)
recepient := "address"
passphrase := "pass"

tx := CreateTransaction(recepient,1,"KAPU-GOLang is saying whoop whooop",passphrase, "")
payload.Transactions = append(payload.Transactions, tx)
res, httpresponse, err := kapuclient.PostTransaction(payload)
```
## More information about KAPU Ecosystem and etc
* [KAPU Ecosystem ](https://github.com/kapucoin

Please, use github issues for questions or feedback. For confidential requests or specific demands, contact us on our public channels.

## Authors
Chris (kristjan.kosic@gmail.com), with a lot of help from FX Thoorens fx@ark.io and ARK Community
Jarunik (jens@hviid.ch), with a lot of help from Chris

# License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Copyright (c) 2017 ARK, chris
