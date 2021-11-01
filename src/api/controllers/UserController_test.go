package controllers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/frankko93/spaceguru-challenge/clients"
	"github.com/frankko93/spaceguru-challenge/commands"
	"github.com/frankko93/spaceguru-challenge/src/api/app"
	"github.com/frankko93/spaceguru-challenge/src/api/controllers"
	"github.com/frankko93/spaceguru-challenge/src/api/dao"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	assert := assert.New(t)

	type want struct {
		Status int
		Resp   map[string]interface{}
	}
	tests := []struct {
		name      string
		body      controllers.RequestBody
		headers   controllers.RequestHeaders
		want      want
		mocksType string
	}{
		{
			name: "ok",
			body: controllers.RequestBody{
				"email": "mock@gmail.com",
			},
			mocksType: "ok",
			want: want{
				Status: http.StatusCreated,
				Resp: map[string]interface{}{
					"email":    "mock@gmail.com",
					"id":       "a7a579c2-d92a-423b-be7b-7b2c207bd595",
					"password": "mock@gmail.com",
				},
			},
		},
		{
			name:      "error empty email",
			body:      controllers.RequestBody{},
			mocksType: "ok",
			want: want{
				Status: http.StatusBadRequest,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "bad_request",
					"message": "Invalid JSON Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag",
					"status":  float64(400),
				},
			},
		},
		{
			name: "error existing email",
			body: controllers.RequestBody{
				"email": "mock@gmail.com",
			},
			mocksType: "error_existing_email",
			want: want{
				Status: http.StatusBadRequest,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "bad_request",
					"message": "Error existing email",
					"status":  float64(400),
				},
			},
		},
	}

	for _, tt := range tests {
		createMocksCreateUser(tt.mocksType)
		t.Run(
			"Create_Users",
			func(t *testing.T) {
				res := controllers.PerformRequest("POST", "/v1/users", tt.body, tt.headers, app.Router)

				response := []byte(res.Body.String())

				var wantResp map[string]interface{}
				json.Unmarshal(response, &wantResp)

				if tt.want.Resp["id"] != nil {
					tt.want.Resp["id"] = wantResp["id"] // id uuid
				}

				assert.Equal(tt.want.Status, res.Code)
				assert.Equal(tt.want.Resp, wantResp)
			},
		)
	}

}

func createMocksCreateUser(typeMock string) {

	db, _ := clients.SpaceGuruDB()
	switch typeMock {
	case "error_existing_email":
		commands.DeleteAllEntities(db)

		user := models.User{
			Email: "mock@gmail.com",
		}
		dao.InsertUser(context.Background(), user)

	}
}

func Test_LoginUser(t *testing.T) {
	assert := assert.New(t)

	type want struct {
		Status int
		Resp   map[string]interface{}
	}
	tests := []struct {
		name      string
		body      controllers.RequestBody
		headers   controllers.RequestHeaders
		want      want
		mocksType string
	}{
		{
			name: "ok",
			body: controllers.RequestBody{
				"email":    "mock@gmaill",
				"password": "mock@gmaill",
			},
			mocksType: "ok",
			want: want{
				Status: http.StatusOK,
				Resp: map[string]interface{}{
					"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1vY2tAZ21haWxsIn0.F-vIcdbiYerP_Un0M2EHRm3zCE5vrojqj-hF6OPsVmY",
				},
			},
		},
		{
			name: "user not created",
			body: controllers.RequestBody{
				"email":    "mock@gmaill2",
				"password": "mock@gmaill2",
			},
			mocksType: "",
			want: want{
				Status: http.StatusBadRequest,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "bad_request",
					"message": "Error validate User",
					"status":  float64(400),
				},
			},
		},
	}

	for _, tt := range tests {
		createMocksLoginUser(tt.mocksType)
		t.Run(
			"Login_Users",
			func(t *testing.T) {
				res := controllers.PerformRequest("POST", "/v1/users/login", tt.body, tt.headers, app.Router)

				response := []byte(res.Body.String())

				var wantResp map[string]interface{}
				json.Unmarshal(response, &wantResp)

				assert.Equal(tt.want.Status, res.Code)
				assert.Equal(tt.want.Resp, wantResp)
			},
		)
	}

}

