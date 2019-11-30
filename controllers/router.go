package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tngranados/tech-challenge-time/storage"
)

// SetupRouter creates the router of the web service.
func SetupRouter(debug bool, db *sql.DB) (*gin.Engine, error) {
	// Setup stores
	sessionStore, err := storage.NewSessionStore(db)
	if err != nil {
		return nil, fmt.Errorf("Error setting up the user store: %s", err.Error())
	}

	// Setup controllers
	sessionCtrl := NewSessionsCtrl(db, sessionStore)

	// Setup router
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.Use(ErrorMiddleware())

	api := router.Group("api")

	// Sessions
	api.GET("sessions", sessionCtrl.GetAll)
	api.GET("sessions/:id", sessionCtrl.Get)
	api.GET("finished-sessions", sessionCtrl.GetFinished)
	api.GET("unfinished-sessions", sessionCtrl.GetUnfinished)
	api.POST("sessions", sessionCtrl.Add)
	api.PUT("sessions", sessionCtrl.Update)
	api.DELETE("sessions/:id", sessionCtrl.Delete)

	return router, nil
}
