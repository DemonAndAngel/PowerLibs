package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ClientInterface interface {
	Send(request RequestInterface, options *object.HashMap) ResponseContract
	SendAsync(request RequestInterface, options *object.HashMap) PromiseInterface

	Request(method string, uri string, options *object.HashMap, returnRaw bool, outHeader interface{}, outBody interface{}) (ResponseContract,error)
	RequestAsync(method string, uri string, options *object.HashMap, returnRaw bool, outHeader interface{}, outBody interface{})

	SetClientConfig(config *object.HashMap) ClientInterface
	GetClientConfig() *object.HashMap
}
