package server

import (
	"cellular-data-tracker/client"
	"cellular-data-tracker/repository"
	"cellular-data-tracker/service"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Configuration struct {
	Port         int
	Database     Database
	DataProvider DataProvider
}

type Database struct {
	ConnectionUrl string
}

type DataProvider struct {
	Email    string
	Password string
	ApiToken string
	Numbers  []string
}

type Application struct {
	fiber               *fiber.App
	cellularDataService service.CellularDataService
	infoLog             *log.Logger
	errorLog            *log.Logger
}

func Start(configuration *Configuration) {
	db := openDbConnection(configuration.Database.ConnectionUrl)
	defer db.Close()

	app := &Application{
		fiber:    getConfiguredFiber(),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		cellularDataService: service.NewCellularDataService(
			client.NewCellularDataUsageStatisticsClient(),
			repository.NewCellularDataStatisticsRepository(10*time.Second, db),
			repository.NewPhoneRepository(10*time.Second, db),
		),
	}
	app.setupHandlers()

	err := app.fiber.Listen(fmt.Sprintf(":%d", configuration.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func openDbConnection(connectionUrl string) *sql.DB {
	db, err := sql.Open("postgres", connectionUrl)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getConfiguredFiber() *fiber.App {
	engine := html.New("./static", ".html")

	//middleware
	return fiber.New(fiber.Config{Views: engine})
}
