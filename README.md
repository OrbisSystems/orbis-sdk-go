# Orbis Golang SDK

### Orbis SDK is designed to help clients consume Orbis APIs with ease. Orbis SDK takes care of all the heavy lifting and provides out-of-the-box ready-to-use objects.

## Installation
Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

`go mod init`

Then, reference orbis-sdk-go in a Go program with import:

```go
  import "github.com/OrbisSystems/orbis-sdk-go"
```

Go version of the SDK is `1.20`

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the stripe-go module automatically.

Alternatively, you can also explicitly go get the package into a project:

`go get -u github.com/OrbisSystems/orbis-sdk-go`

## Documentation

Below are a few simple examples:

### Orbis config
```go
cfg := config.Config{
    Host:     "host-to-orbis-server", // ask client support. JUST hostname, without setting schema. Example: localhost, NOT https://localhost 
    LogLevel: config.TraceLogLevel,
}
```

### Token storage
You can use redis
```go
tokenStorage, err := storage.NewRedisStorage(storage.Config{
    Host: "localhost",
    Port: 6379,
})
```
or in-memory storage 
```go
tokenStorage := storage.NewInMemoryStorage()
```

### Init AUTH system
```go
authSys := auth.New(tokenStorage)
```

### Set up the client
```go
cli := client.NewSDK(cfg, authSys)
```

### Auth
You can check do you need to loging
```go
needToLogin, err := cli.Account.NeedToLogin(ctx)
```

For login use the next API
```go
cli.Account.LoginByEmail(ctx, model.LoginByEmailRequest{
    Email:      "test@test.com",
    Password:   "passpass",
    DeviceID:   "devideID",
    RememberMe: true,
})
```
Also, you can use API Keys for loging using `LoginByAPIKey` method.

### Client using
After login, you can use one of the many APIs we provide you via this client. 
Here are some examples:

#### News
```go
response, err := cli.News.GetByFilter(ctx, model.NewsFilterRequest{
    Paging: model.Paging{
        Limit:  100,
        Offset: 2,
    },
})
```

#### TipRank
```go
response, err := cli.TipRank.AnalystPortfolios(ctx, model.PortfoliosRequest{
    ExpertUID: "dwdqw31",
    Number:    12,
})
```

#### Logos
```go
response, err := cli.Logos.SymbolLogos(ctx, "AAPL")
```

#### Funds
```go
response, err := cli.Funds.GetFundDetails(ctx, "AAPL")
```

### WebSocket
By using web socket connection you can get news feed from our system
```go
newsCh, err := cli.WS.Subscribe(ctx, model.NewsSubscription)
if err != nil {
    log.Error(err)
    return
}

// newsCh is chan model.News
```

## Example of working code for using Orbis System
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/client"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
)

func main() {
	// creating basic context
	ctx := context.Background()

	// set configs
	cfg := config.Config{
		Host:     "host-to-orbis-server",
		LogLevel: config.InfoLogLevel,
	}

	// init token storage using REDIS
	tokenStorage, err := storage.NewRedisStorage(storage.Config{
		Host: "localhost",
		Port: 6379,
	})
	if err != nil {
		panic(err)
	}

	// init auth system
	authSys := auth.New(tokenStorage)

	// create SDK Client
	cli := client.NewSDK(cfg, authSys)

	// check do we need to call Log In flow
	needToLogin, err := cli.Account.NeedToLogin(ctx)
	if err != nil {
		return
	}

	if needToLogin {
		// if we need to Log In - do it using email and password
		if err := cli.Account.LoginByEmail(ctx, model.LoginByEmailRequest{
			Email:      "some-email@loca.com",
			Password:   "strong_password",
			DeviceID:   "qwqeqweq",
			RememberMe: true,
		}); err != nil {
			fmt.Println("login error")
		}
	}

	// get some latest news
	news, err := cli.News.GetByFilter(ctx, model.NewsFilterRequest{
		Paging: model.Paging{
			Limit: 10,
		},
	})
	fmt.Println(err)
	fmt.Println(news)

	// get some fund details for specific symbol
	fundDetails, err := cli.Funds.GetFundDetails(ctx, "AAPL")
	fmt.Println(err)
	fmt.Println(fundDetails)

	// create web socket subscription for getting news in realtime
	newsCh, err := cli.WS.Subscribe(ctx, model.NewsSubscription)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			// reading news from web socket
			case wsNews := <-newsCh:
				fmt.Println(wsNews)
			}
		}
	}()

	time.Sleep(time.Hour * 48)
}
```

### Check full list of available [API](./API.md)