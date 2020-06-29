package web

import (
	"context"
	"database/sql"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

// SessionData contains data for flash messages
type SessionData struct {
	FlashMessage string
	// UserID uuid.UUID
}

// NewSessionManager manages sessions for Goreddit
func NewSessionManager(dataSourceName string) (*scs.SessionManager, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	sessions := scs.New()
	sessions.Store = postgresstore.New(db)

	return sessions, nil
}

// GetSessionData grabs data from the SessionManager
func GetSessionData(ctx context.Context, session *scs.SessionManager) SessionData {
	var data SessionData

	data.FlashMessage = session.PopString(ctx, "flash")
	// data.UserID, _ = session.Get(ctx, "user_id").(uuid.UUID)

	return data
}
