package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/tngranados/tech-challenge-time/models"
)

// TestSessions tests the methods in the SessionCtrl.
func TestSessions(t *testing.T) {
	router, cleanup := createTestServer(t)
	defer cleanup()

	var sessions []*models.Session
	now := time.Now().UTC()
	// Add a few sessions
	for i := 1; i <= 5; i++ {
		session := &models.Session{
			ID:         i,
			Name:       "session" + strconv.Itoa(i),
			StartedAt:  now.Add(time.Duration(-i) * time.Hour),
			FinishedAt: now.Add(time.Duration(-i+1) * time.Hour),
		}

		addSession(t, router, session)
		sessions = append(sessions, session)
	}

	// GetAll should return the just created sessions
	getSessionsAndCompare(t, router, "sessions", sessions)

	// Update session, then Get it and compare
	sessions[0].Name = "new name"
	body, err := json.Marshal(sessions[0])
	require.NoError(t, err)

	w := performRequest(t, router, "PUT", "/api/sessions", body)
	err = checkCode(http.StatusOK, w.Code)
	require.NoError(t, err)

	w = performRequest(t, router, "GET", fmt.Sprintf("/api/sessions/%v", sessions[0].ID), nil)
	err = checkCode(http.StatusOK, w.Code)
	require.NoError(t, err)

	recoveredSession := &models.Session{}
	err = json.Unmarshal(w.Body.Bytes(), recoveredSession)
	require.NoError(t, err)
	require.Equal(t, sessions[0], recoveredSession)

	// Add a couple of unfinished sessions
	var unfinishedSessions []*models.Session
	for i := 1; i <= 2; i++ {
		session := &models.Session{
			ID:        len(sessions) + i,
			Name:      "unfinished session" + strconv.Itoa(i),
			StartedAt: now.Add(time.Duration(-i) * time.Hour),
		}

		addSession(t, router, session)
		unfinishedSessions = append(unfinishedSessions, session)
	}

	// GetFinished should return the previously created finished sessions
	getSessionsAndCompare(t, router, "finished-sessions", sessions)

	// GetUnfinished should return the just created unfinished sessions
	getSessionsAndCompare(t, router, "unfinished-sessions", unfinishedSessions)

	// GetAll should return all of the sessions
	getSessionsAndCompare(t, router, "sessions", append(sessions, unfinishedSessions...))

	// Delete session should return ok
	w = performRequest(t, router, "DELETE", fmt.Sprintf("/api/sessions/%v", sessions[0].ID), nil)

	err = checkCode(http.StatusOK, w.Code)
	require.NoError(t, err)

	// Remove the same session from our expected list
	sessions = sessions[1:]

	// Check if lists are equal
	getSessionsAndCompare(t, router, "sessions", append(sessions, unfinishedSessions...))
}

func addSession(tb testing.TB, router *gin.Engine, session *models.Session) {
	body, err := json.Marshal(session)
	require.NoError(tb, err)
	w := performRequest(tb, router, "POST", "/api/sessions", body)

	err = checkCode(http.StatusOK, w.Code)
	require.NoError(tb, err)
}

func getSessionsAndCompare(tb testing.TB, router *gin.Engine, path string, expectedList []*models.Session) {
	w := performRequest(tb, router, "GET", "/api/"+path, nil)

	err := checkCode(http.StatusOK, w.Code)
	require.NoError(tb, err)

	defer w.Result().Body.Close()
	receivedList := []*models.Session{}

	err = json.NewDecoder(w.Result().Body).Decode(&receivedList)
	require.NoError(tb, err)
	require.Equal(tb, expectedList, receivedList)
}
