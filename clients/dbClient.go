package clients

import (
	"fmt"
	"log"

	//base de datos en memoria

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	//driver mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/frankko93/spaceguru-challenge/commands"
	"github.com/jinzhu/gorm"
)

var spaceguruDB *gorm.DB

func InitDB() (*gorm.DB, error) {

	var err error

	log.Println("init DB")
	spaceguruDB, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		log.Println("error al querer conectarnos a la db en memoria", err)
		return spaceguruDB, err

	}
	err = commands.CreateTables(spaceguruDB)
	if err != nil {
		log.Println("error al crear tablas de db en memoria", err)
		return spaceguruDB, err

	}
	return spaceguruDB, err
}

func SpaceGuruDB() (*gorm.DB, error) {

	var err error

	stats := spaceguruDB.DB().Stats()

	if stats.OpenConnections >= 50 {
		err := fmt.Errorf("number of connections exceeded: %v", stats.OpenConnections)
		return nil, err
	}

	err = spaceguruDB.DB().Ping()
	if err != nil {
		log.Println("error connect wdwadmin db", err)
		return nil, err
	}
	return spaceguruDB, nil
}
