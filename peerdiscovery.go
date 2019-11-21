package peerdiscovery

import (
	"fmt"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/schollz/peerdiscovery"
	"log"
	"time"
)

const channelName = "youngspace.xyz/go/peerdiscovery"

type PeerDiscoveryPlugin struct{}

var _ flutter.Plugin = &PeerDiscoveryPlugin{} // compile-time type check

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getBatteryLevel' method.
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#MethodChannel
//
// The plugin is using a MethodChannel through the StandardMethodCodec.
//
// You can also use the more basic BasicMessageChannel, which supports basic,
// asynchronous message passing using a custom message codec.
// You can also use the specialized BinaryCodec, StringCodec, and JSONMessageCodec
// struct, or create your own codec.
//
// The JSONMessageCodec was deliberately left out because in Go its better to
// decode directly to known structs.
func (PeerDiscoveryPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getBatteryLevel", handleGetBatteryLevel)
	channel.HandleFunc("peerDicovers", peerDicovers)
	return nil // no error
}

// handleGetBatteryLevel is called when the method getBatteryLevel is invoked by
// the dart code.
// The function returns a fake battery level, without raising an error.
//
// Supported return types of StandardMethodCodec codec are described in a table:
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#StandardMessageCodec
func handleGetBatteryLevel(arguments interface{}) (reply interface{}, err error) {
	batteryLevel := int32(55) // Your platform-specific API
	return batteryLevel, nil
}

func peerDicovers(arguments interface{}) (reply interface{}, err error){
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


	return ipAddress, nil
}
