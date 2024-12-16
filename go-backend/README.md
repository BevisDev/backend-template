# Go Backend

## Getting started

### Prerequisites

- [Go 1.23.3](https://go.dev/dl/) or higher

### Dependencies

- [Gin Gonic](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [Lumberjack](https://github.com/natefinch/lumberjack)
- [Gorm](https://gorm.io/docs/index.html)

### Init Go Module

Create folder project

```sh
$ mkdir go-backend
$ cd go-backend
```

Init Go module

```sh
$ go mod init github.com/BevisDev/go-backend
```

### Getting Framework

```sh
$ go get -u github.com/gin-gonic/gin
```

### Running Gin

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

To run the code, use the `go run` command, like:

```sh
$ go run main.go
```

Then visit [`0.0.0.0:8080/ping`](http://0.0.0.0:8080/ping) in your browser to see the response!

## Getting Viper

```sh
$ go get github.com/spf13/viper
```

## Getting handler Logger

```sh
$ go get -u go.uber.org/zap
```

For writting logs to rolling files

```sh
$ go get github.com/natefinch/lumberjack
```
