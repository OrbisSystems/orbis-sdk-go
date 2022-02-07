package ws

const (
	wssScheme = "wss"
	wsScheme  = "ws"

	channelStreamingQuotesHostBSON
	channelStreamingQuotesHostJSON
	channelStreamingEquitiesHost
	channelStreamingOrdersHost
	channelStreamingTransferUpdatesHost
	channelStreamingTransferMessengerHost
)

type Scheme string

var (
	WSS Scheme = "wss"
	WS  Scheme = "ws"
)

type Channel string

var (
	StreamingQuotesHostBSON        Channel = "/stream?auth=%s"
	StreamingQuotesHostJSON        Channel = "/stream/json?auth=%s"
	StreamingEquitiesHost          Channel = "/equities?auth=%s"
	StreamingOrdersHost            Channel = "/orders/feed?auth=%s"
	StreamingTransferUpdatesHost   Channel = "/transfers/feed?auth=%s"
	StreamingTransferMessengerHost Channel = "/messenger?auth=%s"
)
