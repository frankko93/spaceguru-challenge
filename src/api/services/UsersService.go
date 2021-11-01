package services

import (
	"context"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/dao"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
)

type MyCustomClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func CreateUser(ctx context.Context, newUser models.User) (models.User, apierrors.ApiError) {

	var apiErr apierrors.ApiError
	var respUser models.User

	tags := utils.BuildTags(utils.Tags{
		"event":  "CreateUser",
		"source": "UsersService",
	})

	respUser, apiErr = dao.InsertUser(ctx, newUser)

	if apiErr != nil {
		log.Println("Error insert User", tags)
		return respUser, apiErr
	}

	return respUser, nil

}

func Login(ctx context.Context, login models.Login) (string, apierrors.ApiError) {

	var err error

	tags := utils.BuildTags(utils.Tags{
		"event":  "Login",
		"source": "UsersService",
	})

	respUser, apiErr := dao.SearchUser(ctx, login.Email)
	if apiErr != nil || (respUser.Email == "" && respUser.Password == "") {
		log.Println("Error validate User", apiErr, tags)
		return "", apierrors.NewBadRequestApiError("Error validate User")
	}
	if login.Password != respUser.Password {
		log.Println("Error: User or Password incorrect", apiErr, tags)
		return "", apierrors.NewBadRequestApiError("Error: User or Password incorrect")
	}

	signedToken, err := utils.CreateToken(login.Email)
	if err != nil {
		log.Println("Error creating token", err, tags)
		return "", apierrors.NewInternalServerApiError("Error creating token", err)
	}

	return signedToken, nil

}

func SearchUsersDrivers(ctx context.Context, params models.DriversSearchParams) (models.Users, apierrors.ApiError) {

	var apiErr apierrors.ApiError
	var respUserDao []models.User
	var respUsers models.Users
	var count = 0

	tags := utils.BuildTags(utils.Tags{
		"event":  "CreateFavoritesUser",
		"source": "UsersService",
	})

	respUserDao, count, apiErr = dao.SearchUsersDrivers(ctx, params)

	if apiErr != nil {
		log.Println(apiErr.Message(), tags)
		return respUsers, apierrors.NewInternalServerApiError("error search Users Drivers", apiErr)
	}
	respUsers.Data = respUserDao
	respUsers.Total = int64(count)
	respUsers.PageSize = params.PageSize
	respUsers.Page = params.Page
	respUsers.TotalPages = 1
	if count > 0 {
		respUsers.TotalPages = (int64(count) / params.PageSize)
	}
	if (int64(count) % params.PageSize) != 0 {
		respUsers.TotalPages = respUsers.TotalPages + 1
	}
	return respUsers, nil
}
