// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms endpoints
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package msorderms

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "msorderms" service endpoints.
type Endpoints struct {
	SayHello                   goa.Endpoint
	CreateOrder                goa.Endpoint
	GetStatusOrderByID         goa.Endpoint
	SayHelloEndpoint           goa.Endpoint
	CreateOrderEndpoint        goa.Endpoint
	GetStatusOrderByIDEndpoint goa.Endpoint
}

// NewEndpoints wraps the methods of the "msorderms" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		SayHello:                   NewSayHelloEndpoint(s),
		CreateOrder:                NewCreateOrderEndpoint(s),
		GetStatusOrderByID:         NewGetStatusOrderByIDEndpoint(s),
		SayHelloEndpoint:           NewSayHelloEndpointEndpoint(s),
		CreateOrderEndpoint:        NewCreateOrderEndpointEndpoint(s),
		GetStatusOrderByIDEndpoint: NewGetStatusOrderByIDEndpointEndpoint(s),
	}
}

// Use applies the given middleware to all the "msorderms" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.SayHello = m(e.SayHello)
	e.CreateOrder = m(e.CreateOrder)
	e.GetStatusOrderByID = m(e.GetStatusOrderByID)
	e.SayHelloEndpoint = m(e.SayHelloEndpoint)
	e.CreateOrderEndpoint = m(e.CreateOrderEndpoint)
	e.GetStatusOrderByIDEndpoint = m(e.GetStatusOrderByIDEndpoint)
}

// NewSayHelloEndpoint returns an endpoint function that calls the method
// "SayHello" of service "msorderms".
func NewSayHelloEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SayHelloPayload)
		return s.SayHello(ctx, p)
	}
}

// NewCreateOrderEndpoint returns an endpoint function that calls the method
// "CreateOrder" of service "msorderms".
func NewCreateOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*OrdineRequest)
		return s.CreateOrder(ctx, p)
	}
}

// NewGetStatusOrderByIDEndpoint returns an endpoint function that calls the
// method "GetStatusOrderById" of service "msorderms".
func NewGetStatusOrderByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetStatusOrderByIDPayload)
		return s.GetStatusOrderByID(ctx, p)
	}
}

// NewSayHelloEndpointEndpoint returns an endpoint function that calls the
// method "SayHello" of service "msorderms".
func NewSayHelloEndpointEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SayHelloPayload)
		return s.SayHelloEndpoint(ctx, p)
	}
}

// NewCreateOrderEndpointEndpoint returns an endpoint function that calls the
// method "CreateOrder" of service "msorderms".
func NewCreateOrderEndpointEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*OrdineRequest)
		return s.CreateOrderEndpoint(ctx, p)
	}
}

// NewGetStatusOrderByIDEndpointEndpoint returns an endpoint function that
// calls the method "GetStatusOrderById" of service "msorderms".
func NewGetStatusOrderByIDEndpointEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetStatusOrderByIDPayload)
		return s.GetStatusOrderByIDEndpoint(ctx, p)
	}
}