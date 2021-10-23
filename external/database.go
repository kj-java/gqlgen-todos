package external

import (
	"github.com/maruware/gqlgen-todos/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=user password=mypassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}