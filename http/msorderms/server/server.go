package main

import (
	"context"
	"net/http"
	"orderms/gen/http/msorderms/server"
	"orderms/gen/msorderms"

	goahttp "goa.design/goa/v3/http"
)

type svc struct{}

func (s *svc) SayHello(ctx context.Context, p *msorderms.SayHelloPayload) (string, error) {
	sayHelloReturn := "Hello " + p.Name + "!"
	return sayHelloReturn, nil
}

func (s *svc) CreateOrder(ctx context.Context, p *msorderms.OrdineRequest) (res *msorderms.StatoOrdine, err error) {
	return &msorderms.StatoOrdine{}, nil
}

func (s *svc) GetStatusOrderByID(ctx context.Context, p *msorderms.GetStatusOrderByIDPayload) (res *msorderms.StatoOrdine, err error) {
	return &msorderms.StatoOrdine{}, nil
}

func main() {
	s := &svc{}
	endpoints := msorderms.NewEndpoints(s)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil, nil)
	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:30001",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
