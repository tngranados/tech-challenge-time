package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tngranados/tech-challenge-time/models"
)

// TestSessionStore checks the methods in the SessionStore
func TestSessionStore(t *testing.T) {
	db, cleanup := NewTestDB(t)
	defer cleanup()

	sessionStore, err := NewSessionStore(db)
	require.NoError(t, err)

	testSession := &models.Session{
		ID:         1,
		Name:       "Test",
		StartedAt:  time.Now().Add(-30 * time.Minute).UTC(),
		FinishedAt: time.Now().UTC(),
	}
	testSession2 := &models.Session{
		ID:        2,
		Name:      "Test2",
		StartedAt: time.Now().Add(-1 * time.Hour).UTC(),
	}

	// Insert a couple of sessions in a transaction.
	tx, err := db.Begin()
	require.NoError(t, err)

	err = sessionStore.Insert(tx, testSession)
	require.NoError(t, err)
	err = sessionStore.Insert(tx, testSession2)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)

	// Get a session
	recoveredSession, err := sessionStore.Get(db, testSession.ID)
	require.NoError(t, err)
	require.Equal(t, testSession, recoveredSession)

	// Get all sessions
	sessionsList := []*models.Session{testSession, testSession2}

	recoveredList, err := sessionStore.GetAll(db)
	require.NoError(t, err)
	require.Equal(t, sessionsList, recoveredList)

	// Get all finished
	finishedList := []*models.Session{testSession}

	recoveredFinishedList, err := sessionStore.GetAllFinished(db)
	require.NoError(t, err)
	require.Equal(t, finishedList, recoveredFinishedList)

	// Get all unfinished sessions
	unfinishedList := []*models.Session{testSession2}

	recoveredUnfinishedList, err := sessionStore.GetAllUnfinished(db)
	require.NoError(t, err)
	require.Equal(t, unfinishedList, recoveredUnfinishedList)

	// Update session
	testSession2.FinishedAt = time.Now().UTC()

	err = sessionStore.Update(db, testSession2)
	require.NoError(t, err)

	updatedSession, err := sessionStore.Get(db, testSession2.ID)
	require.NoError(t, err)
	require.Equal(t, testSession2, updatedSession)

	// Delete session
	err = sessionStore.Delete(db, testSession.ID)
	require.NoError(t, err)

	deletedItem, err := sessionStore.Get(db, testSession.ID)
	require.NoError(t, err)
	require.Nil(t, deletedItem)
}
