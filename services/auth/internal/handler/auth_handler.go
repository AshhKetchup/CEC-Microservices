package handler

import (
	pb "cec/services/auth/proto/gen"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// JWT configuration (will be loaded from env)
var (
	jwtSecretKey  string
	jwtExpiration time.Duration
	jwtIssuer     string
)

// Custom claims structure
type customClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load("auth.env"); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	// Load JWT configuration
	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		panic("JWT_SECRET_KEY is required in .env")
	}

	expHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		panic("Invalid JWT_EXPIRATION_HOURS in .env")
	}
	jwtExpiration = time.Duration(expHours) * time.Hour

	jwtIssuer = os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		jwtIssuer = "cec-auth-service" // default
	}
}

type AuthHandler struct {
	db *sql.DB
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (s *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Validate input
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// Get user from database
	var userID string
	var hashedPassword string
	err := s.db.QueryRowContext(ctx,
		"SELECT id, password FROM users WHERE email = $1",
		req.Email,
	).Scan(&userID, &hashedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// Generate JWT token (using your actual JWT implementation)
	token, err := GenerateJWT(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Validate input
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// Check if user already exists
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

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	// Create user in database
	var userID string
	err = s.db.QueryRowContext(ctx,
		"INSERT INTO users (email, password, created_at) VALUES ($1, $2, $3) RETURNING id",
		req.Email,
		string(hashedPassword),
		time.Now(),
	).Scan(&userID)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &pb.RegisterResponse{
		Id: userID,
	}, nil
}

// Helper functions
func GenerateJWT(userID string) (string, error) {
	// Implement your actual JWT generation logic here
	return "generated.jwt.token", nil
}

// // Add these database initialization steps in your application setup
// func CreateTables(db *sql.DB) error {
// 	_, err := db.Exec(`
// 		CREATE TABLE IF NOT EXISTS users (
// 			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
// 			email TEXT UNIQUE NOT NULL,
// 			password TEXT NOT NULL,
// 			created_at TIMESTAMP NOT NULL
// 		);
// 	`)
// 	return err
// }

func ValidateJWT(tokenString string) (bool, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return false, "", err
	}

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return true, claims.UserID, nil
	}

	return false, "", errors.New("invalid token")
}

func (s *AuthHandler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Your token validation logic here
	isValid, userId, err := ValidateJWT(req.Token)
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
		Base: &pb.BaseResponse{
			Code:    int32(codes.OK),
			Message: "Token validated successfully",
		},
	}, nil
}
