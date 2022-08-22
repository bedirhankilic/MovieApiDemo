package dataaccesslayer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bedirhankilic/movieapicase/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInfo DbModel
var Db *gorm.DB

func createMysqlConnection() gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbInfo.User, DbInfo.Pass, DbInfo.Server, DbInfo.DbName)
	return mysql.Open(dsn)
}

func DbContext() (*gorm.DB, error) {
	var err error
	Db, err = gorm.Open(createMysqlConnection(), &gorm.Config{})
	return Db, err
}

func InitDb() error {
	ok := false
	if DbInfo.Server, ok = os.LookupEnv("DB_SERVER"); !ok {
		return errors.New("ENV ERROR")
	}
	if DbInfo.DbName, ok = os.LookupEnv("DB_NAME"); !ok {
		return errors.New("ENV ERROR")
	}
	if DbInfo.User, ok = os.LookupEnv("DB_USER"); !ok {
		return errors.New("ENV ERROR")
	}
	if DbInfo.Pass, ok = os.LookupEnv("DB_PASS"); !ok {
		return errors.New("ENV ERROR")
	}

	createDb()

	db, err := DbContext()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Movie{}, &models.User{})

	db.Create(&models.User{
		Name:     "Bedirhan",
		Email:    "recaibedirhankilic@gmail.com",
		Username: "bedirhankilic",
		Password: "123456",
	})
	return err
}

func createDb() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", DbInfo.User, DbInfo.Pass, DbInfo.Server))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+DbInfo.DbName)
	if err != nil {
		fmt.Printf("Error %s when creating DB\n", err)
		return
	}
	db.Close()
	no, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error %s when fetching rows", err)
		return
	}
	fmt.Printf("rows affected %d\n", no)
}
