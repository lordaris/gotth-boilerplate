package app

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Server struct {
		Port string
	}
}

// Application holds dependencies and configuration
type Application struct {
	Config Config
	DB     *sqlx.DB
}

// NewApplication creates a new Application instance
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

// ConnectToDatabase establishes connection to PostgreSQL
func (app *Application) ConnectToDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		app.Config.DB.Host,
		app.Config.DB.User,
		app.Config.DB.Password,
		app.Config.DB.Name,
		app.Config.DB.Port,
	)
	log.Printf("Connecting to database: %s@%s:%s/%s",
		app.Config.DB.User,
		app.Config.DB.Host,
		app.Config.DB.Port,
		app.Config.DB.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return err
	}

	app.DB = db
	log.Println("Successfully connected to database")
	return nil
}
