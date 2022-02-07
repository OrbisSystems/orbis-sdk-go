package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Send pings to peer with this period
const pingPeriod = 30 * time.Second

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan []byte
	msgChan   chan []byte
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
}

// NewWebSocketClient create new websocket connection
func NewWebSocketClient(scheme Scheme,
	channel Channel,
	msgChan chan []byte,
	orbisConfig config.OrbisConfig,
	auth auth.API) (*WebSocketClient, error) {
	conn := WebSocketClient{
		sendBuf: make(chan []byte, 1),
		msgChan: msgChan,
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	u := url.URL{
		Scheme: string(scheme),
		Host:   orbisConfig.WSHost,
		Path:   fmt.Sprintf(string(channel), token),
	}
	conn.configStr = u.String()

	go conn.listen()
	go conn.listenWrite()
	go conn.ping()
	return &conn, nil
}

func (conn *WebSocketClient) Connect() *websocket.Conn {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, _, err := websocket.DefaultDialer.Dial(conn.configStr, nil)
			if err != nil {
				log.Println("Cannot connect to websocket: ", conn.configStr, "err:", err)
				continue
			}
			log.Println("connected to websocket to ", conn.configStr)
			conn.wsconn = ws
			return conn.wsconn
		}
	}
}

func (conn *WebSocketClient) listen() {
	log.Println("listen for the messages: ", conn.configStr)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			for {
				ws := conn.Connect()
				if ws == nil {
					return
				}
				_, bytMsg, err := ws.ReadMessage()
				if err != nil {
					log.Println("listen:", err, "Cannot read websocket message")
					conn.closeWs()
					break
				}
				conn.msgChan <- bytMsg
			}
		}
	}
}

// Write data to the websocket server
func (conn *WebSocketClient) Write(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	for {
		select {
		case conn.sendBuf <- data:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

func (conn *WebSocketClient) listenWrite() {
	for data := range conn.sendBuf {
		ws := conn.Connect()
		if ws == nil {
			err := fmt.Errorf("conn.ws is nil")
			log.Println("listenWrite:", err, "No websocket connection")
			continue
		}

		if err := ws.WriteMessage(
			websocket.TextMessage,
			data,
		); err != nil {
			log.Println("listenWrite", nil, "WebSocket Write Error")
		}
		log.Println("listenWrite", nil, fmt.Sprintf("send: %s", data))
	}
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	conn.ctxCancel()
	conn.closeWs()
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	conn.mu.Lock()
	if conn.wsconn != nil {
		conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.wsconn.Close()
		conn.wsconn = nil
	}
	conn.mu.Unlock()
}

func (conn *WebSocketClient) ping() {
	log.Println("ping", nil, "ping pong started")
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ws := conn.Connect()
			if ws == nil {
				continue
			}
			if err := conn.wsconn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod/2)); err != nil {
				conn.closeWs()
			}
		case <-conn.ctx.Done():
			return
		}
	}
}
