// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms service
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package msorderms

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The service manager order
type Service interface {
	// SayHello implements SayHello.
	SayHello(context.Context, *SayHelloPayload) (res string, err error)
	// CreateOrder implements CreateOrder.
	CreateOrder(context.Context, *OrdineRequest) (res *StatoOrdine, err error)
	// GetStatusOrderByID implements GetStatusOrderById.
	GetStatusOrderByID(context.Context, *GetStatusOrderByIDPayload) (res *StatoOrdine, err error)
	// SayHello implements SayHello.
	SayHelloEndpoint(context.Context, *SayHelloPayload) (res string, err error)
	// CreateOrder implements CreateOrder.
	CreateOrderEndpoint(context.Context, *OrdineRequest) (res *StatoOrdine, err error)
	// GetStatusOrderByID implements GetStatusOrderById.
	GetStatusOrderByIDEndpoint(context.Context, *GetStatusOrderByIDPayload) (res *StatoOrdine, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "msorderms"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"SayHello", "CreateOrder", "GetStatusOrderById", "SayHello", "CreateOrder", "GetStatusOrderById"}

// GetStatusOrderByIDPayload is the payload type of the msorderms service
// GetStatusOrderById method.
type GetStatusOrderByIDPayload struct {
	// IdOrdine
	IDOrdine string
}

// OrdineRequest is the payload type of the msorderms service CreateOrder
// method.
type OrdineRequest struct {
	// Id Ordine
	IDOrdine string
	// Codice Store
	Store string
	// Data Ordine
	DataOrdine *string
	// Tipologia Ordine
	TipologiaOrdine string
	// Nome Cliente
	NomeCliente *string
	// Cognome Cliente
	CognomeCliente *string
	// Righe Ordine
	RigheOrdine []*RigaOrdine
}

type RigaOrdine struct {
	// Barcode Articolo
	Barcode string
	// Codice Articolo
	Modello *string
	// Codice Colore
	Colore *string
	// Codice Taglia
	Taglia *string
	// Numero pezzi
	Quantita int
	// Prezzo
	Prezzo *float32
}

// SayHelloPayload is the payload type of the msorderms service SayHello method.
type SayHelloPayload struct {
	// Name
	Name string
}

// StatoOrdine is the result type of the msorderms service CreateOrder method.
type StatoOrdine struct {
	// Id Ordine
	OrdineID    *string
	StatoOrdine *string
}

// MakeNoCriteria builds a goa.ServiceError from an error.
func MakeNoCriteria(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "no_criteria", false, false, false)
}

// MakeNoMatch builds a goa.ServiceError from an error.
func MakeNoMatch(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "no_match", false, false, false)
}