package server

import (
	"context"

	auth "github.com/backend-ids/authentication"
	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Db         *gorm.DB
	jwtManager *auth.JWTManager
}

func NewAuthServer(jwtManager *auth.JWTManager, db *gorm.DB) pb.AuthServiceServer {
	return &AuthServer{jwtManager: jwtManager, Db: db}
}

// user login return JWT token
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// log.Println("Server : Logging in the User")
	user := models.User{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	var count int64
	s.Db.Model(&models.User{}).Where("name=? and password=?", user.Name, user.Password).Find(&user).Count(&count)
	if count != 0 {
		token, err := s.jwtManager.Generate(&user)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate access token")
		}
		res := &pb.LoginResponse{AccessToken: token}
		return res, nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "wrong credentials")
}
