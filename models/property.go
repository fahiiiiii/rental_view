// models/property.go should look like this:
package models
import(
    "fmt"
)
type Property struct {
    Id              int      `json:"id"`
    CityId          int      `json:"city_id"`
    PropertyId      string   `json:"property_id"`
    Name            string   `json:"name"`
    PropertyType    string   `json:"property_type"`
    Bedrooms        int      `json:"bedrooms"`
    Bathrooms       int      `json:"bathrooms"`
    Description     string   `json:"description"`
    ReviewScore     float64  `json:"review_score"`
    ReviewCount     int      `json:"review_count"`
    ReviewScoreWord string   `json:"review_score_word"`
    ImageType       string   `json:"image_type"`
    Amenities       []string `json:"amenities"`
}

type PropertyDetails struct {
    Property Property `json:"property"`
    Location string   `json:"location"`
}

type PropertyRepository struct {}

func NewPropertyRepository() *PropertyRepository {
    return &PropertyRepository{}
}

func (r *PropertyRepository) GetPropertiesByLocation(cityID, propertyID string) ([]*Property, error) {
    // Mock data - in a real application, this would query a database
    if cityID == "" {
        return nil, fmt.Errorf("city ID is required")
    }
    properties := []*Property{
        {
            Id:              1,
            CityId:         1,
            PropertyId:     "PROP001",
            Name:           "SBH Superior Villa At Dana Paradise",
            PropertyType:   "Villa",
            Bedrooms:       4,
            Bathrooms:      3,
            Description:    "Welcome to our luxurious villa where you can experience true paradise.",
            ReviewScore:    9.8,
            ReviewCount:    520,
            ReviewScoreWord: "Exceptional",
            ImageType:      "exterior",
            Amenities:      []string{"Air Conditioner", "Swimming Pool", "Parking"},
        },
        {
            Id:              2,
            CityId:         1,
            PropertyId:     "PROP002",
            Name:           "Staycae Holiday Homes",
            PropertyType:   "Apartment",
            Bedrooms:       2,
            Bathrooms:      2,
            Description:    "Modern apartment in the heart of the city.",
            ReviewScore:    8.9,
            ReviewCount:    320,
            ReviewScoreWord: "Excellent",
            ImageType:      "interior",
            Amenities:      []string{"Gym", "Security", "Balcony"},
        },
    }

    var filteredProperties []*Property
    for _, prop := range properties {
        matches := true
        
        if cityID != "" && prop.CityId != 1 { // Convert cityID to int in real implementation
            matches = false
        }
        
        if propertyID != "" && prop.PropertyId != propertyID {
            matches = false
        }
        
        if matches {
            filteredProperties = append(filteredProperties, prop)
        }
    }
    
    return filteredProperties, nil
}

func (r *PropertyRepository) GetPropertyDetails(propertyID string) (*PropertyDetails, error) {
    // In a real application, this would query a database
    properties := map[string]*Property{
        "PROP001": {
            Id:              1,
            CityId:         1,
            PropertyId:     "PROP001",
            Name:           "SBH Superior Villa At Dana Paradise",
            PropertyType:   "Villa",
            Bedrooms:       4,
            Bathrooms:      3,
            Description:    "Welcome to our luxurious villa where you can experience true paradise.",
            ReviewScore:    9.8,
            ReviewCount:    520,
            ReviewScoreWord: "Exceptional",
            ImageType:      "exterior",
            Amenities:      []string{"Air Conditioner", "Swimming Pool", "Parking"},
        },
        "PROP002": {
            Id:              2,
            CityId:         1,
            PropertyId:     "PROP002",
            Name:           "Staycae Holiday Homes",
            PropertyType:   "Apartment",
            Bedrooms:       2,
            Bathrooms:      2,
            Description:    "Modern apartment in the heart of the city.",
            ReviewScore:    8.9,
            ReviewCount:    320,
            ReviewScoreWord: "Excellent",
            ImageType:      "interior",
            Amenities:      []string{"Gym", "Security", "Balcony"},
        },
    }

    if property, exists := properties[propertyID]; exists {
        return &PropertyDetails{
            Property:  *property,
            Location: "Dubai, United Arab Emirates",
        }, nil
    }

    return nil, nil
}
// package models

