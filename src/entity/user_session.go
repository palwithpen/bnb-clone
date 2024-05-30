package entity

import (
	"time"
)

// UserSession represents a user session in the database.
type UserSession struct {
	SessionID    int       `pg:"session_id,pk"`                     // Primary key
	UserID       string    `pg:"user_id,notnull,type:varchar(100)"` // Foreign key to users
	SessionToken string    `pg:"session_token,unique,notnull"`      // Unique session token
	CreatedAt    time.Time `pg:"created_at,default:now()"`          // Creation timestamp
	ExpiresAt    time.Time `pg:"expires_at,notnull"`                // Expiration timestamp

	// Foreign key relationship
	User *User `pg:"rel:has-one,fk:user_id"`
}
