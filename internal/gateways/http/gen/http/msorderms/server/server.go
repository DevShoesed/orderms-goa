// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms HTTP server
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package server

import (
	"context"
	"net/http"
	msorderms "orderms/internal/gateways/http/gen/msorderms"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the msorderms service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	SayHello           http.Handler
	CreateOrder        http.Handler
	GetStatusOrderByID http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the msorderms service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *msorderms.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"SayHello", "GET", "/sayHello"},
			{"CreateOrder", "POST", "/order"},
			{"GetStatusOrderByID", "GET", "/order/{idOrdine}"},
		},
		SayHello:           NewSayHelloHandler(e.SayHello, mux, decoder, encoder, errhandler, formatter),
		CreateOrder:        NewCreateOrderHandler(e.CreateOrder, mux, decoder, encoder, errhandler, formatter),
		GetStatusOrderByID: NewGetStatusOrderByIDHandler(e.GetStatusOrderByID, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "msorderms" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.SayHello = m(s.SayHello)
	s.CreateOrder = m(s.CreateOrder)
	s.GetStatusOrderByID = m(s.GetStatusOrderByID)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return msorderms.MethodNames[:] }

// Mount configures the mux to serve the msorderms endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSayHelloHandler(mux, h.SayHello)
	MountCreateOrderHandler(mux, h.CreateOrder)
	MountGetStatusOrderByIDHandler(mux, h.GetStatusOrderByID)
}

// Mount configures the mux to serve the msorderms endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountSayHelloHandler configures the mux to serve the "msorderms" service
// "SayHello" endpoint.
func MountSayHelloHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/sayHello", f)
}

// NewSayHelloHandler creates a HTTP handler which loads the HTTP request and
// calls the "msorderms" service "SayHello" endpoint.
func NewSayHelloHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSayHelloRequest(mux, decoder)
		encodeResponse = EncodeSayHelloResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "SayHello")
		ctx = context.WithValue(ctx, goa.ServiceKey, "msorderms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateOrderHandler configures the mux to serve the "msorderms" service
// "CreateOrder" endpoint.
func MountCreateOrderHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/order", f)
}

// NewCreateOrderHandler creates a HTTP handler which loads the HTTP request
// and calls the "msorderms" service "CreateOrder" endpoint.
func NewCreateOrderHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateOrderRequest(mux, decoder)
		encodeResponse = EncodeCreateOrderResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "CreateOrder")
		ctx = context.WithValue(ctx, goa.ServiceKey, "msorderms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetStatusOrderByIDHandler configures the mux to serve the "msorderms"
// service "GetStatusOrderById" endpoint.
func MountGetStatusOrderByIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/order/{idOrdine}", f)
}

// NewGetStatusOrderByIDHandler creates a HTTP handler which loads the HTTP
// request and calls the "msorderms" service "GetStatusOrderById" endpoint.
func NewGetStatusOrderByIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetStatusOrderByIDRequest(mux, decoder)
		encodeResponse = EncodeGetStatusOrderByIDResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetStatusOrderById")
		ctx = context.WithValue(ctx, goa.ServiceKey, "msorderms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}