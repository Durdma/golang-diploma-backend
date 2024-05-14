package httpv1

import "github.com/gin-gonic/gin"

func (h *Handler) initUniversityRoutes(api *gin.RouterGroup) {
	university := api.Group("/", h.setUniversityFromRequest)
	{
		university.GET("/sign-in")
		university.POST("/sign-in")

		university.POST("/auth/refresh")

		university.GET("/verify/:hash")

		authSettings := university.Group("/", h.userIdentity)
		{
			authSettings.GET("/settings")
			authSettings.PUT("/settings")
		}

		newsGroup := university.Group("/news")
		{
			newsGroup.GET("")

			newsGroup.GET("/:id")

			newsGroup.GET("/:id/report")
			newsGroup.POST("/:id/report")

			authNewsGroup := newsGroup.Group("/", h.userIdentity)
			{
				authNewsGroup.GET("/new")
				authNewsGroup.POST("/new")

				authNewsGroup.GET("/:id/edit")
				authNewsGroup.PUT("/:id/edit")
				authNewsGroup.DELETE("/:id/edit")
			}
		}

		structureGroup := university.Group("/structure")
		{
			structureGroup.GET("")

			structureGroup.GET("/:id")

			structureGroup.GET("/:id/:depart_id")

			authStructureGroup := structureGroup.Group("/", h.userIdentity)
			{
				authStructureGroup.GET("/new")
				authStructureGroup.POST("/new")

				authStructureGroup.GET("/:id/edit")
				authStructureGroup.PUT("/:id/edit")
				authStructureGroup.DELETE("/:id/edit")

				authStructureGroup.GET("/:id/new")
				authStructureGroup.POST("/:id/new")

				authStructureGroup.GET("/:id/:depart_id/edit")
				authStructureGroup.PUT("/:id/:depart_id/edit")
				authStructureGroup.DELETE("/:id/:depart_id/edit")
			}
		}

		documentsGroup := university.Group("/documents")
		{
			documentsGroup.GET("")

			documentsGroup.GET("/:id")

			documentsGroup.GET("/:id/url")

			authDocumentsGroup := documentsGroup.Group("/", h.userIdentity)
			{
				authDocumentsGroup.GET("/new")
				authDocumentsGroup.POST("/new")

				authDocumentsGroup.GET("/:id/edit")
				authDocumentsGroup.PUT("/:id/edit")
				authDocumentsGroup.DELETE("/:id/edit")
			}
		}

		programsGroup := university.Group("/programs")
		{
			programsGroup.GET("")

			programsGroup.GET("/:id")

			authProgramsGroup := programsGroup.Group("/", h.userIdentity)
			{
				authProgramsGroup.GET("/new")
				authProgramsGroup.POST("/new")

				authProgramsGroup.GET("/:id/edit")
				authProgramsGroup.PUT("/:id/edit")
				authProgramsGroup.DELETE("/:id/edit")
			}
		}

		applicantsGroup := university.Group("/applicants")
		{
			applicantsGroup.GET("")

			applicantsGroup.GET("/:id")

			applicantsGroup.GET("/:id/report")
			applicantsGroup.POST("/:id/report")

			authAppNewsGroup := applicantsGroup.Group("/", h.userIdentity)
			{
				authAppNewsGroup.GET("/new")
				authAppNewsGroup.POST("/new")

				authAppNewsGroup.GET("/:id/edit")
				authAppNewsGroup.PUT("/:id/edit")
				authAppNewsGroup.DELETE("/:id/edit")
			}

			appDocs := applicantsGroup.Group("/documents")
			{
				appDocs.GET("")

				appDocs.GET("/:id")

				appDocs.GET("/:id/:url")

				authAppDocs := appDocs.Group("/", h.userIdentity)
				{
					authAppDocs.GET("/new")
					authAppDocs.POST("/new")

					authAppDocs.GET("/:id/edit")
					authAppDocs.PUT("/:id/edit")
					authAppDocs.DELETE("/:id/edit")
				}
			}

			appTables := applicantsGroup.Group("/table")
			{
				appTables.GET("")

				appTables.GET("/:id")

				authAppTables := appTables.Group("/", h.userIdentity)
				{
					authAppTables.GET("/new")
					authAppTables.POST("/new")

					authAppTables.GET("/:id/persons")
					authAppTables.POST("/:id/persons")

					authAppTables.GET("/:id/persons/:person_id")
					authAppTables.PUT("/:id/persons/:person_id")
					authAppTables.DELETE("/:id/persons/:person_id")
				}
			}
		}
	}
}
