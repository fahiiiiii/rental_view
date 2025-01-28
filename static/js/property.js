async function loadProperties() {
    const urlParams = new URLSearchParams(window.location.search);
    const cityId = urlParams.get('city_id');
    
    if (!cityId) {
        $('#propertyContainer').html(
            '<div class="text-center p-4">' +
            '<p class="text-lg text-gray-600">Please select a destination</p>' +
            '</div>'
        );
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/v1/property/list?city_id=${encodeURIComponent(cityId)}`, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            mode: 'cors',
            credentials: 'include'
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        
        // Debugging: Print fetched data
        console.log("Fetched data:", data);

        if (!Array.isArray(data)) {
            throw new Error("Invalid data format: Expected an array.");
        }

        // Transform the raw database data into the expected format
        const transformedProperties = data.map(property => ({
            property_id: property.propertyId,  // Convert propertyId -> property_id
            name: property.name,
            property_type: property.propertyType,  // Convert propertyType -> property_type
            bedrooms: property.bedrooms,
            bathrooms: property.bathrooms,
            amenities: Array.isArray(property.amenities) 
                ? property.amenities 
                : JSON.parse(property.amenities || "[]"),  // Fix amenities parsing
            image_urls: ['/static/images/default-property.jpg'], // default image
            price: property.price || null,
            review_score: property.review_score || null,
            review_count: property.review_count || 0
        }));

        console.log("Transformed properties:", transformedProperties);

        displayProperties(transformedProperties);
    } catch (error) {
        console.error('Error loading properties:', error);
        $('#propertyContainer').html(
            '<div class="text-center p-4">' +
            `<p class="text-lg text-red-600">Error loading properties: ${error.message}</p>` +
            '</div>'
        );
    }
}


// async function loadProperties() {
//     const urlParams = new URLSearchParams(window.location.search);
//     const cityId = urlParams.get('city_id');
    
//     if (!cityId) {
//         $('#propertyContainer').html(
//             '<div class="text-center p-4">' +
//             '<p class="text-lg text-gray-600">Please select a destination</p>' +
//             '</div>'
//         );
//         return;
//     }

//     try {
//         const response = await fetch(`http://localhost:8080/v1/property/list?city_id=${encodeURIComponent(cityId)}`, {
//             method: 'GET',
//             headers: {
//                 'Accept': 'application/json',
//                 'Content-Type': 'application/json'
//             },
//             mode: 'cors',
//             credentials: 'include'
//         });
        
//         if (!response.ok) {
//             throw new Error(`HTTP error! status: ${response.status}`);
//         }

//         const data = await response.json();
        
//         // Transform the raw database data into the expected format
//         const transformedProperties = data.map(property => ({
//             property_id: property.property_id,
//             name: property.name,
//             property_type: property.property_type,
//             bedrooms: property.bedrooms,
//             bathrooms: property.bathrooms,
//             amenities: Array.isArray(property.amenities) 
//                 ? property.amenities 
//                 : JSON.parse(property.amenities),
//             image_urls: ['/static/images/default-property.jpg'], // default image
//             price: null,  // add if you have price data
//             review_score: null, // add if you have review data
//             review_count: 0 // add if you have review data
//         }));

//         displayProperties(transformedProperties);
//     } catch (error) {
//         console.error('Error loading properties:', error);
//         $('#propertyContainer').html(
//             '<div class="text-center p-4">' +
//             `<p class="text-lg text-red-600">Error loading properties: ${error.message}</p>` +
//             '</div>'
//         );
//     }
// }

function displayProperties(properties) {
    const container = $('#propertyContainer');
    container.empty();

    if (properties.length === 0) {
        container.html(
            '<div class="text-center p-4">' +
            '<p class="text-lg text-gray-600">No properties found for this location</p>' +
            '</div>'
        );
        return;
    }

    properties.forEach(property => {
        // Handle amenities - backend sends it as a string
        let amenitiesList = [];
        if (typeof property.amenities === 'string') {
            amenitiesList = property.amenities.split(',').map(item => item.trim());
        }

        // Handle image URLs
        const imageUrl = Array.isArray(property.image_urls) && property.image_urls.length > 0 
            ? property.image_urls[0] 
            : '/static/images/default-property.jpg';
        
        const propertyCard = `
            <div class="property-card">
                <div class="property-image-container">
                    <img src="${imageUrl}" 
                         alt="${property.name}" 
                         class="property-image">
                    ${property.price ? `<div class="price">AED ${property.price.toLocaleString()}</div>` : ''}
                </div>
                <div class="property-info">
                    <h3>${property.name}</h3>
                    ${property.review_score ? `
                    <div class="rating">
                        <span class="rating-number">${property.review_score}</span>
                        <span>(${property.review_count} reviews)</span>
                    </div>` : ''}
                    <div class="amenities">
                        ${amenitiesList.slice(0, 4).join(' • ')}
                        ${amenitiesList.length > 4 ? ' • ...' : ''}
                    </div>
                    <div class="location">${property.property_type} • ${property.bedrooms} bedrooms</div>
                    <button class="view-button" onclick="viewProperty('${property.property_id}')">View Availability</button>
                </div>
            </div>
        `;
        container.append(propertyCard);
    });
}


function displayProperties(properties) {
    const container = $('#propertyContainer');
    container.empty();

    if (properties.length === 0) {
        container.html(
            '<div class="text-center p-4">' +
            '<p class="text-lg text-gray-600">No properties found for this location</p>' +
            '</div>'
        );
        return;
    }

    properties.forEach(property => {
        // Check if amenities is a string and convert to array if needed
        let amenitiesList = [];
        if (typeof property.amenities === 'string') {
            amenitiesList = property.amenities.split(',').map(item => item.trim());
        } else if (Array.isArray(property.amenities)) {
            amenitiesList = property.amenities;
        }
        
        const propertyCard = `
            <div class="property-card">
                <div class="property-image-container">
                    <img src="${property.image_urls?.[0] || '/static/images/default-property.jpg'}" 
                         alt="${property.name}" 
                         class="property-image">
                    ${property.price ? `<div class="price">AED ${property.price.toLocaleString()}</div>` : ''}
                </div>
                <div class="property-info">
                    <h3>${property.name}</h3>
                    ${property.review_score ? `
                    <div class="rating">
                        <span class="rating-number">${property.review_score}</span>
                        <span>(${property.review_count} reviews)</span>
                    </div>` : ''}
                    <div class="amenities">
                        ${amenitiesList.slice(0, 4).join(' • ')}
                        ${amenitiesList.length > 4 ? ' • ...' : ''}
                    </div>
                    <div class="location">${property.property_type} • ${property.bedrooms} bedrooms</div>
                    <button class="view-button" onclick="viewProperty('${property.property_id}')">View Availability</button>
                </div>
            </div>
        `;
        container.append(propertyCard);
    });
}

// Add this function to handle property viewing
function viewProperty(propertyId) {
    // You can implement the property viewing logic here
    console.log('Viewing property:', propertyId);
    // For example:
    // window.location.href = `/property/detail?id=${propertyId}`;
}


// Initialize on document ready
$(document).ready(function() {
    loadProperties();
    
    // Handle search button click
    $('.search-button').click(function(e) {
        e.preventDefault();
        const searchInput = $('.search-input').val();
        if (searchInput) {
            // Refresh the page with the new search parameters
            // You'll need to implement the logic to convert the search input to a city_id
            console.log('Searching for:', searchInput);
        }
    });
});

// Initialize when document is ready
$(document).ready(function() {
    loadProperties();
    
    // Handle search form submission
    $('.search-button').click(function(e) {
        e.preventDefault();
        const searchInput = $('.search-input').val();
        
        // Load cities data and find the matching city
        fetch('/static/data/cities.json')
            .then(response => response.json())
            .then(cities => {
                const city = cities.find(c => 
                    `${c.city_name}, ${c.country}`.toLowerCase() === searchInput.toLowerCase()
                );
                
                if (city) {
                    window.location.href = `/property/list?city_id=${encodeURIComponent(city.city_id)}`;
                } else {
                    alert('City not found. Please select a valid destination.');
                }
            })
            .catch(error => {
                console.error('Error loading cities:', error);
                alert('Error loading cities. Please try again.');
            });
    });
});

