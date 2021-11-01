{
	"info": {
		"_postman_id": "60a1cbd1-6e59-40b5-8431-4fb91ca84752",
		"name": "Space Guru",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Create User",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"email\": \"spaceguru@test.com\",\n\t\"name\": \"franco\",\n\t\"surname\": \"aballay\",\n    \"status\": \"active\",\n    \"type\": \"driver\",\n    \"password\": \"1234\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/users",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"users"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create Travel",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "x-auth",
						"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNwYWNlZ3VydUB0ZXN0LmNvbSJ9.bt1CY-lXLT03CtL-Q2K9xX0bhgHQHjibXqzHVBE14yY",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"user_id\": \"841c537f-16d2-4556-92a6-e0c6475d3871\",\n\t\"vehicle_id\": \"1\",\n\t\"status\": \"in_progress\",\n    \"route\": \"aaaaaaaa\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/users/travel",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"users",
						"travel"
					]
				}
			},
			"response": []
		},
		{
			"name": "Search Drivers",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [
					{
						"key": "x-auth",
						"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNwYWNlZ3VydUB0ZXN0LmNvbSJ9.bt1CY-lXLT03CtL-Q2K9xX0bhgHQHjibXqzHVBE14yY",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"title\": \"Apartamento cerca a la estación de transmilenio\",\n\t\"location\": {\n\t\t\"longitude\": -74.0665887,\n\t\t\"latitude\": 4.6371593\n\t},\n\t\"pricing\": {\n\t\t\"salePrice\": 450000000\n\t},\n\t\"propertyType\": \"HOUSE\",\n\t\"bedrooms\": 3,\n\t\"bathrooms\": 2,\n\t\"parkingSpots\": 1,\n\t\"area\": 60,\n\t\"photos\": [\n\t\t\"https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg\",\n\t\t\"https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg\"\n\t]\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/users/drivers?status=finished",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"users",
						"drivers"
					],
					"query": [
						{
							"key": "status",
							"value": "finished"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"email\": \"spaceguru@test.com\",\n\t\"password\": \"1234\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/users/login",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"users",
						"login"
					]
				}
			},
			"response": []
		}
