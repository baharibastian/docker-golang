package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func InitDatabase() (db *gorm.DB, err error) {
	dbDriver := "mysql"
	var connectionString string

	if dbDriver == "mysql" {
		connectionString = buildMysqlConnectionString()
	} else if dbDriver == "postgres" {
		connectionString = buildPostgresqlConnectionString()
	} else if dbDriver == "sqlite3" {
		connectionString = buildSqliteConnectionString()
	}

	logrus.Infoln("Connection String " + connectionString)

	db, err = openConnection(dbDriver, connectionString)

	return
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func buildMysqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"db")

	return
}

func buildPostgresqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_POSTGRES_SSL_MODE"))

	return
}

func buildSqliteConnectionString() (connectionString string) {
	connectionString = viper.GetString("DB_NAME")
	return
}

func BuildRedisInstance() (*redis.Client) {
	redisdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	return redisdb
}
