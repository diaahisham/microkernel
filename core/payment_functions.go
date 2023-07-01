package core

import (
	"github.com/gin-gonic/gin"
	"microkernel/register"
)

func Pay(context *gin.Context) {
	// get the payment model from request
	var body PaymentRequestModel
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(
			400,
			"bad request body",
		)
		return
	}
	// get the payment plugin from the register
	// payment plugin must implement the interface PaymentPluginInterface
	paymentPlugin, err := register.PaymentRegister(body.PaymentMethod)
	if err != nil {
		context.JSON(
			400,
			err.Error(),
		)
		return
	}
	// run the pay function
	result, err := paymentPlugin.Pay(body.Payload)
	if err != nil {
		context.JSON(
			400,
			err.Error(),
		)
		return
	}
	// return the success message that is returned from pay function
	context.JSON(
		200,
		result,
	)
	return
}

func Refund(context *gin.Context) {
	// get the payment model from request
	var body PaymentRequestModel
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(
			400,
			"bad request body",
		)
		return
	}
	// get the payment plugin from the register
	// payment plugin must implement the interface PaymentPluginInterface
	paymentPlugin, err := register.PaymentRegister(body.PaymentMethod)
	if err != nil {
		context.JSON(
			400,
			err.Error(),
		)
		return
	}
	// run the refund function
	result, err := paymentPlugin.Refund(body.Payload)
	if err != nil {
		context.JSON(
			400,
			err.Error(),
		)
		return
	}
	// return the success message that is returned from refund function
	context.JSON(
		200,
		result,
	)
	return
}
