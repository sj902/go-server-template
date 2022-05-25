package main

import (
	"fmt"
	"os"
	"server/confs"
	v1 "server/handlers/v1"
	"server/middleware"
	"server/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ENV") == "prod" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	}
}

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Getenv("PORT")); err != nil {
		log.Errorf("Error: %v", err.Error())
		os.Exit(exitFail)
	}
}

// CMD: PORT=3000 go run main.go
func run(port string) error {

	confs.Init(port)
	cnf := confs.Conf()

	externalIdMap := make(map[string]string)

	server := v1.Server{
		ExtIdMap: externalIdMap,
	}

	log.Info("Starting server")

	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	// Add middlewares
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	app.Use(cors.New())

	app.Use(pprof.New())

	jwtMiddleware := middleware.NewCustomMiddleware(middleware.CustomMiddlewareConfig{})

	router.SetupRoutes(app, jwtMiddleware, server)

	return startListening(app, *cnf)
}

func startListening(app *fiber.App, cnf confs.Config) error {
	return app.Listen(fmt.Sprintf(":%s", cnf.ServerPort))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	log.Error(err)
	msg := "Internal Server Error"
	e, ok := err.(*fiber.Error)
	if ok && e.Code < 500 {
		code = e.Code
		msg = e.Error()
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).SendString(msg)
}
