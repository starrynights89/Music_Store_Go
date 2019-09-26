package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"dblayer"
	"models"
	"github.com/gin-gonic/gin"
)

// Interface for polymorphism 
type HandleInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	// This creates a new pointer to the Handler object
	return new(Handler), nil
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return 
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		/*
			First argument is the http status code, 
			whereas the second argument is the body of the request
		*/
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

