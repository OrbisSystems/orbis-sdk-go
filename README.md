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

### Check full list of available [API](./API.md)