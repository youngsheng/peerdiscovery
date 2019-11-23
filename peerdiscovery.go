package peerdiscovery

import (
	"fmt"
	"github.com/schollz/peerdiscovery"
	"log"
	"time"
)

type(
	DataProcessor struct{
	// add fields here
	}
)

func (p *DataProcessor) PeerDicovers() (str string, err error){
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     -1,
		Payload:   []byte("ubuntu"),
		Delay:     500 * time.Millisecond,
		TimeLimit: 10 * time.Second,
		/*
			Notify: func(d peerdiscovery.Discovered) {
				log.Println(d)
			},
		*/
	})

	ipAddress := []string{}
	if err != nil {
		log.Fatal(err)
	} else {
		if len(discoveries) > 0 {
			fmt.Printf("\nFound %d other computers\n", len(discoveries))
			for i, d := range discoveries {
				fmt.Printf("%d) '%s' with payload '%s'\n", i, d.Address, d.Payload)
				ipAddress = append(ipAddress, d.Address)
			}
		} else {
			fmt.Println("\nFound no devices. You need to run this on another computer at the same time.")
		}
	}

	ipAddressStr := ""
	for _, ip := range ipAddress {
		ipAddressStr += (ip + ",")
	}

	return ipAddressStr, nil
}
