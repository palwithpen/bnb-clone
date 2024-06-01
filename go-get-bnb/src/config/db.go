package config

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"palwithpen.github.com/airbnb/src/entity"
)

var DB *pg.DB

func Init() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "172.19.169.166:31431",
		Database: "bnb",
	}

	DB = pg.Connect(opts)

	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connection established")
	if err := createSchema(); err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}
	log.Println("Created Tables")
}

func createSchema() error {
	models := []interface{}{
		(*entity.User)(nil),
		(*entity.Role)(nil),
		(*entity.Permission)(nil),
		(*entity.Role)(nil),
		(*entity.RolePermission)(nil),
		(*entity.UserSession)(nil),

		// Add other models here
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	log.Println("Database schema created")
	return nil
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
