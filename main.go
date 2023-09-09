package main

import (
	"fmt"
	"strings"
	"sustain/database"
	"sustain/handlers"
	"sustain/logs"
	"sustain/repositories"
	"sustain/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	db, err := database.CreateDB()
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(db)

	userRepositoryDB := repositories.NewUserRespositoryDB(db)
	userService := services.NewUserService(userRepositoryDB)
	userHandler := handlers.NewUserHandler(userService)

	groupRepositoryDB := repositories.NewGroupRespositoryDB(db)
	groupService := services.NewGroupService(groupRepositoryDB)
	groupHandler := handlers.NewGroupHandler(groupService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max",
	}))

	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"port": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(viper.GetString("app.port"))
			},
			"msg": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				message := fmt.Sprintf("Response body: %s", c.Response().Body())
				return output.WriteString(message)
			},
		},
		Format:     "[${time}] Status: ${status} - Medthod: ${method} API: ${path} Port:${port}\n Message: ${msg}\n\n",
		TimeFormat: "2 Jan 2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Post("/users/login", userHandler.Login)
	app.Post("/users/register", userHandler.Registers)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/token", userHandler.TokenCheck)
	app.Post("/group/create", groupHandler.CreateGroup)
	app.Post("/group/add", groupHandler.AddUserToGroup)
	app.Delete("/group/delete/:group_id", groupHandler.DeleteGroup)
	app.Post("/group/delete/user", groupHandler.DeleteUserGroup)
	app.Put("/group/update", groupHandler.UpdateGroup)

	logs.Info("Product service started at port " + viper.GetString("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
