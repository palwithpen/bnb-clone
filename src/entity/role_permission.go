package entity

// RolePermission represents the many-to-many relationship between roles and permissions.
type RolePermission struct {
	RoleID       int `pg:"role_id,pk,notnull"`       // Foreign key to roles
	PermissionID int `pg:"permission_id,pk,notnull"` // Foreign key to permissions

	// Foreign key relationships
	Role       *Role       `pg:"rel:has-one,fk:role_id"`
	Permission *Permission `pg:"rel:has-one,fk:permission_id"`
}
