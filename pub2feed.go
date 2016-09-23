package esdatapub

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-oci8"
	"github.com/xtracdev/orapub"
	"os"
	"strings"
	"time"
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

func connectToDB(connectStr string) (*sql.DB, error) {
	db, err := sql.Open("oci8", connectStr)
	if err != nil {
		log.Warnf("Error connecting to oracle: %s", err.Error())
		return nil, err
	}

	//Are we really in an ok state for starters?
	err = db.Ping()
	if err != nil {
		log.Infof("Error connecting to oracle: %s", err.Error())
		return nil, err
	}

	return db, nil
}



func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

	for {

		polledEventsSpec, err := publisher.PollEvents()
		if err != nil {
			log.Warnf("Error polling for events: %s", err.Error())
			continue
		}

		log.Infof("Process %d events", len(polledEventsSpec))

		for i := 0; i < len(polledEventsSpec); i += 100 {

			batch := polledEventsSpec[i:min(i+100, len(polledEventsSpec))]
			log.Infof("===> processing batch with starting index %d batch size %d", i, len(batch))

			for _, eventContext := range batch {

				e, err := publisher.RetrieveEventDetail(eventContext.AggregateId, eventContext.Version)
				if err != nil {
					log.Warnf("Error reading event to process (%v): %s", eventContext, err)
					continue
				}

				//TODO - make error codes available to interested users of OraPub
				publisher.ProcessEvent(e)
			}

			log.Infof("Deleting %d events", len(batch))
			err = publisher.DeleteProcessedEvents(batch)
			if err != nil {
				log.Warnf("Error cleaning up processed events: %s", err)
			}

		}

		if len(polledEventsSpec) == 0 {
			log.Infof("Nothing to do... time for a 5 second sleep")
			time.Sleep(5 * time.Second)
		}

	}
}
