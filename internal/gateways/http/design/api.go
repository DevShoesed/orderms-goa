package design

import . "goa.design/goa/v3/dsl"

var _ = API("msorderms", func() {
	Title("Order Mangament Service")
	Description("Microservice for manager Order")

	Server("msorderms", func() {
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:30001")
			//URI("grpc://localhost:30002")

		})
	})
})

var _ = Service("msorderms", func() {
	Description("The service manager order")

	Method("SayHello", func() {
		Payload(func() {
			Field(1, "name", String, "Name")
			Required("name")
		})
		Result(String)

		HTTP(func() {
			GET("/sayHello")
			Param("name")
		})

		// GRPC(func ()  {})
	})

	Method("CreateOrder", func() {
		Payload(CreateOrderPayload)
		Result(StatoOrdine)

		HTTP(func() {
			POST("/order")
		})
	})

	Method("GetStatusOrderById", func() {
		Payload(func() {
			Field(1, "idOrdine", String, "IdOrdine")
			Required("idOrdine")
		})
		Result(StatoOrdine)

		HTTP(func() {
			GET("/order/{idOrdine}")
		})
	})
})

var RigaOrdine = Type("RigaOrdine", func() {
	Attribute("barcode", String, "Barcode Articolo")
	Attribute("modello", String, "Codice Articolo")
	Attribute("colore", String, "Codice Colore")
	Attribute("taglia", String, "Codice Taglia")
	Attribute("quantita", Int, "Numero pezzi")
	Attribute("prezzo", Float32, "Prezzo")

	Required("barcode", "quantita")
})

var CreateOrderPayload = Type("OrdineRequest", func() {
	Attribute("idOrdine", String, "Id Ordine")
	Attribute("store", String, "Codice Store")
	Attribute("dataOrdine", String, "Data Ordine")
	Attribute("tipologiaOrdine", String, "Tipologia Ordine", func() {
		Enum("corriere", "negozio")
	})
	Attribute("nomeCliente", String, "Nome Cliente")
	Attribute("cognomeCliente", String, "Cognome Cliente")
	Attribute("righeOrdine", ArrayOf(RigaOrdine), "Righe Ordine")

	Required("idOrdine", "store", "tipologiaOrdine")
})

var StatoOrdine = Type("StatoOrdine", func() {
	Attribute("ordineId", String, "Id Ordine")
	Attribute("statoOrdine", func() {
		Enum("da elaborare", "in elaborazione", "presa in carico", "confermato", "annullato")
	})
})
