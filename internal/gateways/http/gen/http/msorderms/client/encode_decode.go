// Code generated by goa v3.10.2, DO NOT EDIT.
//
// msorderms HTTP client encoders and decoders
//
// Command:
// $ goa gen orderms/internal/gateways/http/design --output
// internal/gateways/http

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	msorderms "orderms/internal/gateways/http/gen/msorderms"

	goahttp "goa.design/goa/v3/http"
)

// BuildSayHelloRequest instantiates a HTTP request object with method and path
// set to call the "msorderms" service "SayHello" endpoint
func (c *Client) BuildSayHelloRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SayHelloMsordermsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("msorderms", "SayHello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSayHelloRequest returns an encoder for requests sent to the msorderms
// SayHello server.
func EncodeSayHelloRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*msorderms.SayHelloPayload)
		if !ok {
			return goahttp.ErrInvalidType("msorderms", "SayHello", "*msorderms.SayHelloPayload", v)
		}
		values := req.URL.Query()
		values.Add("name", p.Name)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSayHelloResponse returns a decoder for responses returned by the
// msorderms SayHello endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeSayHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("msorderms", "SayHello", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("msorderms", "SayHello", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateOrderRequest instantiates a HTTP request object with method and
// path set to call the "msorderms" service "CreateOrder" endpoint
func (c *Client) BuildCreateOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateOrderMsordermsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("msorderms", "CreateOrder", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateOrderRequest returns an encoder for requests sent to the
// msorderms CreateOrder server.
func EncodeCreateOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*msorderms.OrdineRequest)
		if !ok {
			return goahttp.ErrInvalidType("msorderms", "CreateOrder", "*msorderms.OrdineRequest", v)
		}
		body := NewCreateOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("msorderms", "CreateOrder", err)
		}
		return nil
	}
}

// DecodeCreateOrderResponse returns a decoder for responses returned by the
// msorderms CreateOrder endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCreateOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("msorderms", "CreateOrder", err)
			}
			err = ValidateCreateOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("msorderms", "CreateOrder", err)
			}
			res := NewCreateOrderStatoOrdineCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("msorderms", "CreateOrder", resp.StatusCode, string(body))
		}
	}
}

// BuildGetStatusOrderByIDRequest instantiates a HTTP request object with
// method and path set to call the "msorderms" service "GetStatusOrderById"
// endpoint
func (c *Client) BuildGetStatusOrderByIDRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		idOrdine string
	)
	{
		p, ok := v.(*msorderms.GetStatusOrderByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("msorderms", "GetStatusOrderById", "*msorderms.GetStatusOrderByIDPayload", v)
		}
		idOrdine = p.IDOrdine
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetStatusOrderByIDMsordermsPath(idOrdine)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("msorderms", "GetStatusOrderById", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetStatusOrderByIDResponse returns a decoder for responses returned by
// the msorderms GetStatusOrderById endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetStatusOrderByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetStatusOrderByIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("msorderms", "GetStatusOrderById", err)
			}
			err = ValidateGetStatusOrderByIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("msorderms", "GetStatusOrderById", err)
			}
			res := NewGetStatusOrderByIDStatoOrdineOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("msorderms", "GetStatusOrderById", resp.StatusCode, string(body))
		}
	}
}

// marshalMsordermsRigaOrdineToRigaOrdineRequestBody builds a value of type
// *RigaOrdineRequestBody from a value of type *msorderms.RigaOrdine.
func marshalMsordermsRigaOrdineToRigaOrdineRequestBody(v *msorderms.RigaOrdine) *RigaOrdineRequestBody {
	if v == nil {
		return nil
	}
	res := &RigaOrdineRequestBody{
		Barcode:  v.Barcode,
		Modello:  v.Modello,
		Colore:   v.Colore,
		Taglia:   v.Taglia,
		Quantita: v.Quantita,
		Prezzo:   v.Prezzo,
	}

	return res
}

// marshalRigaOrdineRequestBodyToMsordermsRigaOrdine builds a value of type
// *msorderms.RigaOrdine from a value of type *RigaOrdineRequestBody.
func marshalRigaOrdineRequestBodyToMsordermsRigaOrdine(v *RigaOrdineRequestBody) *msorderms.RigaOrdine {
	if v == nil {
		return nil
	}
	res := &msorderms.RigaOrdine{
		Barcode:  v.Barcode,
		Modello:  v.Modello,
		Colore:   v.Colore,
		Taglia:   v.Taglia,
		Quantita: v.Quantita,
		Prezzo:   v.Prezzo,
	}

	return res
}