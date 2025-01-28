// controllers/property.go
package controllers

import (
    beego "github.com/beego/beego/v2/server/web"
    // "rental_view/services"
    "net/url"
)

type PropertyController struct {
    beego.Controller
}

func (c *PropertyController) Prepare() {
    // Add CORS headers
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
}

func (c *PropertyController) List() {
    cityID := c.GetString("city_id")
    if cityID == "" {
        c.Data["json"] = map[string]interface{}{
            "code":    400,
            "message": "City ID is required",
            "data":    nil,
        }
        c.ServeJSON()
        return
    }

    // URL decode the city_id
    decodedCityID, err := url.QueryUnescape(cityID)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "code":    400,
            "message": "Invalid city ID format",
            "data":    nil,
        }
        c.ServeJSON()
        return
    }

    // Mock data for testing
    properties := []map[string]interface{}{
        {
            "id":           1,
            "city_id":      decodedCityID,
            "property_id":  "PROP001",
            "name":         "SBH Superior Villa At Dana Paradise",
            "property_type": "Villa",
            "bedrooms":     4,
            "bathrooms":    3,
            "review_score": 9.8,
            "review_count": 520,
            "price":        58781,
            "image_urls":   []string{"/static/images/villa1.jpg"},
            "amenities":    "Air Conditioner, Swimming Pool, Parking",
        },
        {
            "id":           2,
            "city_id":      decodedCityID,
            "property_id":  "PROP002",
            "name":         "Staycae Holiday Homes",
            "property_type": "Apartment",
            "bedrooms":     2,
            "bathrooms":    2,
            "review_score": 8.9,
            "review_count": 320,
            "price":        42000,
            "image_urls":   []string{"/static/images/apartment1.jpg"},
            "amenities":    "Gym, Security, Balcony",
        },
    }

    c.Data["json"] = map[string]interface{}{
        "code":    200,
        "message": "Success",
        "data":    properties,
    }
    c.ServeJSON()
}
// package controllers

// import (
//     beego "github.com/beego/beego/v2/server/web"
//     "rental_view/services"
//     "net/url" 
// )

// type PropertyController struct {
//     beego.Controller
// }

// type PropertyResponse struct {
//     Code    int         `json:"code"`
//     Message string      `json:"message"`
//     Data    interface{} `json:"data"`
// }

// func (c *PropertyController) List() {
//     cityID := c.GetString("city_id")
//     if cityID == "" {
//         c.Data["json"] = PropertyResponse{
//             Code:    400,
//             Message: "City ID is required",
//             Data:    nil,
//         }
//         c.ServeJSON()
//         return
//     }

//     // URL decode the city_id
//     decodedCityID, err := url.QueryUnescape(cityID)
//     if err != nil {
//         c.Data["json"] = PropertyResponse{
//             Code:    400,
//             Message: "Invalid city ID format",
//             Data:    nil,
//         }
//         c.ServeJSON()
//         return
//     }

//     propertyService := services.NewPropertyService()
//     properties, err := propertyService.GetPropertiesByLocation([]string{decodedCityID}, nil)
//     if err != nil {
//         c.Data["json"] = PropertyResponse{
//             Code:    500,
//             Message: "Error: " + err.Error(),
//             Data:    nil,
//         }
//     } else {
//         c.Data["json"] = PropertyResponse{
//             Code:    200,
//             Message: "Success",
//             Data:    properties,
//         }
//     }
//     c.ServeJSON()
// }
func (c *PropertyController) Index() {
    c.TplName = "property/list.html"
}
// package controllers

// import (
//      "github.com/beego/beego/v2/server/web"
//     "rental_view/services"
// )

// type PropertyController struct {
//     web.Controller
// }

// func (c *PropertyController) List() {
//     cityIDs := []string{c.GetString("city_id")}
//     propertyService := &services.PropertyService{}
    
