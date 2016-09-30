package esdatapub

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-oci8"
	"github.com/xtracdev/orapub"
	"os"
	"strings"
)

var connectStr string

func init() {
	var configErrors []string

	user := os.Getenv("FEED_DB_USER")
	if user == "" {
		configErrors = append(configErrors, "Configuration missing FEED_DB_USER env variable")
	}

	password := os.Getenv("FEED_DB_PASSWORD")
	if password == "" {
		configErrors = append(configErrors, "Configuration missing FEED_DB_PASSWORD env variable")
	}

	dbhost := os.Getenv("FEED_DB_HOST")
	if dbhost == "" {
		configErrors = append(configErrors, "Configuration missing FEED_DB_HOST env variable")
	}

	dbPort := os.Getenv("FEED_DB_PORT")
	if dbPort == "" {
		configErrors = append(configErrors, "Configuration missing FEED_DB_PORT env variable")
	}

	dbSvc := os.Getenv("FEED_DB_SVC")
	if dbSvc == "" {
		configErrors = append(configErrors, "Configuration missing FEED_DB_SVC env variable")
	}

	if len(configErrors) != 0 {
		log.Fatal(strings.Join(configErrors, "\n"))
	}

	connectStr = fmt.Sprintf("%s/%s@//%s:%s/%s",
		user, password, dbhost, dbPort, dbSvc)

}

func ProcessEventRecords() error {

	publisher := new(orapub.OraPub)
	err := publisher.Connect(connectStr, 5)
	if err != nil {
		log.Warnf("Unable to connect publisher reader")
		return err
	}

	err = publisher.InitializeProcessors()
	if err != nil {
		return err
	}

	publisher.ProcessEvents(true)

	return nil
}
