package controllers

import (
    beego "github.com/beego/beego/v2/server/web"
    "rental_view/services"
)

type PropertyDetailsController struct {
    beego.Controller
}

func (c *PropertyDetailsController) Get() {
    propertyID := c.GetString("property_id")
    propertyService := services.NewPropertyService()
    
    details, err := propertyService.GetPropertyDetails(propertyID)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "code":    500,
            "message": "Error: " + err.Error(),
            "data":    nil,
        }
    } else if details == nil {
        c.Data["json"] = map[string]interface{}{
            "code":    404,
            "message": "Property not found",
            "data":    nil,
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "code":    200,
            "message": "Success",
            "data":    details,
        }
    }
    c.ServeJSON()
}
// func (c *PropertyDetailsController) Get() {
//     propertyId := c.GetString("property_id")
    
//     propertyDetails := &models.PropertyDetails{
//         Property: models.Property{
//             Id:              1,
//             PropertyId:      propertyId,
//             Name:            "SBH Superior Villa At Dana Paradise",
//             PropertyType:    "Villa",
//             Bedrooms:        4,
//             Bathrooms:       3,
//             Description:     "Welcome to our luxurious villa where you can experience true paradise.",
//             ReviewScore:     9.8,
//             ReviewCount:     520,
//             ReviewScoreWord: "Exceptional",
//             ImageType:       "exterior",
//             Amenities:       []string{"Air Conditioner", "Swimming Pool", "Parking"},
//         },
//         Location: "Dubai, United Arab Emirates",
//     }
    
//     c.Data["json"] = propertyDetails
//     c.ServeJSON()
// }