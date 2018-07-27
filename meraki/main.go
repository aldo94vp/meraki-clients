package meraki

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aldo94vp/meraki-clients/structs"
)

// BadResponse structure
type BadResponse struct {
	Errors []string `json:"errors"`
}

// ReturnIndexSSID check if SSID of Meraki exists in list of Networks of Wauth (indexOf)
func ReturnIndexSSID(name string, SSIDS structs.SSIDs) int {
	for index, ssid := range SSIDS {
		if ssid.Name == name {
			return index
		}
	}
	return -1
}

// GetNetworks from Meraki Organization
func GetNetworks(config structs.Config) structs.Networks {
	var networks = structs.Networks{}
	endpoint := "/organizations/" + config.OrganizationID + "/networks"

	resp, err := GetRequest(config, endpoint)

	if err != nil {
		fmt.Println("Can't get networks:", err)
		return nil
	}

	json.Unmarshal(resp, &networks)

	return networks
}

// GetESSIDs from Meraki Network
func GetESSIDs(config structs.Config, networkID string) structs.SSIDs {
	var ssids = structs.SSIDs{}
	endpoint := "/networks/" + networkID + "/ssids"

	resp, err := GetRequest(config, endpoint)

	if err != nil {
		fmt.Println("Can't get SSIDs:", err)
		return nil
	}

	json.Unmarshal(resp, &ssids)

	return ssids
}

// GetDevices from Meraki Network
func GetDevices(config structs.Config, networkID string) structs.Devices {
	var devices = structs.Devices{}
	endpoint := "/networks/" + networkID + "/devices"

	resp, err := GetRequest(config, endpoint)

	if err != nil {
		fmt.Println("Can't get networks:", err)
		return nil
	}

	json.Unmarshal(resp, &devices)

	return devices
}

// GetClients from Meraki Device
func GetClients(config structs.Config, serialDevice string) structs.Clients {
	var clients = structs.Clients{}
	endpoint := "/devices/" + serialDevice + "/clients?timespan=60"

	resp, err := GetRequest(config, endpoint)

	if err != nil {
		fmt.Println("Can't get networks:", err)
		return nil
	}

	json.Unmarshal(resp, &clients)

	return clients
}

// GetClient from Meraki Device
func GetClient(config structs.Config, networkID string, clientID string) structs.Client {
	var client = structs.Client{}
	endpoint := "/networks/" + networkID + "/clients/" + clientID

	resp, err := GetRequest(config, endpoint)

	if err != nil {
		fmt.Println("Can't get networks:", err)
		return structs.Client{}
	}

	json.Unmarshal(resp, &client)

	return client
}

// GetRequest from Meraki Device
func GetRequest(config structs.Config, endpoint string) ([]byte, error) {
	url := config.URL + endpoint
	println("Doing http request to: ", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Cisco-Meraki-API-Key", config.Key)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Can't do request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Can't parse body")
	}
	if resp.StatusCode != 200 {
		var error = BadResponse{}
		json.Unmarshal(body, &error)
		fmt.Println("You have an error because:", error)
		return nil, errors.New("Status code: " + string(resp.StatusCode))
	}

	return body, err
}
