package routers

import (
    "rental_view/controllers"
    beego "github.com/beego/beego/v2/server/web"
)

func init() {
    // Serve static files
    beego.SetStaticPath("/static", "static")
    
    // Main routes
    beego.Router("/", &controllers.MainController{})
    beego.Router("/destinations", &controllers.MainController{}, "get:GetCities")
    
    // Property routes
    beego.Router("/property/list", &controllers.PropertyController{}, "get:Index")
    beego.Router("/property/details", &controllers.PropertyDetailsController{}, "get:Get")
    beego.Router("/v1/property/list", &controllers.PropertyController{}, "get:List")
    beego.Router("/v1/property/details", &controllers.PropertyDetailsController{}, "get:Get")
}
// package routers

// import (
//     "rental_view/controllers"
//     beego "github.com/beego/beego/v2/server/web"
// )

// func init() {
//     // Serve static files
//     beego.SetStaticPath("/static", "static")
    
//     // Main routes
//     beego.Router("/", &controllers.MainController{})
//     beego.Router("/destinations", &controllers.MainController{}, "get:GetCities")
    
//     // Property routes 
//     beego.Router("/property/list", &controllers.PropertyController{}, "get:Index")
//     beego.Router("/v1/property/list", &controllers.PropertyController{}, "get:List")
//     beego.Router("/v1/property/details", &controllers.PropertyDetailsController{})
// }
// // package routers

// // import (
// //     "rental_view/controllers"
// //     beego "github.com/beego/beego/v2/server/web"
// // )

// // func init() {
// //     // Serve static files
// //     beego.SetStaticPath("/static", "static")
    
// //     beego.Router("/", &controllers.MainController{})
// //     beego.Router("/destinations", &controllers.MainController{}, "get:GetCities")
    
    
// //     // Routes
// //     beego.Router("/", &controllers.PropertyController{}, "get:Index")
// //     beego.Router("/v1/property/list", &controllers.PropertyController{}, "get:List")
// // 	// beego.Router("/v1/property/details", &controllers.PropertyController{}, "get:Details")
// // 	beego.Router("/v1/property/details", &controllers.PropertyDetailsController{})
// // }
