package entity

type Permission struct {
	PermissionID   int    `pg:"permission_id,pk"`                                 // Primary key
	PermissionName string `pg:"permission_name,unique,notnull,type:varchar(100)"` // Unique permission name
	Description    string `pg:"description"`                                      // Description of the permission
}

type Role struct {
	RoleID      int    `pg:"role_id,pk"`                                 // Primary key
	RoleName    string `pg:"role_name,unique,notnull,type:varchar(100)"` // Unique role name
	Description string `pg:"description"`                                // Description of the role
}
