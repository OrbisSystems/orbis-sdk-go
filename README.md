# Orbis Golang SDK

### Orbis SDK is designed to help clients consume Orbis APIs with ease. Orbis SDK takes care of all the heavy lifting and provides out of the box ready to use objects.

## Installation
Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

`go mod init`

Then, reference orbis-sdk-go in a Go program with import:

```go
  import "github.com/OrbisSystems/orbis-sdk-go"
```

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the stripe-go module automatically.

Alternatively, you can also explicitly go get the package into a project:

`go get -u github.com/OrbisSystems/orbis-sdk-go`

## Documentation

Below are a few simple examples:

### Set up 

```go
var orbisConfig = config.OrbisConfig{ApiUrl: "url"}
var authStorage = storage.NewInMemory()
var authSrv = auth.NewAuth(authStorage, "key")
var configuredClient = http.NewClient(10 * time.Second, authSrv)
```

### Orbis config
```go
var orbisConfig = config.OrbisConfig{ApiUrl: "url"}
```

### Auth Storage
> Provides an interface to store the auth token

You can use default implementation `InMemory`

Also, you are able to create your one just implements the interface

```go
type Storage interface {
    Save(key string, value []byte) error
    Get(key string) ([]byte, bool, error)
}
```

Example of initialization default implementation 
```go
var inMemory = storage.NewInMemory()
```

### Auth Service
> Provides wrapper for authentication system

Example of usage:
```go
var authService = auth.NewAuth(authStorage, http.DefaultAuthKey)
```

**http.DefaultAuthKey** is a second parameter
that setting key for your token in auth storage

### Refresh token
Omni-SDK provides your ability for automation refreshing token

```go
var account = account.NewClient(orbisConfig)
var chanErr = make(chan error, 1)
go accountService.Refresh(chanErr)
```
**accountService.Refresh** will automatically refresh your token

### WebSocket connection

```go
var msgChan = make(chan []byte, 1)
var wsClient, err = ws.NewWebSocketClient(ws.WSS, ws.StreamingOrdersHost, msgChan, config.OrbisConfig{WSHost: "WSHost"}, authSrv)
if err != nil {
log.Fatal("couldn't initialize ws client", " err:", err)
}
go wsClient.Connect()
```