//     properties, err := propertyService.GetPropertiesByLocation(cityIDs, nil)
//     if err != nil {
//         c.Data["json"] = PropertyResponse{
//             Code:    500,
//             Message: "Error: " + err.Error(),
//             Data:    nil,
//         }
//     } else {
//         c.Data["json"] = PropertyResponse{
//             Code:    200,
//             Message: "Success",
//             Data:    properties,
//         }
//     }
//     c.ServeJSON()
// }
// // package controllers

// // import (
// //     "github.com/beego/beego/v2/server/web"
// //     "rental_view/models"
// // )

// // type PropertyController struct {
// //     web.Controller
// // }

// // type PropertyResponse struct {
// //     Code    int         `json:"code"`
// //     Message string      `json:"message"`
// //     Data    interface{} `json:"data"`
// // }

// // // Mock data with the new structure
// // var mockProperties = []models.Property{
// //     {
// //         Id:           1,
// //         CityId:       1,
// //         PropertyId:   "PROP001",
// //         Name:         "SBH Superior Villa At Dana Paradise",
// //         PropertyType: "Villa",
// //         Bedrooms:     4,
// //         Bathrooms:    3,
// //         Amenities:    []string{"Air Conditioner", "Swimming Pool", "Parking"},
// //     },
// //     {
// //         Id:           2,
// //         CityId:       2,
// //         PropertyId:   "PROP002",
// //         Name:         "Staycae Holiday Homes",
// //         PropertyType: "Apartment",
// //         Bedrooms:     2,
// //         Bathrooms:    2,
// //         Amenities:    []string{"Gym", "Security", "Balcony"},
// //     },
// //     {
// //         Id:           3,
// //         CityId:       2,
// //         PropertyId:   "PROP003",
// //         Name:         "Luxury Pool View Apartment",
// //         PropertyType: "Apartment",
// //         Bedrooms:     3,
// //         Bathrooms:    2,
// //         Amenities:    []string{"Pool View", "Kitchen", "Wi-Fi"},
// //     },
// // }

// // func (c *PropertyController) List() {
// //     // Get filter parameters
// //     cityId := c.GetString("city_id")
// //     propertyType := c.GetString("property_type")
    
// //     // Filter properties based on query parameters
// //     var filteredProperties []models.Property
    
// //     if cityId != "" || propertyType != "" {
// //         for _, prop := range mockProperties {
// //             // Add property if it matches all provided filters
// //             matches := true
            
// //             if cityId != "" && prop.CityId != web.AppConfig.DefaultInt(cityId, 0) {
// //                 matches = false
// //             }
            
// //             if propertyType != "" && prop.PropertyType != propertyType {
// //                 matches = false
// //             }
            
// //             if matches {
// //                 filteredProperties = append(filteredProperties, prop)
// //             }
// //         }
// //     } else {
// //         filteredProperties = mockProperties
// //     }
    
// //     c.Data["json"] = PropertyResponse{
// //         Code:    200,
// //         Message: "Success",
// //         Data:    filteredProperties,
// //     }
// //     c.ServeJSON()
// // }
// // // controllers/property.go
// func (c *PropertyController) Index() {
//     c.TplName = "property/list.html"
// }

// // // package controllers

// // // import (
// // // 	beego "github.com/beego/beego/v2/server/web"
// // // 	"rental_view/services"
// // // )

// // // type PropertyController struct {
// // // 	beego.Controller
// // // }

// // // func (c *PropertyController) Get() {
// // //     cityIDs := c.GetStrings("city_id")
// // //     propertyIDs := c.GetStrings("property_id")

// // //     propertyService := &services.PropertyService{}
// // //     properties, err := propertyService.GetPropertiesByLocation(cityIDs, propertyIDs)
// // //     if err != nil {
// // //         c.Data["error"] = err.Error()
// // //         c.TplName = "error.tpl"
// // //         return
// // //     }
// // //     c.Data["properties"] = properties
// // //     c.TplName = "index.tpl"
// // // }