func createMocksLoginUser(typeMock string) {

	db, _ := clients.SpaceGuruDB()
	switch typeMock {
	case "ok":
		commands.DeleteAllEntities(db)

		user := models.User{
			Email: "mock@gmaill",
		}
		dao.InsertUser(context.Background(), user)

	}
}

func Test_CreateFavoritesUser(t *testing.T) {
	assert := assert.New(t)

	type want struct {
		Status int
		Resp   map[string]interface{}
	}
	tests := []struct {
		name      string
		body      controllers.RequestBody
		headers   controllers.RequestHeaders
		want      want
		mocksType string
	}{
		{
			name: "user_not_created",
			body: controllers.RequestBody{
				"propertyId": float64(1),
			},
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			mocksType: "",
			want: want{
				Status: http.StatusInternalServerError,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "internal_server_error",
					"message": "Error: could not get user",
					"status":  float64(500),
				},
			},
		},
		{
			name: "ok",
			body: controllers.RequestBody{
				"propertyId": float64(1),
			},
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			mocksType: "ok",
			want: want{
				Status: http.StatusCreated,
				Resp: map[string]interface{}{
					"id":         "2f888bb5-c457-4089-955f-d7e2e983fce1",
					"propertyID": float64(1),
					"userID":     "2c965f4f-9f7e-4da6-a448-6ecfc03e4f69",
				},
			},
		},
		{
			name: "unauthorized",
			body: controllers.RequestBody{
				"propertyId": float64(1),
			},
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session2",
			},
			mocksType: "",
			want: want{
				Status: http.StatusUnauthorized,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "unauthorized",
					"message": "Error validate x-auth session",
					"status":  float64(401),
				},
			},
		},
		{
			name: "bad_request",
			body: controllers.RequestBody{
				"propertyId": "a",
			},
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			mocksType: "not_user",
			want: want{
				Status: http.StatusBadRequest,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "bad_request",
					"message": "Invalid JSON",
					"status":  float64(400),
				},
			},
		},
	}

	for _, tt := range tests {
		createMocksCreateFavoritesUser(tt.mocksType)
		t.Run(
			"Create_FavoritesUser",
			func(t *testing.T) {
				res := controllers.PerformRequest("POST", "/v1/users/me/favorites", tt.body, tt.headers, app.Router)

				response := []byte(res.Body.String())

				var wantResp map[string]interface{}
				json.Unmarshal(response, &wantResp)

				if tt.want.Resp["userID"] != nil {
					tt.want.Resp["userID"] = wantResp["userID"] // id uuid
				}
				if tt.want.Resp["id"] != nil {
					tt.want.Resp["id"] = wantResp["id"] // id uuid
				}
				assert.Equal(tt.want.Status, res.Code)
				assert.Equal(tt.want.Resp, wantResp)
			},
		)
	}

}

func createMocksCreateFavoritesUser(typeMock string) {

	db, _ := clients.SpaceGuruDB()
	switch typeMock {
	case "ok":
		commands.DeleteAllEntities(db)

		user := models.User{
			Email: "mocked-session",
		}
		dao.InsertUser(context.Background(), user)

		property := models.Property{
			Title:       "Apartamento cerca a la estación de transmilenio",
			Description: "Apartamento con 3 cuartos y 2 baños localizado a 100 metros de la avenida caracas en la zona de Chapinero",
			Location: models.Location{
				Longitude: -74.0665887,
				Latitude:  4.6371593,
			},
			Pricing: models.Pricing{
				SalePrice:         450000000,
				AdministrativeFee: 250000,
			},
			PropertyType: "APARTMENT",
			Bedrooms:     6,
			Bathrooms:    2,
			ParkingSpots: 1,
			Area:         60,
			Photos: []string{
				"https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg",
				"https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg",
			},
		}

		dao.InsertProperty(context.Background(), property)
	}
}

