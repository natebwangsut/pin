# Pin

[![Go Report Card](https://goreportcard.com/badge/github.com/natebwangsut/pin)](https://goreportcard.com/report/github.com/natebwangsut/pin)
[![Go Coverage](https://img.shields.io/badge/go%20coverage-here-brightgreen)](https://gocover.io/github.com/natebwangsut/pin)

> Converting string to PIN number.

## Example

```console
$ pin "hello"
43556

$ pin "hello-world"
4355696753

$ pin "HELLO-WORLD"
4355696753

$ pin -i "hello-world"
43556-96753
```
