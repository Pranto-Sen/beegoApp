package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	beego.Controller
}

type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// func (c *VotingController) Get() {
// 	c.TplName = "index.tpl"
// }

func (c *VotingController) GetCatImages() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	apiUrl := "https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	responseChannel := make(chan []CatImage)
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
		var catImages []CatImage
		if err := json.Unmarshal(body, &catImages); err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- catImages
	}()

	select {
	case catImages := <-responseChannel:
		if len(catImages) > 0 {
			c.Data["json"] = catImages[0]
		} else {
			c.Data["json"] = map[string]interface{}{"error": "No cat image found"}
		}
	case err := <-errorChannel:
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}

	c.ServeJSON()
}