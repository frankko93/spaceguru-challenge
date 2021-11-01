package services

import (
	"context"
	"log"

	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/dao"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
)

func CreateTravel(ctx context.Context, travel models.Travel) (models.Travel, apierrors.ApiError) {

	var apiErr apierrors.ApiError
	var respTravel models.Travel

	tags := utils.BuildTags(utils.Tags{
		"event":  "CreateFavoritesUser",
		"source": "TravelService",
	})

	respTravel, apiErr = dao.InsertTravel(ctx, travel)
	if apiErr != nil {
		log.Println("Error: create travel", apiErr, tags)
		return respTravel, apiErr
	}

	return respTravel, nil
}
