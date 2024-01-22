package app

import (
	"errors"
	"fmt"
	"net/http"
	"sas/internal/controller/httpv1"
	"sas/internal/repository"
	"sas/internal/server"
	"sas/internal/service"
)

func Run() {
	repo := repository.NewTmpRepo()

	adminService := service.NewAdminService(repo)

	handlers := httpv1.NewHandler(adminService)

	srv := server.NewServer(handlers.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("%v\n", err.Error())
		}
	}()

	fmt.Println("Server started!")

}
