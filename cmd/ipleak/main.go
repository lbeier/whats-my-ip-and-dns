package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	letterBytes   = "abcdef0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}

type info struct {
	IP            string  `json:"ip"`
	CountryCode   string  `json:"country_code"`
	ContinentCode string  `json:"continent_code"`
	CountryName   string  `json:"country_name"`
	Timezone      string  `json:"time_zone"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	ContinentName string  `json:"continent_name"`
}

func jsonInfo(url string, wg *sync.WaitGroup, ch chan map[string]info, prefix string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	jsonInfo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	i := info{}
	err = json.Unmarshal(jsonInfo, &i)
	if err != nil {
		panic(err)
	}

	info := map[string]info{
		prefix: i,
	}
	ch <- info
}

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan map[string]info, 2)

	go jsonInfo("https://ipv4.ipleak.net/json/", &wg, ch, "IP")

	randomString := randomString()
	go jsonInfo(fmt.Sprintf("https://%s.ipleak.net/json/", randomString), &wg, ch, "DNS")

	wg.Wait()

	response := []map[string]info{
		<-ch,
		<-ch,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Write(responseJSON)
}

func randomString() string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, 40)
	for i, cache, remain := 39, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