// type Property struct {
//     Id              int      `json:"id"`
//     CityId          int      `json:"city_id"`
//     PropertyId      string   `json:"property_id"`
//     Name            string   `json:"name"`
//     PropertyType    string   `json:"property_type"`
//     Bedrooms        int      `json:"bedrooms"`
//     Bathrooms       int      `json:"bathrooms"`
//     Description     string   `json:"description"`
//     ReviewScore     float64  `json:"review_score"`
//     ReviewCount     int      `json:"review_count"`
//     ReviewScoreWord string   `json:"review_score_word"`
//     ImageType       string   `json:"image_type"`
//     Amenities       []string `json:"amenities"`
// }

// type PropertyDetails struct {
//     Property Property `json:"property"`
//     Location string   `json:"location"`
// }

// type PropertyRepository struct {}

// func NewPropertyRepository() *PropertyRepository {
//     return &PropertyRepository{}
// }
// // In models/property.go, add this method to PropertyRepository
// func (r *PropertyRepository) GetPropertiesByLocation(cityID, propertyID string) ([]*Property, error) {
//     // Mock data - in a real application, this would query a database
//     properties := []*Property{
//         {
//             Id:              1,
//             CityId:         1,
//             PropertyId:     "PROP001",
//             Name:           "SBH Superior Villa At Dana Paradise",
//             PropertyType:   "Villa",
//             Bedrooms:       4,
//             Bathrooms:      3,
//             Description:    "Welcome to our luxurious villa where you can experience true paradise.",
//             ReviewScore:    9.8,
//             ReviewCount:    520,
//             ReviewScoreWord: "Exceptional",
//             ImageType:      "exterior",
//             Amenities:      []string{"Air Conditioner", "Swimming Pool", "Parking"},
//         },
//         {
//             Id:              2,
//             CityId:         1,
//             PropertyId:     "PROP002",
//             Name:           "Staycae Holiday Homes",
//             PropertyType:   "Apartment",
//             Bedrooms:       2,
//             Bathrooms:      2,
//             Description:    "Modern apartment in the heart of the city.",
//             ReviewScore:    8.9,
//             ReviewCount:    320,
//             ReviewScoreWord: "Excellent",
//             ImageType:      "interior",
//             Amenities:      []string{"Gym", "Security", "Balcony"},
//         },
//     }

//     var filteredProperties []*Property
//     for _, prop := range properties {
//         matches := true
        
//         if cityID != "" && prop.CityId != 1 { // Convert cityID to int in real implementation
//             matches = false
//         }
        
//         if propertyID != "" && prop.PropertyId != propertyID {
//             matches = false
//         }
        
//         if matches {
//             filteredProperties = append(filteredProperties, prop)
//         }
//     }
    
//     return filteredProperties, nil
// }

// // In services/property_service.go, modify the GetPropertiesByLocation method:
// func (s *PropertyService) GetPropertiesByLocation(cityIDs, propertyIDs []string) ([]*Property, error) {
//     cityID := ""
//     propertyID := ""
    
//     if len(cityIDs) > 0 {
//         cityID = cityIDs[0]
//     }
    
//     if len(propertyIDs) > 0 {
//         propertyID = propertyIDs[0]
//     }
    
//     return s.repository.GetPropertiesByLocation(cityID, propertyID)
// }
// func (r *PropertyRepository) GetPropertyDetails(propertyID string) (*PropertyDetails, error) {
//     // In a real application, this would query a database
//     properties := map[string]*Property{
//         "PROP001": {
//             Id:              1,
//             CityId:         1,
//             PropertyId:     "PROP001",
//             Name:           "SBH Superior Villa At Dana Paradise",
//             PropertyType:   "Villa",
//             Bedrooms:       4,
//             Bathrooms:      3,
//             Description:    "Welcome to our luxurious villa where you can experience true paradise.",
//             ReviewScore:    9.8,
//             ReviewCount:    520,
//             ReviewScoreWord: "Exceptional",
//             ImageType:      "exterior",
//             Amenities:      []string{"Air Conditioner", "Swimming Pool", "Parking"},
//         },
//         "PROP002": {
//             Id:              2,
//             CityId:         1,
//             PropertyId:     "PROP002",
//             Name:           "Staycae Holiday Homes",
//             PropertyType:   "Apartment",
//             Bedrooms:       2,
//             Bathrooms:      2,
//             Description:    "Modern apartment in the heart of the city.",
//             ReviewScore:    8.9,
//             ReviewCount:    320,
//             ReviewScoreWord: "Excellent",
//             ImageType:      "interior",
//             Amenities:      []string{"Gym", "Security", "Balcony"},
//         },
//     }

//     if property, exists := properties[propertyID]; exists {
//         return &PropertyDetails{
//             Property:  *property,
//             Location: "Dubai, United Arab Emirates",
//         }, nil
//     }

//     return nil, nil
// }