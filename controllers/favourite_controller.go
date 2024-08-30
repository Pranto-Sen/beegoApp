package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type FavouriteController struct {
	beego.Controller
}

type Favourite struct {
	ID        int      `json:"id"`
	ImageID   string   `json:"image_id"`
	SubID     string   `json:"sub_id"`
	CreatedAt string   `json:"created_at"`
	Image     CatImage `json:"image"`
}

func (c *FavouriteController) AddFavourite() {
	apiKey, err := beego.AppConfig.String("catapi_key")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to retrieve API key"}
		c.ServeJSON()
		return
	}

	var requestBody struct {
		ImageID string `json:"image_id"`
		SubID   string `json:"sub_id"`
	}
	err = json.NewDecoder(c.Ctx.Request.Body).Decode(&requestBody)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to decode request body"}
		c.ServeJSON()
		return
	}

	apiUrl := "https://api.thecatapi.com/v1/favourites"
	jsonBody, _ := json.Marshal(requestBody)

	responseChannel := make(chan *http.Response)
	errorChannel := make(chan error)

	go func() {
		req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- resp
	}()

	select {
	case resp := <-responseChannel:
		defer resp.Body.Close()
		responseBody, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			c.Data["json"] = map[string]interface{}{
				"error":         "External API returned an error",
				"status_code":   resp.StatusCode,
				"response_body": string(responseBody),
			}
		} else {
			var result map[string]interface{}
			json.Unmarshal(responseBody, &result)
			c.Data["json"] = result
		}
	case err := <-errorChannel:
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}

	c.ServeJSON()
}

func (c *FavouriteController) GetFavourites() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	apiUrl := "https://api.thecatapi.com/v1/favourites"

	subID := c.GetString("sub_id")
	limit := c.GetString("limit", "10")
	page := c.GetString("page", "0")

	fullUrl := fmt.Sprintf("%s?sub_id=%s&limit=%s&page=%s&order=DESC", apiUrl, subID, limit, page)

	responseChannel := make(chan *http.Response)
	errorChannel := make(chan error)

	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", fullUrl, nil)
		req.Header.Set("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- resp
	}()

	select {
	case resp := <-responseChannel:
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var favourites []Favourite
		json.Unmarshal(body, &favourites)
		c.Data["json"] = favourites
	case err := <-errorChannel:
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}

	c.ServeJSON()
}

func (c *FavouriteController) DeleteFavourite() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	favouriteID := c.Ctx.Input.Param(":id")
	apiUrl := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s", favouriteID)

	responseChannel := make(chan *http.Response)
	errorChannel := make(chan error)

	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", apiUrl, nil)
		req.Header.Set("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- resp
	}()

	select {
	case resp := <-responseChannel:
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		c.Data["json"] = result
	case err := <-errorChannel:
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}

	c.ServeJSON()
}
