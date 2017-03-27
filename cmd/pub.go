package main

import (
	ad "github.com/xtracdev/es-atom-data"
	dp "github.com/xtracdev/es-data-pub"
	"github.com/xtracdev/orapub"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: go run %s <num event pub loops to process>")
	}

	numEvents, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	//Instantiate and register the event processors
	processor := ad.NewESAtomPubProcessor()
	orapub.RegisterEventProcessor("es-atom-data",processor)

	//Loop until death
	err = dp.ProcessNEvents(numEvents)
	if err != nil {
		log.Fatal(err.Error())
	}
}