func Test_SearchFavoritesUser(t *testing.T) {
	assert := assert.New(t)

	type want struct {
		Status int
		Resp   map[string]interface{}
	}
	tests := []struct {
		name      string
		body      controllers.RequestBody
		headers   controllers.RequestHeaders
		want      want
		mocksType string
	}{
		// {
		// 	name: "Invalid user",
		// 	body: controllers.RequestBody{},
		// 	headers: controllers.RequestHeaders{
		// 		"x-auth": "mocked-session2",
		// 	},
		// 	mocksType: "",
		// 	want: want{
		// 		Status: http.StatusUnauthorized,
		// 		Resp: map[string]interface{}{
		// 			"cause":   []interface{}{},
		// 			"error":   "unauthorized",
		// 			"message": "Error validate x-auth session",
		// 			"status":  float64(401),
		// 		},
		// 	},
		// },
		// {
		// 	name: "User not created",
		// 	body: controllers.RequestBody{},
		// 	headers: controllers.RequestHeaders{
		// 		"x-auth": "mocked-session",
		// 	},
		// 	mocksType: "",
		// 	want: want{
		// 		Status: http.StatusInternalServerError,
		// 		Resp: map[string]interface{}{
		// 			"cause":   []interface{}{},
		// 			"error":   "internal_server_error",
		// 			"message": "Error: could not get user",
		// 			"status":  float64(500),
		// 		},
		// 	},
		// },
		// {
		// 	name: "ok-empty",
		// 	body: controllers.RequestBody{},
		// 	headers: controllers.RequestHeaders{
		// 		"x-auth": "mocked-session",
		// 	},
		// 	mocksType: "ok-empty",
		// 	want: want{
		// 		Status: http.StatusOK,
		// 		Resp: map[string]interface{}{
		// 			"data":       []interface{}{},
		// 			"page":       float64(1),
		// 			"pageSize":   float64(10),
		// 			"total":      float64(0),
		// 			"totalPages": float64(1),
		// 		},
		// 	},
		// },
		// {
		// 	name: "ok",
		// 	body: controllers.RequestBody{},
		// 	headers: controllers.RequestHeaders{
		// 		"x-auth": "mocked-session",
		// 	},
		// 	mocksType: "ok",
		// 	want: want{
		// 		Status: http.StatusOK,
		// 		Resp: map[string]interface{}{
		// 			"data": []interface{}{
		// 				map[string]interface{}{
		// 					"area":        float64(60),
		// 					"bathrooms":   float64(2),
		// 					"bedrooms":    float64(6),
		// 					"createdAt":   "2021-10-25T21:52:41.848561-03:00",
		// 					"description": "Apartamento con 3 cuartos y 2 baños localizado a 100 metros de la avenida caracas en la zona de Chapinero",
		// 					"id":          float64(1),
		// 					"location": map[string]interface{}{
		// 						"latitude":  float64(4.6371593),
		// 						"longitude": float64(-74.0665887),
		// 					},
		// 					"parkingSpots": float64(1),
		// 					"photos":       interface{}(nil),
		// 					"pricing": map[string]interface{}{
		// 						"administrativeFee": float64(250000),
		// 						"salePrice":         float64(450000000),
		// 					},
		// 					"propertyType": "APARTMENT",
		// 					"status":       "",
		// 					"title":        "Apartamento cerca a la estación de transmilenio",
		// 					"updatedAt":    "2021-10-25T21:52:41.848561-03:00",
		// 				},
		// 			},
		// 			"page":       float64(1),
		// 			"pageSize":   float64(10),
		// 			"total":      float64(1),
		// 			"totalPages": float64(1),
		// 		},
		// 	},
		// },
	}

	for _, tt := range tests {

		createMocksSearchFavoritesUser(tt.mocksType)
		t.Run(
			"Search_FavoritesUser",
			func(t *testing.T) {
				res := controllers.PerformRequest("GET", "/v1/users/me/favorites", tt.body, tt.headers, app.Router)

				response := []byte(res.Body.String())

				var wantResp map[string]interface{}
				json.Unmarshal(response, &wantResp)

				if tt.want.Resp["data"] != nil { // casos de error que no viene data
					respData := tt.want.Resp["data"]
					if len(respData.([]interface{})) > 0 { // casos en data[]
						resp := respData.([]interface{})[0]

						wantData := wantResp["data"]
						want := wantData.([]interface{})[0]

						if resp.(map[string]interface{})["createdAt"] != nil {
							resp.(map[string]interface{})["createdAt"] = want.(map[string]interface{})["createdAt"] // Date automatico
						}
						if resp.(map[string]interface{})["updatedAt"] != nil {
							resp.(map[string]interface{})["updatedAt"] = want.(map[string]interface{})["updatedAt"] // Date automatico
						}

						respData.([]interface{})[0] = resp
						tt.want.Resp["data"] = respData

						wantData.([]interface{})[0] = want
						wantResp["data"] = wantData
					}

				}

				assert.Equal(tt.want.Status, res.Code)
				assert.Equal(tt.want.Resp, wantResp)
			},
		)
	}

}

