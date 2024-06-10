package httpv1

import "github.com/gin-gonic/gin"

func (h *Handler) initUniversityRoutes(api *gin.RouterGroup) {
	university := api.Group("/university", h.setDomainFromRequest)
	{
		university.GET("/verify/:hash")
		university.GET("/header-image", h.getHeaderImage)
		university.GET("/logo-image", h.getLogoImage)
		university.GET("/css", h.getCSS)
		university.GET("/css/colors", h.getColors)

		authSettings := university.Group("/settings", h.setUserFromRequest)
		{
			authSettings.GET("/info", h.getUniversity)
			authSettings.GET("/settings")
			authSettings.PATCH("/css", h.patchCSS)
		}

		historyGroup := university.Group("/history")
		{
			historyGroup.GET("", h.getUniversityHistory)

			authHistoryGroup := historyGroup.Group("/", h.setUserFromRequest)
			{
				authHistoryGroup.POST("", h.postUniversityHistory)
				authHistoryGroup.PATCH("")
			}
		}

		newsGroup := university.Group("/news")
		{
			newsGroup.GET("", h.getAllNews)
			newsGroup.GET("/header-image/:name", h.getImage)
			newsGroup.GET("/:id")

			newsGroup.POST("", h.postNews)
			newsGroup.POST("/header-image/:name", h.setImage)

			newsGroup.GET("/:id/report")
			newsGroup.POST("/:id/report")

			authNewsGroup := newsGroup.Group("/", h.setUserFromRequest)
			{
				authNewsGroup.GET("/new")
				authNewsGroup.POST("/new")

				authNewsGroup.GET("/:id/edit")
				authNewsGroup.PUT("/:id/edit")
				authNewsGroup.DELETE("/:id/edit")
			}
		}

		documentsGroup := university.Group("/docs")
		{
			documentsGroup.GET("", h.getAllUniversityDocs)
			documentsGroup.GET("/:doc_id", h.getDocs)

			documentsGroup.POST("", h.postDocs)
			documentsGroup.POST("/:doc_id", h.setDocs)
			documentsGroup.GET("/bachelors", h.getAllBachelors)
			documentsGroup.POST("/bachelors", h.postStudyPlanDocs)
			documentsGroup.POST("/bachelors/:doc_id", h.setDocs)
			documentsGroup.GET("/mags", h.getAllMags)
			documentsGroup.POST("/mags", h.postStudyPlanDocs)
			documentsGroup.POST("/mags/:doc_id", h.setDocs)
			documentsGroup.GET("/enrollee", h.getAllEnrollsDocs)
			documentsGroup.POST("/enrollee", h.postStudyPlanDocs)
			documentsGroup.POST("/enrollee/:doc_id", h.setDocs)
			//authDocumentsGroup := documentsGroup.Group("/")
			//{
			//	//authDocumentsGroup.GET("/new")
			//	//authDocumentsGroup.POST("/new")
			//	//
			//	//authDocumentsGroup.GET("/:id/edit")
			//	//authDocumentsGroup.PUT("/:id/edit")
			//	//authDocumentsGroup.DELETE("/:id/edit")
			//}
		}

		programsGroup := university.Group("/programs")
		{
			programsGroup.GET("")

			programsGroup.GET("/:id")

			authProgramsGroup := programsGroup.Group("/")
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

			authAppNewsGroup := applicantsGroup.Group("/")
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

				authAppDocs := appDocs.Group("/")
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

				authAppTables := appTables.Group("/")
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
