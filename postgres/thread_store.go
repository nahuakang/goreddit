package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nahuakang/goreddit"
)

// ThreadStore inherits methods from sqlx.DB
type ThreadStore struct {
	*sqlx.DB
}

// Thread method gets a thread from the database based on id input
func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("Error getting thread: %w", err)
	}
	return t, nil
}

// Threads method gets all the threads in the database
func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var tt []goreddit.Thread
	if err := s.Select(&tt, `SELCT * FROM threads`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("Error getting threads: %w", err)
	}
	return tt, nil
}

// CreateThread creates a thread in the database
func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("Error creating thread: %w", err)
	}
	return nil
}

// UpdateThread updates a thread in the database
func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3) RETURNING *`,
		t.Title,
		t.Description,
		t.ID); err != nil {
		return fmt.Errorf("Error updating thread: %w", err)
	}
	return nil
}

// DeleteThread deletes a thread in the database
func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("Error deleting thread: %w", err)
	}
	return nil
}