func createMocksSearchFavoritesUser(typeMock string) {

	db, _ := clients.SpaceGuruDB()
	switch typeMock {
	case "ok":
		commands.DeleteAllEntities(db)

		user := models.User{
			Email: "mocked-session",
		}
		userDao, _ := dao.InsertUser(context.Background(), user)

		property := models.Property{
			Title:       "Apartamento cerca a la estación de transmilenio",
			Description: "Apartamento con 3 cuartos y 2 baños localizado a 100 metros de la avenida caracas en la zona de Chapinero",
			Location: models.Location{
				Longitude: -74.0665887,
				Latitude:  4.6371593,
			},
			Pricing: models.Pricing{
				SalePrice:         450000000,
				AdministrativeFee: 250000,
			},
			PropertyType: "APARTMENT",
			Bedrooms:     6,
			Bathrooms:    2,
			ParkingSpots: 1,
			Area:         60,
			Photos: []string{
				"https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg",
				"https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg",
			},
		}

		dao.InsertProperty(context.Background(), property)

		favorite := models.UserFavorites{
			PropertyID: 1,
			UserID:     userDao.ID,
		}

		dao.InsertUserFavorite(context.Background(), favorite)

		// case "ok-empty":
		// 	commands.DeleteAllEntities(db)

		// 	user := models.User{
		// 		Email: "mocked-session",
		// 	}
		// 	dao.InsertUser(context.Background(), user)

		// 	property := models.Property{
		// 		Title:       "Apartamento cerca a la estación de transmilenio",
		// 		Description: "Apartamento con 3 cuartos y 2 baños localizado a 100 metros de la avenida caracas en la zona de Chapinero",
		// 		Location: models.Location{
		// 			Longitude: -74.0665887,
		// 			Latitude:  4.6371593,
		// 		},
		// 		Pricing: models.Pricing{
		// 			SalePrice:         450000000,
		// 			AdministrativeFee: 250000,
		// 		},
		// 		PropertyType: "APARTMENT",
		// 		Bedrooms:     6,
		// 		Bathrooms:    2,
		// 		ParkingSpots: 1,
		// 		Area:         60,
		// 		Photos: []string{
		// 			"https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg",
		// 			"https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg",
		// 		},
		// 	}
		// 	dao.InsertProperty(context.Background(), property)

	}
}
