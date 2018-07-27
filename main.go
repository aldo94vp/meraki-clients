package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aldo94vp/meraki-clients/meraki"
	"github.com/aldo94vp/meraki-clients/structs"
)

func main() {
	var config structs.Config

	// Unmarshal config data
	file, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(file, &config)

	networks := meraki.GetNetworks(config)
	var ssids = structs.SSIDs{}
	//fmt.Println(networks)
	for _, network := range networks {
		//fmt.Println(network)
		for _, ssid := range meraki.GetESSIDs(config, network.ID) {
			time.Sleep(300 * time.Millisecond)
			if ssid.Enabled && meraki.ReturnIndexSSID(ssid.Name, ssids) < 0 {
				ssids = append(ssids, ssid)
			}
		}
		devices := meraki.GetDevices(config, network.ID)
		for _, device := range devices {
			//fmt.Println(device)
			clients := meraki.GetClients(config, device.Serial)
			for _, client := range clients {
				fmt.Println(client)
				client = meraki.GetClient(config, network.ID, client.ID)
				index := meraki.ReturnIndexSSID(client.SSID, ssids)
				if index > -1 {
					ssids[index].Count++
				}
				fmt.Println(time.Now())
				time.Sleep(350 * time.Millisecond)
				fmt.Println(time.Now())
			}
		}
	}
	fmt.Println(ssids)

}
