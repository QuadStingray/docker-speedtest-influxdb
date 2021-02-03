package speedtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SpeedTestServer struct {
	// The right side is the name of the JSON variable
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Lat         float64 `json:"latitude"`
	Lon         float64 `json:"longitude"`
	Roundrobin  bool    `json:"roundrobin"`
	Site        string  `json:"site"`
	UplinkSpeed string  `json:"uplink_speed"`
}

func ListServer() ([]SpeedTestServer, error) {
	var (
		err      error
		servers  []SpeedTestServer
		response *http.Response
		body     []byte
	)

	response, err = http.Get("https://siteinfo.mlab-oti.measurementlab.net/v1/sites/locations.json")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &servers)
	if err != nil {
		fmt.Println(err)
	}

	return servers, nil
}

func FindServerByFQDN(fqdn string) (SpeedTestServer, error) {
	serverList, _ := ListServer()
	for _, server := range serverList {
		if strings.Contains(fqdn, "."+server.Site+".") || strings.Contains(fqdn, "-"+server.Site+"-") || strings.Contains(fqdn, "-"+server.Site+".") {
			return server, nil
		}

	}
	serverList2, _ := ListServer()
	for _, server := range serverList2 {
		if strings.Contains(fqdn, "."+server.Site+".") || strings.Contains(fqdn, "-"+server.Site+"-") || strings.Contains(fqdn, "-"+server.Site+".") {
			return server, nil
		}

	}
	return SpeedTestServer{}, nil
}
