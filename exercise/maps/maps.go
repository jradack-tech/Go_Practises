//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func serverStatus(status int) string {
	switch status {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	case Maintenance:
		return "Maintenance"
	case Retired:
		return "Retired"
	default:
		return ""
	}
}

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func printStatusOfServer(servers map[string]int) {
	fmt.Println("Number of servers is", len(servers))
	for serverName, status := range servers {
		fmt.Println(serverName, "-----", serverStatus(status))
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	serverList := make(map[string]int)

	//* Set all of the server statuses to `Online` when creating the map
	for _, server := range servers {
		serverList[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printStatusOfServer(serverList)

	//  - change server status of `darkstar` to `Retired`
	serverList["darkstar"] = 3

	//  - change server status of `aiur` to `Offline`
	serverList["aiur"] = 1

	//  - call display server info function
	printStatusOfServer(serverList)

	//  - change server status of all servers to `Maintenance`
	fmt.Println("All servers are in Maintenance...")
	for _, server := range servers {
		serverList[server] = Maintenance
	}

	//  - call display server info function
	printStatusOfServer(serverList)
}
