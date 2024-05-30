package entity

import "time"

/*
type User struct {
	ID       string     `pg:"id,pk,type:varchar(100)"` // 'pk' stands for primary key
	Name     string     `pg:"name ,type:varchar(100)"`
	Email    string     `pg:"email,type:varchar(100)"`
	Password string     `pg:"password,type:varchar(100)"`
	Profiles []*Profile `pg:"rel:has-many"` // One-to-many relationship
}

type Profile struct {
	ID     string `pg:"id,pk,type:varchar(100)"`           // 'pk' stands for primary key
	UserID string `pg:"user_id,notnull,on_delete:CASCADE"` // Foreign key field with constraint
	User   *User  `pg:"rel:has-one"`                       // Reference to the User model
	Bio    string `pg:"bio,type:varchar(100)"`
}
*/

type User struct {
	UserID       string    `pg:"user_id,pk,type:varchar(100)"`              // Primary key
	Username     string    `pg:"username,unique,notnull,type:varchar(100)"` // Unique username
	Email        string    `pg:"email,unique,notnull,type:varchar(100)"`    // Unique email
	PasswordHash string    `pg:"password_hash,notnull,type:varchar(100)"`   // Password hash
	CreatedAt    time.Time `pg:"created_at,default:now()"`                  // Creation timestamp
	UpdatedAt    time.Time `pg:"updated_at,default:now()"`                  // Update timestamp
}
