# CLI for Greeting Manager

A CLI application written in Go to actuate local endpoints of Greeting Manager

## Prerequisites

- Go 1.23

## Setup

1. Checkout source from [GitHub](https://github.com/bergerdenes/greeting-cli)
2. Build with

`$ go build`

## Usage

The CLI binary has a self-descriptive help, however here are some examples:

```
$ ./greeting-cli create --name "John Doe" --age 42 --gender MALE --keywords ["it","telecom"]
$ ./greeting-cli list
$ ./greeting-cli get --id 1
$ ./greeting-cli greet --id 1
```

## Skipped items

Configurability - i.e. it always connects to the localhost:8080