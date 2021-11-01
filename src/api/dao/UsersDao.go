package dao

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/frankko93/spaceguru-challenge/clients"
	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func InsertUser(ctx context.Context, newUser models.User) (models.User, apierrors.ApiError) {
	var res *gorm.DB

	var userDB models.User

	tags := utils.BuildTags(utils.Tags{
		"event":  "InsertUser",
		"source": "UsersDao",
	})

	//validate if it exists
	userDB, _ = SearchUser(ctx, newUser.Email)
	if userDB.Email != "" {
		log.Println("Error existing email", tags)
		return newUser, apierrors.NewBadRequestApiError("Error existing email")
	}

	//id
	newUser.ID = uuid.NewV4().String()

	db, err := clients.SpaceGuruDB()
	if err != nil {
		log.Println("Error connecting database", err, tags)
		return newUser, apierrors.NewInternalServerApiError("Error connecting database", err)
	}

	startTime := time.Now()
	//insert user
	res = db.Table("users").Create(&newUser)
	endTime := time.Now()
	log.Println(fmt.Sprintf("query_elapsed_time: %f", endTime.Sub(startTime).Seconds()), tags)

	if res.Error != nil {
		log.Println("Error insert user", res.Error, tags)
		return newUser, apierrors.NewInternalServerApiError("Error insert user", res.Error)
	}

	return newUser, nil
}

func SearchUser(ctx context.Context, email string) (models.User, apierrors.ApiError) {
	var res *gorm.DB

	var user models.User

	tags := utils.BuildTags(utils.Tags{
		"event":  "SearchUser",
		"source": "UserDao",
	})

	db, err := clients.SpaceGuruDB()
	if err != nil {
		log.Println("Error connecting database", err, tags)
		return user, apierrors.NewInternalServerApiError("Error connecting database", err)
	}

	startTime := time.Now()
	//search users
	res = db.Table("users").Where("email = (?)", email).Find(&user)
	endTime := time.Now()
	log.Println(fmt.Sprintf("query_elapsed_time: %f", endTime.Sub(startTime).Seconds()), tags)

	if res.Error != nil && res.Error.Error() != "record not found" {
		log.Println("Error search user", res.Error, tags)
		return user, apierrors.NewInternalServerApiError("Error search user", res.Error)
	}

	return user, nil
}

func SearchUsersDrivers(ctx context.Context, params models.DriversSearchParams) ([]models.User, int, apierrors.ApiError) {
	var res *gorm.DB
	var users []models.User
	var count int

	tags := utils.BuildTags(utils.Tags{
		"event":  "SearchUsersDrivers",
		"source": "UsersDao",
	})

	db, err := clients.SpaceGuruDB()
	if err != nil {
		log.Println("Error connecting database", err, tags)
		return users, count, apierrors.NewInternalServerApiError("Error connecting database", err)
	}
	query, queryParams := createWhereDrivers(params)

	startTime := time.Now()
	//search users drivers
	res = db.Table("users").Debug().Joins("left join travels on travels.user_id = users.id").Where(query, queryParams...).Offset(params.PageSize * (params.Page - 1)).Limit(params.PageSize).Scan(&users)
	if res.Error != nil {
		log.Println("Error search users drivers", res.Error, tags)
		return users, count, apierrors.NewInternalServerApiError("Error search users drivers", res.Error)
	}
	//count properties
	res = db.Table("users").Joins("left join drivers on drivers.user_id = users.id").Where(query, queryParams...).Count(&count)
	endTime := time.Now()
	log.Println(fmt.Sprintf("query_elapsed_time: %f", endTime.Sub(startTime).Seconds()), tags)

	if res.Error != nil {
		log.Println("Error search users drivers count ", res.Error, tags)
		return users, count, apierrors.NewInternalServerApiError(fmt.Sprintf("Error search users drivers %v", res.Error), err)
	}
	return users, count, nil
}

func createWhereDrivers(params models.DriversSearchParams) (string, []interface{}) {

	query := "users.type = ? "
	whereParams := []interface{}{"driver"}

	if params.Status != "" {

		if params.Status == "free" {
			query = fmt.Sprintf(" %v AND status != (?) AND status != (?) ", query)
			whereParams = append(whereParams, "finished")
			whereParams = append(whereParams, "cancelled")
		} else {
			query = fmt.Sprintf(" %v AND status != (?) ", query)
			statusList := strings.Split(params.Status, ",")
			whereParams = append(whereParams, statusList)
		}
	}

	return query, whereParams
}
