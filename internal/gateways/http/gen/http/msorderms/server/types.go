// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms HTTP server types
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package server

import (
	msorderms "orderms/internal/gateways/http/gen/msorderms"
	msordermsviews "orderms/internal/gateways/http/gen/msorderms/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateOrderRequestBody is the type of the "msorderms" service "createOrder"
// endpoint HTTP request body.
type CreateOrderRequestBody struct {
	// Id Ordine
	IDOrdine *string `form:"idOrdine,omitempty" json:"idOrdine,omitempty" xml:"idOrdine,omitempty"`
	// Codice Store
	Store *string `form:"store,omitempty" json:"store,omitempty" xml:"store,omitempty"`
	// Data Ordine
	DataOrdine *string `form:"dataOrdine,omitempty" json:"dataOrdine,omitempty" xml:"dataOrdine,omitempty"`
	// Tipologia Ordine
	TipologiaOrdine *string `form:"tipologiaOrdine,omitempty" json:"tipologiaOrdine,omitempty" xml:"tipologiaOrdine,omitempty"`
	// Nome Cliente
	NomeCliente *string `form:"nomeCliente,omitempty" json:"nomeCliente,omitempty" xml:"nomeCliente,omitempty"`
	// Cognome Cliente
	CognomeCliente *string `form:"cognomeCliente,omitempty" json:"cognomeCliente,omitempty" xml:"cognomeCliente,omitempty"`
	// Righe Ordine
	RigheOrdine []*RigaOrdineRequestBody `form:"righeOrdine,omitempty" json:"righeOrdine,omitempty" xml:"righeOrdine,omitempty"`
}

// GetStatusOrderByIDResponseBody is the type of the "msorderms" service
// "getStatusOrderById" endpoint HTTP response body.
type GetStatusOrderByIDResponseBody struct {
	// Id Ordine
	OrdineID    *string `form:"ordineId,omitempty" json:"ordineId,omitempty" xml:"ordineId,omitempty"`
	StatoOrdine *string `form:"statoOrdine,omitempty" json:"statoOrdine,omitempty" xml:"statoOrdine,omitempty"`
}

// CreateOrderResponseBody is the type of the "msorderms" service "createOrder"
// endpoint HTTP response body.
type CreateOrderResponseBody struct {
	// Id Ordine
	OrdineID    *string `form:"ordineId,omitempty" json:"ordineId,omitempty" xml:"ordineId,omitempty"`
	StatoOrdine *string `form:"statoOrdine,omitempty" json:"statoOrdine,omitempty" xml:"statoOrdine,omitempty"`
}

// GetStatusOrderByIDNoMatchResponseBody is the type of the "msorderms" service
// "getStatusOrderById" endpoint HTTP response body for the "no_match" error.
type GetStatusOrderByIDNoMatchResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RigaOrdineRequestBody is used to define fields on request body types.
type RigaOrdineRequestBody struct {
	// Barcode Articolo
	Barcode *string `form:"barcode,omitempty" json:"barcode,omitempty" xml:"barcode,omitempty"`
	// Codice Articolo
	Modello *string `form:"modello,omitempty" json:"modello,omitempty" xml:"modello,omitempty"`
	// Codice Colore
	Colore *string `form:"colore,omitempty" json:"colore,omitempty" xml:"colore,omitempty"`
	// Codice Taglia
	Taglia *string `form:"taglia,omitempty" json:"taglia,omitempty" xml:"taglia,omitempty"`
	// Numero pezzi
	Quantita *int `form:"quantita,omitempty" json:"quantita,omitempty" xml:"quantita,omitempty"`
	// Prezzo
	Prezzo *float32 `form:"prezzo,omitempty" json:"prezzo,omitempty" xml:"prezzo,omitempty"`
}

// NewGetStatusOrderByIDResponseBody builds the HTTP response body from the
// result of the "getStatusOrderById" endpoint of the "msorderms" service.
func NewGetStatusOrderByIDResponseBody(res *msordermsviews.StatoOrdineView) *GetStatusOrderByIDResponseBody {
	body := &GetStatusOrderByIDResponseBody{
		OrdineID:    res.OrdineID,
		StatoOrdine: res.StatoOrdine,
	}
	return body
}

// NewCreateOrderResponseBody builds the HTTP response body from the result of
// the "createOrder" endpoint of the "msorderms" service.
func NewCreateOrderResponseBody(res *msordermsviews.StatoOrdineView) *CreateOrderResponseBody {
	body := &CreateOrderResponseBody{
		OrdineID:    res.OrdineID,
		StatoOrdine: res.StatoOrdine,
	}
	return body
}

// NewGetStatusOrderByIDNoMatchResponseBody builds the HTTP response body from
// the result of the "getStatusOrderById" endpoint of the "msorderms" service.
func NewGetStatusOrderByIDNoMatchResponseBody(res *goa.ServiceError) *GetStatusOrderByIDNoMatchResponseBody {
	body := &GetStatusOrderByIDNoMatchResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSayHelloPayload builds a msorderms service sayHello endpoint payload.
func NewSayHelloPayload(name string) *msorderms.SayHelloPayload {
	v := &msorderms.SayHelloPayload{}
	v.Name = name

	return v
}

// NewGetStatusOrderByIDPayload builds a msorderms service getStatusOrderById
// endpoint payload.
func NewGetStatusOrderByIDPayload(idOrdine string) *msorderms.GetStatusOrderByIDPayload {
	v := &msorderms.GetStatusOrderByIDPayload{}
	v.IDOrdine = idOrdine

	return v
}

// NewCreateOrderOrdineRequest builds a msorderms service createOrder endpoint
// payload.
func NewCreateOrderOrdineRequest(body *CreateOrderRequestBody) *msorderms.OrdineRequest {
	v := &msorderms.OrdineRequest{
		IDOrdine:        *body.IDOrdine,
		Store:           *body.Store,
		DataOrdine:      body.DataOrdine,
		TipologiaOrdine: *body.TipologiaOrdine,
		NomeCliente:     body.NomeCliente,
		CognomeCliente:  body.CognomeCliente,
	}
	if body.RigheOrdine != nil {
		v.RigheOrdine = make([]*msorderms.RigaOrdine, len(body.RigheOrdine))
		for i, val := range body.RigheOrdine {
			v.RigheOrdine[i] = unmarshalRigaOrdineRequestBodyToMsordermsRigaOrdine(val)
		}
	}

	return v
}

// ValidateCreateOrderRequestBody runs the validations defined on
// CreateOrderRequestBody
func ValidateCreateOrderRequestBody(body *CreateOrderRequestBody) (err error) {
	if body.IDOrdine == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("idOrdine", "body"))
	}
	if body.Store == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("store", "body"))
	}
	if body.TipologiaOrdine == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tipologiaOrdine", "body"))
	}
	if body.TipologiaOrdine != nil {
		if !(*body.TipologiaOrdine == "corriere" || *body.TipologiaOrdine == "negozio") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.tipologiaOrdine", *body.TipologiaOrdine, []interface{}{"corriere", "negozio"}))
		}
	}
	for _, e := range body.RigheOrdine {
		if e != nil {
			if err2 := ValidateRigaOrdineRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateRigaOrdineRequestBody runs the validations defined on
// RigaOrdineRequestBody
func ValidateRigaOrdineRequestBody(body *RigaOrdineRequestBody) (err error) {
	if body.Barcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("barcode", "body"))
	}
	if body.Quantita == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("quantita", "body"))
	}
	return
}
