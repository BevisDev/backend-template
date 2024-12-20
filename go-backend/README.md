# Go Backend

## Getting started

### Prerequisites

- [Go 1.23.4](https://go.dev/doc/install) or higher

### Dependencies

- [Gin Gonic](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [Lumberjack](https://github.com/natefinch/lumberjack)
- [Gorm](https://gorm.io/docs/index.html)

## Technology stack

> **Note:**
>
> To switch to a difference enviroment, you need to set the environment variable
>
> On Windows:
>
> ```sh
> setx GO_PROFILE dev
> ```
>
> On Linux:
>
> ```sh
> export GO_PROFILE=dev
> ```

### Getting Makefile Tools

To install `make`

On Windows: using <b>Chocolatey</b>

Open PowerShell with <b>Administrator privileges </b> and run the following command:

```sh
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]::Tls12; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

After Chocolatey is installed, you can install `make` by running the following command in the PowerShell or Command Prompt:

```sh
choco install make
```

On Linux using `apt` to install

```sh
sudo apt update
sudo apt install make
```

### Init Go Module

Create folder project

```sh
mkdir go-backend
cd go-backend
```

Init Go module

```sh
go mod init github.com/BevisDev/go-backend
```

### Getting Framework

```sh
go get -u github.com/gin-gonic/gin
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
go run main.go
```

Then visit [`0.0.0.0:8080/ping`](http://0.0.0.0:8080/ping) in your browser to see the response!

## Getting Viper

```sh
go get github.com/spf13/viper
```

## Getting handler Logger

```sh
go get -u go.uber.org/zap
```

For writting logs to rolling files

```sh
go get github.com/natefinch/lumberjack
```
