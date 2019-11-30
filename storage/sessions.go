package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/tngranados/tech-challenge-time/models"
)

// SessionStore is the entity that manages Sessions storage in the database.
type SessionStore struct {
}

// NewSessionStore returns a new SessionStore entity.
func NewSessionStore(db *sql.DB) (*SessionStore, error) {
	// Create table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			started_at DATETIME,
			finished_at DATETIME
		)
	`)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sessions table in the database")
	}

	return &SessionStore{}, nil
}

func (s *SessionStore) selectSessions(dbSession DBSession, condition string, args ...interface{}) ([]*models.Session, error) {
	if condition != "" {
		condition = " WHERE " + condition
	}
	query := "SELECT id, name, started_at, finished_at FROM sessions" + condition + ";"

	rows, err := dbSession.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*models.Session
	for rows.Next() {
		session := &models.Session{}
		err = rows.Scan(&session.ID, &session.Name, &session.StartedAt, &session.FinishedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, session)
	}

	return result, nil
}

// Get returns the session with the provided id.
func (s *SessionStore) Get(dbSession DBSession, id int) (*models.Session, error) {
	list, err := s.selectSessions(dbSession, "id = $1", id)
	if err != nil {
		return nil, err
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return list[0], nil
}

// GetAll returns all the sessions in the database.
func (s *SessionStore) GetAll(dbSession DBSession) ([]*models.Session, error) {
	return s.selectSessions(dbSession, "")
}

// GetAllFinished returns all the finished sessions in the database.
func (s *SessionStore) GetAllFinished(dbSession DBSession) ([]*models.Session, error) {
	zeroTime := time.Time{}
	return s.selectSessions(dbSession, "finished_at > $1", zeroTime)
}

// GetAllUnfinished returns all the unfinished sessions in the database.
func (s *SessionStore) GetAllUnfinished(dbSession DBSession) ([]*models.Session, error) {
	zeroTime := time.Time{}
	return s.selectSessions(dbSession, "finished_at = $1", zeroTime)
}

// Insert adds a new session to the database.
func (s *SessionStore) Insert(dbSession DBSession, session *models.Session) error {
	query := "INSERT INTO sessions (name, started_at, finished_at) VALUES ($1, $2, $3);"

	_, err := dbSession.Exec(query, session.Name, session.StartedAt, session.FinishedAt)
	return err
}

// Update updates an session in the database.
func (s *SessionStore) Update(dbSession DBSession, session *models.Session) error {
	query := "UPDATE sessions SET name = $1, started_at = $2, finished_at = $3 WHERE id = $4;"

	res, err := dbSession.Exec(query, session.Name, session.StartedAt, session.FinishedAt, session.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("sql update should have updated 1 session, but updated %v instead", affected)
	}

	return err
}

// Delete removes an session from the database.
func (s *SessionStore) Delete(dbSession DBSession, id int) error {
	query := "DELETE FROM sessions WHERE id = $1;"

	_, err := dbSession.Exec(query, id)
	return err
}
