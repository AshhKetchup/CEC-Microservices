package handler

import (
	pb "auth/proto/gen"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"auth/internal/helper"

	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	db *sql.DB
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (s *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	var userID, hashedPassword, userRole string
	err := s.db.QueryRowContext(ctx,
		"SELECT id, password, role FROM users WHERE email = $1",
		req.Email,
	).Scan(&userID, &hashedPassword, &userRole)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		fmt.Printf("DB Error: %v\n", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := helper.GenerateJWT(userID, userRole)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token.go")
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	var exists bool
	err := s.db.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		req.Email,
	).Scan(&exists)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check user existence")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	var userID string
	err = s.db.QueryRowContext(ctx,
		"INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		req.Email,
		string(hashedPassword),
		req.Role.String(),
		time.Now(),
	).Scan(&userID)

	if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &pb.RegisterResponse{Id: userID}, nil
}

func (s *AuthHandler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	isValid, userId, userRole, err := helper.ValidateJWT(req.Token)
	fmt.Printf("UserID: %s | Role: %s | Error: %v\n", userId, userRole, err)
	if err != nil {
		return &pb.ValidateTokenResponse{
			Valid: false,
			Base: &pb.BaseResponse{
				Code:    int32(codes.Internal),
				Message: "Token validation failed",
			},
		}, nil
	}

	return &pb.ValidateTokenResponse{
		Valid:  isValid,
		UserId: userId,
		Role:   userRole,
		Base: &pb.BaseResponse{
			Code:    int32(codes.OK),
			Message: "Token validated successfully",
		},
	}, nil
}
