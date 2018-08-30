# go-censor [![GoDoc](https://godoc.org/github.com/d4l3k/go-censor?status.svg)](https://godoc.org/github.com/d4l3k/go-censor) [![Build Status](https://travis-ci.com/d4l3k/go-censor.svg?branch=master)](https://travis-ci.com/d4l3k/go-censor)

A go library for censoring bad/offensive words.

```
$ go get -u github.com/d4l3k/go-censor
```

```go
c := censor.New([]string{
  "test",
  "testfoo",
  "bar",
})

c.Pattern = "!@#$"

c.CensorString("foobar test")
// "foo$!@ $!@#"
```
