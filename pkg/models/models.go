package models

type Info struct {
	IP            string  `json:"ip"`
	CountryCode   string  `json:"country_code"`
	ContinentCode string  `json:"continent_code"`
	CountryName   string  `json:"country_name"`
	Timezone      string  `json:"time_zone"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	ContinentName string  `json:"continent_name"`
}

type Response struct {
	IP  Info `json:"ip"`
	DNS Info `json:"dns"`
}
