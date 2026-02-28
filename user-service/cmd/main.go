package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/middleware"
	"user-service/internal/service"
	"user-service/internal/storage"

	grpcserver "user-service/internal/grpc"
	pb "user-service/proto/user"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {

	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	userStorage := storage.NewUserPostgressStorage(database)
	userService := service.NewUserService(userStorage)
	authService := service.NewAuthService(userStorage, "secret_key")
	userHandler := handler.NewUserHandler(userService, authService)

	grpcServer := grpc.NewServer()
	userGRPCServer := grpcserver.NewUserGRPCServer(userService)
	pb.RegisterUserServiceServer(grpcServer, userGRPCServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Не удалось запустить gRPC:", err)
	}

	go func() {
		grpcServer.Serve(lis)
	}()

	r := chi.NewRouter()
	r.Post("/users", userHandler.HandleRegister)
	r.Post("/login", userHandler.HandleLogin)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/users/{id}", userHandler.HandleGetUsers)
		r.Get("/users", userHandler.HandleGetAllUsers)
		r.Delete("/users/{id}", userHandler.HandleDeleteUser)
	})

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", r)
}
