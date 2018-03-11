package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/tutabeier/ipleak/pkg/models"
	"github.com/tutabeier/ipleak/pkg/utils"
	"io/ioutil"
	"net/http"
)

func jsonInfo(url string) models.Info {
	res, err := http.Get("https://ipv4.ipleak.net/json/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	jsonInfo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	info := models.Info{}
	err = json.Unmarshal(jsonInfo, &info)
	if err != nil {
		panic(err)
	}
	return info
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ipInfo := jsonInfo("https://ipv4.ipleak.net/json/")

	randomString := utils.RandStringBytesMaskImprSrc()
	dnsURL := fmt.Sprintf("https://%ss.ipleak.net/json/", randomString)
	dnsInfo := jsonInfo(dnsURL)

	response := models.Response{
		IP:  ipInfo,
		DNS: dnsInfo,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(responseJSON)
}
