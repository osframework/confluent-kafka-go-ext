# confluent-kafka-go-ext
[![Build Status](https://travis-ci.org/osframework/confluent-kafka-go-ext.svg?branch=master)](https://travis-ci.org/osframework/confluent-kafka-go-ext)

Simple utility extension for easy use of [Confluent's Golang Client for Apache Kafka](https://github.com/confluentinc/confluent-kafka-go).

## Requirements

* Go 1.13+ (for Go Modules support)
* [librdkafka](https://github.com/confluentinc/confluent-kafka-go#librdkafka) (depends on your OS)

## Getting Started

Starting with Go 1.13, you can use [Go Modules] to install **confluent-kafka-go-ext**.

Import the `confluent` package from GitHub in your code:
```go
import "github.com/osframework/confluent-kafka-go-ext/confluent"
```

Build your project:
```bash
go build ./...
```

If you are building for Alpine Linux (musl), `-tags musl` must be specified.
```bash
go build -tags musl ./...
```

A dependency to the latest stable version of confluent-kafka-go-ext should be automatically added to your `go.mod` file.

## Tests

Briefly discuss automated tests here.

## Contributing
Contributions to the code, examples, documentation, etc. are appreciated.

1. Make your changes 
2. run make
3. push your branch
4. create a PR