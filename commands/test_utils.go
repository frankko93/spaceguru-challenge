package commands

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type table struct {
	Name  string
	Model interface{}
}

type Vehicles struct {
	ID          int64     `json:"id"`
	Type        string    `json:"type" binding:"required"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Travels struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	VehicleID float64   `json:"vehicle_id"`
	Status    string    `json:"status"`
	Route     string    `json:"route"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
type Users struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Type      string    `json:"type" binding:"required"` // admin/driver
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

var Tables = []table{
	{Name: "users", Model: Users{}},
	{Name: "travels", Model: Travels{}},
	{Name: "vehicles", Model: Vehicles{}},
}

func DeleteAllEntities(db *gorm.DB) {
	db.Delete(Users{})
	db.Delete(Travels{})
	db.Delete(Vehicles{})
}

func CreateTables(db *gorm.DB) error {
	for _, t := range Tables {
		db.DropTable(t.Model)
		if db.Error != nil {
			fmt.Errorf("CreateTables - error delete tabla", db.Error)
		}

		if !db.HasTable(t.Model) {
			db.CreateTable(t.Model)
			if db.Error != nil {
				fmt.Errorf("CreateTables - error crear tabla", db.Error)
				return db.Error
			}
			fmt.Println(t.Name + " created")
		} else {
			fmt.Println("CreateTables - " + t.Name + " tabla creada perro")
		}
	}
	return nil
}
