package routine

import (
	"bytes"
	"encoding/xml"
	"entsoe/src/config"
	"entsoe/src/global"
	"entsoe/src/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func parseXML(xmlBytes []byte, env *global.Env) *map[int]float32 {
	var e PublicationMarketDocument

	err := xml.Unmarshal(xmlBytes, &e)
	if err != nil {
		env.Logger.Error(err)
		return nil
	}
	priceSeries := make(map[int]float32)
	for _, e := range e.TimeSeries.Period.Point {
		priceSeries[e.Position] = utils.ConvertPrice(e.PriceAmount)
	}
	return &priceSeries
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func Routine(interval time.Duration, env *global.Env) {
	for range time.Tick(interval * time.Second) {
		xmlBytes, err := getXML(utils.GetEntsoeURL())
		if err != nil {
			env.Logger.Fatal(err.Error())
			return
		}
		priceData := parseXML(xmlBytes, env)
		currentPrice := (*priceData)[utils.GetHour()]
		var jsonStr = []byte(fmt.Sprintf(`{"price":%v, "timeStamp":"%v"}`, currentPrice, time.Now().Unix()))
		req, err := http.NewRequest("POST", config.GetDataEndpoint(), bytes.NewBuffer(jsonStr))
		req.Header.Add("Authorization", config.GetAccessToken())
		if err != nil {
			env.Logger.Fatalf("Failed to create backend request %v", err.Error())
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			env.Logger.Fatal(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 201 {
			body, _ := ioutil.ReadAll(resp.Body)
			env.Logger.Fatalf("Request failed with status code %v and body: %v", resp.StatusCode, string(body))
			return
		}
	}
}
