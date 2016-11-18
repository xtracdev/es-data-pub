package esdatapub

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/xtracdev/orapub"
)

var connectStr string

func init() {
	var configErrors []string

	user := os.Getenv("DB_USER")
	if user == "" {
		configErrors = append(configErrors, "Configuration missing DB_USER env variable")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		configErrors = append(configErrors, "Configuration missing DB_PASSWORD env variable")
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		configErrors = append(configErrors, "Configuration missing DB_HOST env variable")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		configErrors = append(configErrors, "Configuration missing DB_PORT env variable")
	}

	dbSvc := os.Getenv("DB_SVC")
	if dbSvc == "" {
		configErrors = append(configErrors, "Configuration missing DB_SVC env variable")
	}

	if len(configErrors) != 0 {
		log.Fatal(strings.Join(configErrors, "\n"))
	}

	connectStr = fmt.Sprintf("%s/%s@//%s:%s/%s",
		user, password, dbhost, dbPort, dbSvc)

}

func ProcessEventRecords() error {
	maxTriesEnv := os.Getenv("DB_MAX_TRIES")
	var maxTries int
	if maxTries, err := strconv.Atoi(maxTriesEnv); err == nil {
		log.Debugf("max database reconnections set to %v", maxTries)
	} else {
		maxTries = 5
	}

	publisher := new(orapub.OraPub)
	err := publisher.Connect(connectStr, maxTries)
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
