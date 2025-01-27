package controllers

import (
	"encoding/json"
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type City struct {
	CityName string `json:"city_name"`
}

// Get handles the root path "/" to render the index page
func (c *MainController) Get() {
	c.TplName = "index.html"
}

// GetCities handles the "/destinations" path to serve cities data
func (c *MainController) GetCities() {
	// Read the cities.json file
	filePath := filepath.Join("static", "data", "cities.json")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to read cities data"}
		c.ServeJSON()
		return
	}

	// Parse the JSON data
	var cities []City
	if err := json.Unmarshal(fileContent, &cities); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to parse cities data"}
		c.ServeJSON()
		return
	}

	// Serve the cities as JSON
	c.Data["json"] = cities
	c.ServeJSON()
}
// package controllers

// import (
// 	beego "github.com/beego/beego/v2/server/web"
// )

// type MainController struct {
// 	beego.Controller
// }

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.vip"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.html"
// }
