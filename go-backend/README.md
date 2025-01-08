# Go Backend

## Getting started

### Prerequisites

- [Go 1.23.4](https://go.dev/doc/install) or higher

### Dependencies

- [Gin Gonic](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [Lumberjack](https://github.com/natefinch/lumberjack)
- [Cron](https://github.com/robfig/cron)

### SQL Driver

- [SQL Server](https://github.com/denisenkom/go-mssqldb)

```sh
go get github.com/denisenkom/go-mssqldb
```

- [PostgreSQL](https://github.com/lib/pq)
- [Oracle](https://github.com/godror/godror)
- [Other Driver](https://go.dev/wiki/SQLDrivers)

### Utilities:

- [UUID](https://github.com/google/uuid)
- [Wire](https://github.com/google/wire)

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

On Windows: using **Chocolatey**

Open PowerShell with **Administrator privileges** and run the following command:

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

## Getting Cron

To schedule job runner

```sh
go get github.com/robfig/cron/v3@v3.0.0
```

**Cron Expression Format**

A cron expression represents a set of times, using 6 space-separated fields.

### Example:

```go
c := cron.New(cron.WithSeconds())
// second minute hour day month weekday
c.AddFunc("0 * * * * *", func() { 
    fmt.Println("Running every minute at the 0th second!") 
})
c.Start()
```

| Field name   | Mandatory? | Allowed values  | Allowed special characters |
|--------------|------------|-----------------|----------------------------|
| Seconds      | Yes        | 0-59            | * / , -                    |
| Minutes      | Yes        | 0-59            | * / , -                    |
| Hours        | Yes        | 0-23            | * / , -                    |
| Day of month | Yes        | 1-31            | * / , - ?                  |
| Month        | Yes        | 1-12 or JAN-DEC | * / , -                    |
| Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?                  |

### Getting interact with database

To use map into struct easily:

```sh
go get github.com/jmoiron/sqlx
```

### Getting DI

```sh
go get github.com/google/wire/cmd/wire
```