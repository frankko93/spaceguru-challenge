package controllers

import (
	"log"
	"net/http"

	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/services"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
	"github.com/gin-gonic/gin"
)

func CreateTravel(c *gin.Context) {

	ctx := c.Request.Context()
	var err error
	var travel models.Travel

	tags := utils.BuildTags(utils.Tags{
		"event":  "CreateTravel",
		"source": "TravelController",
	})

	if err = c.BindJSON(&travel); err != nil {
		log.Println("Invalid JSON", err, tags)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("Invalid JSON"))
		return
	}

	userFavorite, apiErr := services.CreateTravel(ctx, travel)
	if apiErr != nil {
		c.JSON(http.StatusInternalServerError, apiErr)
		return
	}
	c.JSON(http.StatusCreated, userFavorite)

}
