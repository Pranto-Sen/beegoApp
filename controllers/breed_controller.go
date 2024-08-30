package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedController struct {
	beego.Controller
}

// func (c *BreedController) Get() {
// 	c.TplName = "index.tpl"
// }

func (c *BreedController) GetBreeds() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	apiUrl := "https://api.thecatapi.com/v1/breeds"

	responseChannel := make(chan []map[string]interface{})
	errorChannel := make(chan error)

	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", apiUrl, nil)
		req.Header.Set("x-api-key", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			errorChannel <- err
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var breeds []map[string]interface{}
		if err := json.Unmarshal(body, &breeds); err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- breeds
	}()

	select {
	case breeds := <-responseChannel:
		c.Data["json"] = breeds
	case err := <-errorChannel:
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}

	c.ServeJSON()
}
