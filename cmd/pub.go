package main

import (
	ad "github.com/xtracdev/es-atom-data"
	dp "github.com/xtracdev/es-data-pub"
	"github.com/xtracdev/orapub"
	"log"
)

func main() {
	//Instantiate and register the event processors
	processor := ad.NewESAtomPubProcessor()
	orapub.RegisterEventProcessor("es-atom-data",processor)

	//Loop until death
	err := dp.ProcessEventRecords()
	if err != nil {
		log.Fatal(err.Error())
	}
}
