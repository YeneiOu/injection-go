package main

import (
	"fmt"
	"injection-go/controllers"
	"injection-go/repositories"
	"injection-go/usecase"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection configuration
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"104.196.19.55", // host
		5432,            // port
		"postgres",      // user
		"p@ssw0rd",      // password
		"test_db",       // dbname
	)

	// Connect to database
	db, err := sqlx.Connect("postgres", dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Initialize repository
	memberRepo := repositories.NewMembershipRepository(db)

	_ = memberRepo
	// fmt.Printf("Database connected successfully, repository initialized: %v\n", memberRepo)

	// Initialize mock repository
	mockRepo := repositories.NewMembershipRepositoryMock()

	// Initialize usecase
	memberService := usecase.NewMembershipService(mockRepo)

	// Initialize controllers
	controller := controllers.NewMembershipHandler(memberService)

	// Initialize router
	router := fiber.New()

	// Routes Membership
	memberRoutes := router.Group("/members")
	memberRoutes.Get("/", controller.GetAll)
	memberRoutes.Get("/:id", controller.GetById)

	router.Listen(":8080")
}
