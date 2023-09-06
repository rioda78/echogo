package database

import (
	"echogo/config"
	"echogo/migrasi"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	conf := config.GetConfig()
	var connectionString = ""
	if conf.Rdms == "mysql" {
		//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		connectionString = conf.DbUsername + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.DbName
	} else {
		connectionString = "host=" + conf.DbHost + " user=" + conf.DbUsername + " password=" + conf.DbPassword + " dbname=" + conf.DbName + " port=" + conf.DbPort + " sslmode=disable"
	}

	if conf.Rdms == "mysql" {
		db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err != nil {
			panic("Could not connect to the database")
		}
		migrasi.Migrate(db)
		return db
	} else {
		db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err != nil {
			panic("Could not connect to the database")
		}
		migrasi.Migrate(db)
		return db
	}

	//db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
