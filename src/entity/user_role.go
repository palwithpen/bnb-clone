package entity

type UserRole struct {
	UserID int `pg:"user_id,pk,notnull"` // Foreign key to users
	RoleID int `pg:"role_id,pk,notnull"` // Foreign key to roles

	User *User `pg:"rel:has-one,fk:user_id"`
	Role *Role `pg:"rel:has-one,fk:role_id"`
}
