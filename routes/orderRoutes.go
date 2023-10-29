package routes

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(route *gin.Engine, db *gorm.DB) {
	route.POST("/api/orders", func(c *gin.Context) {
		controllers.CreateOrder(c, db)
	})

	route.GET("/api/orders", func(c *gin.Context) {
		controllers.GetOrdersWithItems(c, db)
	})

	route.GET("/api/orders/:id", func(c *gin.Context) {
		controllers.GetOrderByIdWithItems(c, db)
	})

	route.PUT("/api/orders/:id", func(c *gin.Context) {
		controllers.UpdateOrder(c, db)
	})

	route.DELETE("/api/orders/:id", func(c *gin.Context) {
		controllers.DeleteOrder(c, db)
	})
}
