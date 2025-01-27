// services/property_service.go should look like this:
package services

import (
    "rental_view/models"
)

type PropertyService struct {
    repository *models.PropertyRepository
}

func NewPropertyService() *PropertyService {
    return &PropertyService{
        repository: models.NewPropertyRepository(),
    }
}

func (s *PropertyService) GetPropertyDetails(propertyID string) (*models.PropertyDetails, error) {
    return s.repository.GetPropertyDetails(propertyID)
}

func (s *PropertyService) GetPropertiesByLocation(cityIDs, propertyIDs []string) ([]*models.Property, error) {
    cityID := ""
    propertyID := ""
    
    if len(cityIDs) > 0 {
        cityID = cityIDs[0]
    }
    
    if len(propertyIDs) > 0 {
        propertyID = propertyIDs[0]
    }
    
    return s.repository.GetPropertiesByLocation(cityID, propertyID)
}
// package services

// import (
//     "rental_view/models"
// )

// type PropertyService struct {
//     repository *models.PropertyRepository
// }

// func NewPropertyService() *PropertyService {
//     return &PropertyService{
//         repository: models.NewPropertyRepository(),
//     }
// }

// func (s *PropertyService) GetPropertyDetails(propertyID string) (*models.PropertyDetails, error) {
//     return s.repository.GetPropertyDetails(propertyID)
// }

// // func (s *PropertyService) GetPropertiesByLocation(cityIDs, propertyIDs []string) ([]*models.Property, error) {
// //     cityID := ""
// //     propertyID := ""
    
// //     if len(cityIDs) > 0 {
// //         cityID = cityIDs[0]
// //     }
    
// //     if len(propertyIDs) > 0 {
// //         propertyID = propertyIDs[0]
// //     }
    
// //     return s.repository.GetPropertiesByLocation(cityID, propertyID)
// // }
// // package services

// // import (
// //     "rental_view/models"
// //     // "strings"
// // )

// // type PropertyService struct{}

// func (s *PropertyService) GetPropertiesByLocation(cityIDs, propertyIDs []string) ([]*models.Property, error) {
//     cityID := ""
//     propertyID := ""
    
//     if len(cityIDs) > 0 {
//         cityID = cityIDs[0]
//     }
    
//     if len(propertyIDs) > 0 {
//         propertyID = propertyIDs[0]
//     }
    
//     return models.GetPropertiesByLocation(cityID, propertyID)
// }