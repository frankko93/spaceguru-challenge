package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/config"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/services"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var apiErr apierrors.ApiError
	var newUser models.User

	tags := utils.BuildTags(utils.Tags{
		"event":  "CreateUser",
		"source": "UserController",
	})

	if err := c.BindJSON(&newUser); err != nil {
		log.Println("Invalid JSON", err, tags)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError(fmt.Sprintf("Invalid JSON %v", err)))
		return
	}

	respProperty, apiErr := services.CreateUser(ctx, newUser)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusCreated, respProperty)
}

func LoginUser(c *gin.Context) {

	ctx := c.Request.Context()
	var login models.Login

	tags := utils.BuildTags(utils.Tags{
		"event":  "LoginUser",
		"source": "UserController",
	})

	if err := c.BindJSON(&login); err != nil {
		log.Println("Invalid JSON", err, tags)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError(fmt.Sprintf("Invalid JSON %v", err)))
		return
	}

	token, apiErr := services.Login(ctx, login)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func SearchUsersDrivers(c *gin.Context) {

	ctx := c.Request.Context()

	tags := utils.BuildTags(utils.Tags{
		"event":  "SearchUsersDrivers",
		"source": "UserController",
	})

	searchParams, apiErr := validateSearchDriversParams(*c, tags)
	if apiErr != nil {
		log.Println(fmt.Sprintf("Invalid params :%v", apiErr), tags)
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}
	_, exist := c.Get("user")
	if !exist {
		log.Println("Error with session", tags)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("Error with session"))
		return
	}

	usersDrivers, apiErr := services.SearchUsersDrivers(ctx, searchParams)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, usersDrivers)

}

func validateSearchDriversParams(c gin.Context, tags []string) (models.DriversSearchParams, apierrors.ApiError) {

	var err error
	var searchParams models.DriversSearchParams

	status, existStatus := c.GetQuery("status")
	if existStatus {
		if !utils.Contains(config.DriversStatusSearch, status) {
			return searchParams, apierrors.NewBadRequestApiError("Invalid Drivers Status")
		}
		searchParams.Status = status
	}

	page, existPage := c.GetQuery("page")
	if existPage {
		searchParams.Page, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			return searchParams, apierrors.NewBadRequestApiError("Invalid Page")
		}
	} else {
		searchParams.Page = 1
	}

	pageSize, existPageSize := c.GetQuery("pageSize")
	if existPageSize {
		searchParams.PageSize, err = strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return searchParams, apierrors.NewBadRequestApiError("Invalid PageSize")
		}
	} else {
		searchParams.PageSize = 10
	}

	return searchParams, nil
}
