package main

import (
	"github.com/RicardoEmm/taskforge/internal/config"
	"github.com/RicardoEmm/taskforge/internal/database"
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/RicardoEmm/taskforge/internal/repository"
	"github.com/RicardoEmm/taskforge/internal/router"
	"github.com/RicardoEmm/taskforge/internal/service"
)

func main() {
	cfg := config.Load()
	db := database.Connet(cfg)

	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	projectRepo := repository.NewGormProjectRepository(db)
	projectService := service.NewProjectService(projectRepo, userRepo)
	projectHandler := handler.NewProjectHandler(projectService)

	taskRepo := repository.NewGormTaskRepository(db)
	taskService := service.NewTaskService(projectRepo, taskRepo, userRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	tokenService := service.NewTokenService(cfg.JWTSecret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)

	authHandler := handler.NewAuthHandler(tokenService, userService)

	setup := router.Handlers{
		AuthHandler:    authHandler,
		ProjectHandler: projectHandler,
		TaskHandler:    taskHandler,
		UserHandler:    userHandler,
	}

	router := router.Setup(setup, tokenService)

	router.Run(":" + cfg.AppPort)
}
