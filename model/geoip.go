package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type GeoIP struct {
	// The right side is the name of the JSON variable
	Ip                 string  `json:"ip"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	RegionCode         string  `json:"region_code"`
	Region             string  `json:"region"`
	City               string  `json:"city"`
	Postal             string  `json:"postal"`
	Lat                float64 `json:"latitude"`
	Lon                float64 `json:"longitude"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	Languages          string  `json:"languages"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

func CheckIpLocation(ip string) (GeoIP, error) {

	var (
		err      error
		geo      GeoIP
		response *http.Response
		body     []byte
	)

	response, err = http.Get("https://ipapi.co/" + ip + "/json/")
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			err = json.Unmarshal(body, &geo)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return geo, err
}

func LocateUser() (GeoIP, error) {

	var (
		err      error
		geo      GeoIP
		response *http.Response
		body     []byte
	)

	response, err = http.Get("https://ipapi.co/json/")
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			err = json.Unmarshal(body, &geo)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// Everything accessible in struct now
	fmt.Println("\n==== IP Geolocation Info ====\n")
	fmt.Println("IP address:\t", geo.Ip)
	fmt.Println("Country Code:\t", geo.CountryName)
	fmt.Println("Country Name:\t", geo.CountryName)
	fmt.Println("Zip Code:\t", geo.Postal)
	fmt.Println("Latitude:\t", geo.Lat)
	fmt.Println("Longitude:\t", geo.Lon)
	fmt.Println("Metro Code:\t", geo.City)

	return geo, err
}

//func DistanceBetweenUserAndIp(ip string, unit string) (float64, error) {
//	geoClient, _ := LocateUser()
//	time.Sleep(time.Duration(1) * time.Second)
//	geoSever, _ := CheckIpLocation(ip)
//	return Distance(geoSever.Lat, geoClient.Lat, geoSever.Lon, geoClient.Lon, unit), nil
//}
//
//func DistanceBetweenIps(ip1 string, ip2 string, unit string) (float64, error) {
//	geoClient, _ := CheckIpLocation(ip1)
//	time.Sleep(time.Duration(1) * time.Second)
//	geoSever, _ := CheckIpLocation(ip2)
//	return Distance(geoSever.Lat, geoClient.Lat, geoSever.Lon, geoClient.Lon, unit), nil
//}
//
//func DistanceBetweenIpAndGeoIp(ip string, geoIp GeoIP, unit string) (float64, error) {
//	time.Sleep(time.Duration(1) * time.Second)
//	geoSever, _ := CheckIpLocation(ip)
//	return Distance(geoSever.Lat, geoIp.Lat, geoSever.Lon, geoIp.Lon, unit), nil
//}

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if unit == "K" {
		dist = dist * 1.609344
	} else if unit == "N" {
		dist = dist * 0.8684
	}

	return dist
}
