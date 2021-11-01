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
				"email":    "spaceguru@test.com",
				"name":     "franco",
				"surname":  "aballay",
				"type":     "driver",
				"password": "1234",
			},
			mocksType: "ok",
			want: want{
				Status: http.StatusCreated,
				Resp: map[string]interface{}{
					"id":       "c0f0f3ca-5265-4f21-b80f-08290ab256c6",
					"email":    "spaceguru@test.com",
					"name":     "franco",
					"password": "1234",
					"surname":  "aballay",
					"type":     "driver",
				},
			},
		},
		{
			name: "error existing email",
			body: controllers.RequestBody{
				"email":    "spaceguru@test.com",
				"name":     "franco",
				"surname":  "aballay",
				"type":     "driver",
				"password": "1234",
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
				res := controllers.PerformRequest("POST", "/users", tt.body, tt.headers, app.Router)

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
			ID:       "1",
			Name:     "Franco",
			Surname:  "Aballay",
			Type:     "admin",
			Password: "1234",
			Email:    "spaceguru@test.com",
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
				"email":    "spaceguru@test.com",
				"password": "1234",
			},
			mocksType: "ok",
			want: want{
				Status: http.StatusOK,
				Resp: map[string]interface{}{
					"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNwYWNlZ3VydUB0ZXN0LmNvbSJ9.bt1CY-lXLT03CtL-Q2K9xX0bhgHQHjibXqzHVBE14yY",
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
				res := controllers.PerformRequest("POST", "/users/login", tt.body, tt.headers, app.Router)

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
			ID:       "1",
			Name:     "Franco",
			Surname:  "Aballay",
			Type:     "admin",
			Password: "1234",
			Email:    "spaceguru@test.com",
		}
		dao.InsertUser(context.Background(), user)

	}
}

func Test_SearchDrivers(t *testing.T) {
	assert := assert.New(t)

	type want struct {
		Status int
		Resp   map[string]interface{}
	}
	tests := []struct {
		name      string
		query     string
		body      controllers.RequestBody
		headers   controllers.RequestHeaders
		want      want
		mocksType string
	}{
		{
			name:      "ok empty",
			query:     "status=free",
			body:      controllers.RequestBody{},
			mocksType: "",
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			want: want{
				Status: http.StatusOK,
				Resp: map[string]interface{}{
					"data":       []interface{}{},
					"page":       float64(1),
					"pageSize":   float64(10),
					"total":      float64(0),
					"totalPages": float64(1),
				},
			},
		},
		{
			name:      "ok",
			query:     "status=free",
			body:      controllers.RequestBody{},
			mocksType: "ok",
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			want: want{
				Status: http.StatusOK,
				Resp: map[string]interface{}{
					"data": []interface{}{
						map[string]interface{}{
							"createdAt": "0001-01-01T00:00:00Z",
							"email":     "spaceguru@test.com",
							"id":        "fd1416b1-6c04-44f1-a438-df3a2a997907",
							"name":      "Franco",
							"password":  "1234",
							"surname":   "Aballay",
							"type":      "driver",
							"updatedAt": "0001-01-01T00:00:00Z",
						},
					},
					"page":       float64(1),
					"pageSize":   float64(10),
					"total":      float64(1),
					"totalPages": float64(1),
				},
			},
		},
		{
			name:      "error params",
			query:     "status=ACTIVE2&page=1&pageSize=2",
			body:      controllers.RequestBody{},
			mocksType: "",
			headers: controllers.RequestHeaders{
				"x-auth": "mocked-session",
			},
			want: want{
				Status: http.StatusBadRequest,
				Resp: map[string]interface{}{
					"cause":   []interface{}{},
					"error":   "bad_request",
					"message": "Invalid Drivers Status",
					"status":  float64(400),
				},
			},
		},
	}

	for _, tt := range tests {
		createMocksSearchDrivers(tt.mocksType)
		t.Run(
			"Search_Property",
			func(t *testing.T) {
				res := controllers.PerformRequest("GET", "/users/drivers?"+tt.query, tt.body, tt.headers, app.Router)

				response := []byte(res.Body.String())

				var wantResp map[string]interface{}
				json.Unmarshal(response, &wantResp)

				if tt.want.Resp["data"] != nil { // casos de error que no viene data
					respData := tt.want.Resp["data"]
					if len(respData.([]interface{})) > 0 { // casos en data[]
						wantData := wantResp["data"]
						want := wantData.([]interface{})
						for index, prop := range respData.([]interface{}) {

							if prop.(map[string]interface{})["createdAt"] != nil {
								prop.(map[string]interface{})["createdAt"] = want[index].(map[string]interface{})["createdAt"] // Date automatico
							}
							if prop.(map[string]interface{})["updatedAt"] != nil {
								prop.(map[string]interface{})["updatedAt"] = want[index].(map[string]interface{})["updatedAt"] // Date automatico
							}
							if prop.(map[string]interface{})["id"] != nil {
								prop.(map[string]interface{})["id"] = want[index].(map[string]interface{})["id"] // Date automatico
							}

							respData.([]interface{})[index] = prop
						}
						tt.want.Resp["data"] = respData
					}

				}

				assert.Equal(tt.want.Status, res.Code)
				assert.Equal(tt.want.Resp, wantResp)
			},
		)
	}
}

func createMocksSearchDrivers(typeMock string) {
	db, _ := clients.SpaceGuruDB()
	switch typeMock {
	case "ok":
		commands.DeleteAllEntities(db)

		commands.DeleteAllEntities(db)

		user := models.User{
			Name:     "Franco",
			Surname:  "Aballay",
			Type:     "driver",
			Password: "1234",
			Email:    "spaceguru@test.com",
		}
		userDao, _ := dao.InsertUser(context.Background(), user)

		travel := models.Travel{
			UserID:    userDao.ID,
			VehicleID: "1",
			Status:    "in_progress",
			Route:     "aaaaaa",
		}
		dao.InsertTravel(context.Background(), travel)

	}
}
