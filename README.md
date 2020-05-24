# go-helper

[![Go Report Card](https://goreportcard.com/badge/github.com/ronzxy/go-helper)](https://goreportcard.com/report/github.com/ronzxy/go-helper)
[![GoDoc](https://godoc.org/github.com/ronzxy/go-helper?status.svg)](https://godoc.org/github.com/ronzxy/go-helper)
[![Github All Releases](https://img.shields.io/github/downloads/ronzxy/go-helper/total.svg)](https://github.com/ronzxy/go-helper/releases)
[![GitHub release](https://img.shields.io/github/release/ronzxy/go-helper/all.svg)](https://github.com/ronzxy/go-helper/releases)
[![GitHub Release Date](https://img.shields.io/github/release-date-pre/ronzxy/go-helper.svg)](https://github.com/ronzxy/go-helper/releases)
[![GitHub license](https://img.shields.io/github/license/ronzxy/go-helper.svg)](https://github.com/ronzxy/go-helper/blob/master/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/ronzxy/go-helper.svg)](https://github.com/ronzxy/go-helper/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/ronzxy/go-helper.svg)](https://github.com/ronzxy/go-helper/network)
[![Sourcegraph](https://sourcegraph.com/github.com/ronzxy/go-helper/-/badge.svg)](https://sourcegraph.com/github.com/ronzxy/go-helper?badge)

A help package for golang, provide some utilities, encryption and database orm struct.


## helper package

#### ByteHelper

ByteHelper provides Hex and Byte format conversion.

#### ConsoleColorHelper

ConsoleColorHelper provides console color support.

#### FileHelper

FileHelper provides file copy and gzip compression support.

#### HttpHelper

HttpHelper provides http GET, POST and PostForm support.

#### PathHelper

#### RandHelper

RandHelper provides rand specified number of numbers or strings.

#### StringHelper

#### TimeHelper

TimeHelper provides full format of time format support.

## algo package

A golang crypto package extended, including encryption and padding.

encryption：

 * AES
 * DES
 * TripleDES
 * HMac
 * SHA
 
Padding：
 
 * ZeroPadding
 * PKCS5Padding

## orm package

Database xorm struct, including:

* PostgreSQL
* MySQL
* Redis
* SQLite
