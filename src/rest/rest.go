package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPIWithHandler(address string) error {
	// Get gin's default engine
	r := gin.Default()
	// Define a handler
	h, _ := NewHandler()
	// get products
	r.GET("/products", h.GetProducts)
	// get promos
	r.GET("/promos", h.GetPromos)
	/*
		// post user sign in
		r.POST("/users/signin", h.SignIn)
	// post user sign out
		In the path below, our relative url needs to include the user id
		Since the id will differ based on the user, the Gin framework allows us to include a wildecard. 
		In Gin, the wildcard wil take the form ':id' to indicate that we are expecting a parameter here 
		with the name 'id'
	r.POST("/user/:id/signout", h.SignOut)
	// get user orders
	r.GET("/user/:id/orders", h.GetOrders) 
	// post purchase charge
	r.POST("/users/charge", h.Charge) 
	*/

	userGroup := r.Group("/user") 
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users") 
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	// run the server
	return r.Run(address)
}
