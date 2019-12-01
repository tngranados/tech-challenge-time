package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tngranados/tech-challenge-time/models"
	"github.com/tngranados/tech-challenge-time/storage"
)

// SessionCtrl handles session related requests.
type SessionCtrl struct {
	db           *sql.DB
	sessionStore *storage.SessionStore
}

// NewSessionsCtrl creates a new instance of the sessions controller.
func NewSessionsCtrl(db *sql.DB, sessionStore *storage.SessionStore) *SessionCtrl {
	return &SessionCtrl{
		db:           db,
		sessionStore: sessionStore,
	}
}

// Get returns the requested session.
func (ctrl *SessionCtrl) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	session, err := ctrl.sessionStore.Get(ctrl.db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if session == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, session)
}

// GetAll returns the list of every session.
func (ctrl *SessionCtrl) GetAll(c *gin.Context) {
	sessions, err := ctrl.sessionStore.GetAll(ctrl.db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, sessions)
}

// GetFinished returns the list of every finished session.
func (ctrl *SessionCtrl) GetFinished(c *gin.Context) {
	sessions, err := ctrl.sessionStore.GetAllFinished(ctrl.db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, sessions)
}

// GetUnfinished returns the list of every unfinished session.
func (ctrl *SessionCtrl) GetUnfinished(c *gin.Context) {
	sessions, err := ctrl.sessionStore.GetAllUnfinished(ctrl.db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, sessions)
}

// Add inserts a session in the database.
func (ctrl *SessionCtrl) Add(c *gin.Context) {
	session := &models.Session{}
	err := c.BindJSON(session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if session.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "A session must have a name")
		return
	}

	err = ctrl.sessionStore.Insert(ctrl.db, session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// Update updates a session in the database.
func (ctrl *SessionCtrl) Update(c *gin.Context) {
	session := &models.Session{}
	err := c.BindJSON(session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if session.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "You can't remove the name of a session")
	}

	err = ctrl.sessionStore.Update(ctrl.db, session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// Delete removes a session from the database.
func (ctrl *SessionCtrl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = ctrl.sessionStore.Delete(ctrl.db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
