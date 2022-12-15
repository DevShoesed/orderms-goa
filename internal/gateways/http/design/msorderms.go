package design

import (
	. "goa.design/goa/v3/dsl"
)

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
		Error("no_criteria")

		HTTP(func() {
			POST("/order")
			Response(StatusCreated)
			Response("no_criteria", StatusBadRequest)
		})
	})

	Method("GetStatusOrderById", func() {
		Payload(func() {
			Field(1, "idOrdine", String, "IdOrdine")
			Required("idOrdine")
		})
		Result(StatoOrdine)
		Error("no_match")

		HTTP(func() {
			GET("/order/{idOrdine}")
			Response(StatusOK)
			Response("no_match", StatusNotFound)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
