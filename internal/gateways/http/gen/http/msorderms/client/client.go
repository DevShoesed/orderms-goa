// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms client HTTP transport
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the msorderms service endpoint HTTP clients.
type Client struct {
	// SayHello Doer is the HTTP client used to make requests to the sayHello
	// endpoint.
	SayHelloDoer goahttp.Doer

	// GetStatusOrderByID Doer is the HTTP client used to make requests to the
	// getStatusOrderById endpoint.
	GetStatusOrderByIDDoer goahttp.Doer

	// CreateOrder Doer is the HTTP client used to make requests to the createOrder
	// endpoint.
	CreateOrderDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the msorderms service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SayHelloDoer:           doer,
		GetStatusOrderByIDDoer: doer,
		CreateOrderDoer:        doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// SayHello returns an endpoint that makes HTTP requests to the msorderms
// service sayHello server.
func (c *Client) SayHello() goa.Endpoint {
	var (
		encodeRequest  = EncodeSayHelloRequest(c.encoder)
		decodeResponse = DecodeSayHelloResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSayHelloRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SayHelloDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("msorderms", "sayHello", err)
		}
		return decodeResponse(resp)
	}
}

// GetStatusOrderByID returns an endpoint that makes HTTP requests to the
// msorderms service getStatusOrderById server.
func (c *Client) GetStatusOrderByID() goa.Endpoint {
	var (
		decodeResponse = DecodeGetStatusOrderByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetStatusOrderByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetStatusOrderByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("msorderms", "getStatusOrderById", err)
		}
		return decodeResponse(resp)
	}
}

// CreateOrder returns an endpoint that makes HTTP requests to the msorderms
// service createOrder server.
func (c *Client) CreateOrder() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateOrderRequest(c.encoder)
		decodeResponse = DecodeCreateOrderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateOrderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateOrderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("msorderms", "createOrder", err)
		}
		return decodeResponse(resp)
	}
}
