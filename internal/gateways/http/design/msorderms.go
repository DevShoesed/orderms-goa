package design

import . "goa.design/goa/v3/dsl"

var _ = Service("msorderms", func() {
	Description("The service manager order")
	HTTP(func() {
		Path("/api/v4/msorderms")
	})

	Method("sayHello", func() {
		Payload(func() {
			Field(1, "name", String, "Name")
			Required("name")
		})
		Result(String)
		HTTP(func() {
			GET("/sayHello")
			Param("name")
			Response(StatusOK)
		})
		// GRPC(func ()  {})
	})

	Method("getStatusOrderById", func() {
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

	Method("createOrder", func() {
		Payload(CreateOrderPayload)
		Result(StatoOrdine)
		Error("no_criteria", String, "Missing criteria")
		HTTP(func() {
			POST("/order")
			Response(StatusCreated)
			Response("no_criteria", StatusBadRequest)
		})
		// GRPC(func() {
		// 	Response(CodeOK)
		// 	Response("no_criteria", CodeInvalidArgument)
		// 	Response("no_match", CodeNotFound)
		// })
	})

})
