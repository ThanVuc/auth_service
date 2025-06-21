package consumers

type brokerModeType string

const (
	BrokerModeDirect  brokerModeType = "direct"
	BrokerModeFanout  brokerModeType = "fanout"
	BrokerModeTopic   brokerModeType = "topic"
	BrokerModeHeaders brokerModeType = "headers"
)

type exchangeNameType string

const (
	ExchangeNameCreateResource exchangeNameType = "direct_create_route_resource"
)

type routingKeyType string

const (
	RoutingKeyCreateResource routingKeyType = "api.resource.route.created"
)

type ConsumerNameType string

const (
	ConsumerNameCreateResource ConsumerNameType = "create_resource_route_consumer"
)
