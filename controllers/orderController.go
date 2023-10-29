package controllers

import (
	"midterm/models"
	"midterm/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context, db *gorm.DB) {
	var request struct {
		OrderedAt    string        `json:"orderedAt" binding:"required"`
		CustomerName string        `json:"customerName" binding:"required"`
		Items        []models.Item `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		OrderedAt:    time.Now(),
		CustomerName: request.CustomerName,
		Items:        request.Items,
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": r})
		}
	}()

	if err := services.CreateOrder(db, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
	}

	for i := range request.Items {
		request.Items[i].OrderID = order.ID
		if err := services.CreateItem(db, &request.Items[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func GetOrdersWithItems(c *gin.Context, db *gorm.DB) {
	orders, err := services.GetOrderstWithItems(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrderByIdWithItems(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := services.GetOrderByIdWithItems(db, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context, db *gorm.DB) {
	var request struct {
		CustomerName string        `json:"customerName"`
		Items        []models.Item `json:"items"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	existingOrder, err := services.GetOrderByIdWithItems(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	existingOrder.CustomerName = request.CustomerName

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": r})
		}
	}()

	if err := services.UpdateOrder(db, uint(id), existingOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
	}

	for i := range existingOrder.Items {
		err := services.DeleteItemById(db, uint(existingOrder.Items[i].ID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			panic(err)
		}
	}

	for i := range request.Items {
		request.Items[i].OrderID = uint(id)
		if err := services.CreateItem(db, &request.Items[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := services.GetOrderByIdWithItems(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": r})
		}
	}()

	for i := range order.Items {
		err := services.DeleteItemById(db, uint(order.Items[i].ID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			panic(err)
		}
	}

	err = services.DeleteOrderById(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		panic(err)
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
