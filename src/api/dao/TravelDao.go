package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/frankko93/spaceguru-challenge/clients"
	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
	"github.com/jinzhu/gorm"
)

func InsertTravel(ctx context.Context, travel models.Travel) (models.Travel, apierrors.ApiError) {
	var err error
	var res *gorm.DB

	tags := utils.BuildTags(utils.Tags{
		"event":  "InsertTravel",
		"source": "UserDao",
	})

	db, err := clients.SpaceGuruDB()
	if err != nil {
		log.Println("Error connecting database", err, tags)
		return travel, apierrors.NewInternalServerApiError("Error connecting database", err)
	}

	startTime := time.Now()
	//insert user
	res = db.Table("travels").Create(&travel)
	endTime := time.Now()
	log.Println(fmt.Sprintf("query_elapsed_time: %f", endTime.Sub(startTime).Seconds()), tags)

	if res.Error != nil {
		log.Println("Error insert travel", res.Error, tags)
		return travel, apierrors.NewInternalServerApiError(fmt.Sprintf("Error insert travel %v", res.Error), err)
	}
	return travel, nil
}
