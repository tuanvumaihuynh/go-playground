package http

import (
	"context"
	"time"

	"github.com/tuanvumaihuynh/go-playground/oapi-codegen/internal/http/gen"
)

type userHandler struct {
}

func newUserHandler() *userHandler {
	return &userHandler{}
}

func (h userHandler) GetUserById(ctx context.Context, request gen.GetUserByIdRequestObject) (gen.GetUserByIdResponseObject, error) {
	return gen.GetUserById200JSONResponse(gen.UserResponse{
		Id:        request.UserId,
		Username:  "username",
		Email:     "email",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}), nil
}

func (h userHandler) ListUsers(ctx context.Context, request gen.ListUsersRequestObject) (gen.ListUsersResponseObject, error) {
	return gen.ListUsers200JSONResponse(gen.UserListResponse{
		Items: []gen.UserResponse{
			{
				Id:        1,
				Username:  "username",
				Email:     "email",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        2,
				Username:  "username2",
				Email:     "email2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		TotalItems: 2,
	}), nil
}

func (h userHandler) CreateUser(ctx context.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {
	return gen.CreateUser201JSONResponse(gen.UserResponse{
		Id:        1,
		Username:  request.Body.Username,
		Email:     string(request.Body.Email),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}), nil
}
