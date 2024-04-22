package httpv1

import "github.com/gin-gonic/gin"

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-in")
		admins.GET("/sign-in")
		admins.POST("/auth/refresh")
		admins.GET("/verify/:hash")
	}

	authenticated := admins.Group("/", h.userIdentity)
	{
		sitesGroup := authenticated.Group("/sites")
		{
			sitesGroup.GET("")

			sitesGroup.GET("/new")
			sitesGroup.POST("/new")

			sitesGroup.GET("/:id")
			sitesGroup.PATCH("/:id")
			sitesGroup.DELETE("/:id")
		}

		// Не в приоритете
		requestsGroup := authenticated.Group("/requests")
		{
			requestsGroup.GET("")

			requestsGroup.GET("/new")
			requestsGroup.POST("/new")

			requestsGroup.GET("/:id")
			requestsGroup.PATCH("/:id")
			requestsGroup.DELETE("/:id")
		}

		employeesGroup := authenticated.Group("/employees")
		{
			employeesGroup.GET("")

			employeesGroup.GET("/new")
			employeesGroup.POST("/new")

			employeesGroup.GET("/:id")
			employeesGroup.PATCH("/:id")
			employeesGroup.DELETE("/:id")
		}

		// Не в приоритете
		notificationsGroup := authenticated.Group("/notifications")
		{
			notificationsGroup.GET("")

			notificationsGroup.GET("/new")
			notificationsGroup.POST("/new")

			notificationsGroup.GET("/:id")
		}
	}
}
