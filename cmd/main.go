package main

import (
	goclean "goclean/internal/delivery/grpc"
	"goclean/internal/domain"
	"goclean/internal/repository/postgres"
	"goclean/internal/usecase"
	"goclean/pkg/database"
	"log"
	"net"
	"os"

	pb "goclean/proto"

	// gc_grpc "goclean/internal/delivery/goclean"

	"google.golang.org/grpc"
)

func main() {
	//set env for db
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "go_clean")
	os.Setenv("DB_PORT", "5432")

	// Initialize database
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&domain.User{})

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, goclean.NewUserHandler(userUseCase))

	log.Printf("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
