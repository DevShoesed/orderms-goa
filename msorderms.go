package msorderms

import (
	"context"
	"log"
	"orderms/internal/gateways/http/gen/msorderms"
)

type msordermsSvc struct {
	logger *log.Logger
}

// NewMsorderms return the msorderms service implementation
func NewMsorderms(logger *log.Logger) msorderms.Service {
	return &msordermsSvc{logger}
}

func (s *msordermsSvc) SayHello(ctx context.Context, p *msorderms.SayHelloPayload) (res string, err error) {
	if p.Name == "" {
		return "", msorderms.NoCriteria("Name è obbligatorio")
	}

	res = "Hello, " + p.Name
	return
}

func (s *msordermsSvc) CreateOrder(ctx context.Context, p *msorderms.OrdineRequest) (res *msorderms.StatoOrdine, err error) {
	if p.IDOrdine == "" || p.Store == "" || p.TipologiaOrdine == "" {
		return nil, msorderms.NoCriteria("idOrdine, store, tipologiaOrdine sono obbligatori")
	}

	s.logger.Print("order create")
	return
}

func (s *msordermsSvc) GetStatusOrderByID(ctx context.Context, p *msorderms.GetStatusOrderByIDPayload) (res *msorderms.StatoOrdine, err error) {

	s.logger.Print("get status order by Id")

	return
}
