package entity

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
