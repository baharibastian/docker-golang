package main

import (
	"net/http"
	"os"

	"github.com/ecojuntak/gorb/middlewares"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/ecojuntak/gorb/database"
	"github.com/urfave/cli"

	"github.com/spf13/viper"

	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/repositories"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func onError(err error, failedMessage string) {
	if err != nil {
		logrus.Errorln(failedMessage)
		logrus.Errorln(err)
	}
}

func loadConfig() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()

	return err
}

func runServer(db *gorm.DB) {
	r := LoadRouter(db)
	corsOption := middlewares.CorsMiddleware()
	redisClient = database.BuildRedisInstance()

	if (!IsMigrated()) {
		err := database.Migrate(db)
		onError(err, "Failed to migrate database schema")

		set := redisClient.Set("is_migrated", "migrated", 0).Err()
		if set != nil {
			logrus.Errorln(err)
		}
	}	

	logrus.Infoln("Server run on " + getAddress())

	http.ListenAndServe(":8000", handlers.CORS(corsOption[0], corsOption[1], corsOption[2])(r))
}

func main() {
	err := loadConfig()
	if err != nil {
		logrus.Errorln(err)
	}

	db, err := database.InitDatabase()
	err = db.DB().Ping()

	if err != nil {
		logrus.Errorln(err)
	}

	defer db.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "GORB"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		// {
		// 	Name:        "migrate",
		// 	Description: "Run database migration",
		// 	Action: func(c *cli.Context) error {
		// 		err = database.Migrate(db)
		// 		onError(err, "Failed to migrate database schema")

		// 		return err
		// 	},
		// },
		// {
		// 	Name:        "seed",
		// 	Description: "Run database seeder",
		// 	Action: func(c *cli.Context) error {
		// 		err = database.RunSeeder(db)
		// 		onError(err, "Failed to generate fake data")

		// 		return err
		// 	},
		// },
		{
			Name:        "start",
			Description: "Start REST API Server",
			Action: func(c *cli.Context) error {
				runServer(db)
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		logrus.Fatalln(err)
	}
}

func LoadRouter(db *gorm.DB) (r *mux.Router) {
	userRepo := repositories.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo)

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/users", userController.Resources).Methods("GET", "POST")
	v1.HandleFunc("/users/{id}", userController.Resources).Methods("GET", "PATCH", "DELETE")

	r.Use(middlewares.LoggerMidldlware)

	return
}

func getAddress() string {
	return viper.GetString("HOST") + ":" + viper.GetString("PORT")
}

func IsMigrated() bool {
	client := database.BuildRedisInstance()

	val, err := client.Get("is_migrated").Result()
	is_migrated := false

	if err != nil {
		panic(err)
	} else {
		if val == "migrated" {
			is_migrated = true
		}
	}

	return is_migrated
}