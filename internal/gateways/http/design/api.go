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

var StatoOrdine = ResultType("StatoOrdine", func() {
	Description("A StatoOrdine describe a state of Order.")
	TypeName("StatoOrdine")
	Attributes(func() {
		Attribute("ordineId", String, "Id Ordine")
		Attribute("statoOrdine", func() {
			Enum("da elaborare", "in elaborazione", "presa in carico", "confermato", "annullato")
		})
	})
	View("default", func() {
		Attribute("ordineId")
		Attribute("statoOrdine")
	})

})
