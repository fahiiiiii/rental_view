$(document).ready(function() {
    // Initial load of properties
    loadProperties();

    // Handle property type filter change
    $('#propertyType').on('change', function() {
        loadProperties();
    });
});

// In property.js
function loadProperties() {
    const urlParams = new URLSearchParams(window.location.search);
    const cityId = urlParams.get('city_id');
    const propertyType = $('#propertyType').val();

    if (!cityId) {
        $('#propertyContainer').html(
            '<div class="text-center p-4">' +
            '<p class="text-lg text-gray-600">Please select a destination</p>' +
            '</div>'
        );
        return;
    }

    $.ajax({
        url: '/v1/property/list',
        method: 'GET',
        data: {
            city_id: encodeURIComponent(cityId),
            property_type: propertyType
        },
        success: function(response) {
            if (response.code === 200) {
                if (response.data && response.data.length > 0) {
                    displayProperties(response.data);
                } else {
                    $('#propertyContainer').html(
                        '<div class="text-center p-4">' +
                        '<p class="text-lg text-gray-600">No properties found for ' + cityId + '</p>' +
                        '</div>'
                    );
                }
            } else {
                console.error('Error:', response.message);
                $('#propertyContainer').html(
                    '<div class="text-center p-4">' +
                    '<p class="text-lg text-red-600">' + response.message + '</p>' +
                    '</div>'
                );
            }
        },
        error: function(xhr, status, error) {
            console.error('Error fetching properties:', error);
            $('#propertyContainer').html(
                '<div class="text-center p-4">' +
                '<p class="text-lg text-red-600">Error loading properties. Please try again.</p>' +
                '</div>'
            );
        }
    });
}
// // static/js/property.js

// $(document).ready(function() {
//     // Initial load of properties
//     loadProperties();

//     // Handle search form submission
//     $('#searchForm').on('submit', function(e) {
//         e.preventDefault();
//         loadProperties();
//     });
// });

// function loadProperties() {
//     const cityId = $('#cityInput').val();
//     const propertyType = $('#propertyType').val();

//     $.ajax({
//         url: '/v1/property/list',
//         method: 'GET',
//         data: {
//             city_id: cityId,
//             property_type: propertyType
//         },
//         success: function(response) {
//             if (response.code === 200) {
//                 displayProperties(response.data);
//             } else {
//                 console.error('Error:', response.message);
//             }
//         },
//         error: function(xhr, status, error) {
//             console.error('Error fetching properties:', error);
//         }
//     });
// }

// function displayProperties(properties) {
//     const container = $('#propertyContainer');
//     container.empty();

//     properties.forEach(property => {
//         const propertyCard = `
//             <div class="property-card">
//                 <div class="property-header">
//                     <h3>${property.name}</h3>
//                     <p>Property ID: ${property.property_id}</p>
//                 </div>
//                 <div class="property-details">
//                     <p>Type: ${property.property_type}</p>
//                     <p>Bedrooms: ${property.bedrooms}</p>
//                     <p>Bathrooms: ${property.bathrooms}</p>
//                 </div>
//                 <div class="amenities">
//                     <h4>Amenities:</h4>
//                     <ul>
//                         ${property.amenities.map(amenity => `<li>${amenity}</li>`).join('')}
//                     </ul>
//                 </div>
//                 <button class="view-btn" onclick="viewProperty('${property.property_id}')">
//                     View Availability
//                 </button>
//             </div>
//         `;
//         container.append(propertyCard);
//     });
// }

// function viewProperty(propertyId) {
//     // Handle view property action
//     console.log('Viewing property:', propertyId);
//     window.location.href = `/property/details?property_id=${propertyId}`;
//     // Add your view property logic here
// }
// // function viewProperty(propertyId) {
//     // window.location.href = `/property/details?property_id=${propertyId}`;

// // }


// // function fetchProperties() {
// //     const cityInput = document.getElementById('cityInput');
// //     const guestsInput = document.getElementById('guestsInput');

// //     // Fetch property data from the backend API
// //     fetch(`/v1/property/list?city=${cityInput.value}&guests=${guestsInput.value}`)
// //         .then(response => response.json())
// //         .then(data => {
// //             const propertyList = document.getElementById('propertyList');
// //             propertyList.innerHTML = '';

// //             data.forEach(property => {
// //                 const card = document.createElement('div');
// //                 card.className = 'property-card';

// //                 if (property.new) {
// //                     const newBadge = document.createElement('div');
// //                     newBadge.className = 'new-badge';
// //                     newBadge.textContent = 'New';
// //                     card.appendChild(newBadge);
// //                 }

// //                 const img = document.createElement('img');
// //                 img.src = property.image;
// //                 img.alt = property.name;
// //                 card.appendChild(img);

// //                 const info = document.createElement('div');
// //                 info.className = 'property-info';

// //                 const name = document.createElement('div');
// //                 name.className = 'property-name';
// //                 name.textContent = property.name;
// //                 info.appendChild(name);

// //                 const details = document.createElement('div');
// //                 details.className = 'property-details';
// //                 details.textContent = `${property.property_type} - ${property.bedrooms} Bedroom(s), ${property.bathrooms} Bathroom(s)`;
// //                 info.appendChild(details);

// //                 const amenities = document.createElement('div');
// //                 amenities.className = 'amenities';
// //                 amenities.textContent = `Amenities: ${property.amenities.join(', ')}`;
// //                 info.appendChild(amenities);

// //                 const price = document.createElement('div');
// //                 price.className = 'price';
// //                 price.textContent = `From ₹${property.price}`;
// //                 info.appendChild(price);

// //                 const button = document.createElement('a');
// //                 button.className = 'button';
// //                 button.href = `/property/${property.id}`;
// //                 button.textContent = 'View Availability';
// //                 info.appendChild(button);

// //                 card.appendChild(info);
// //                 propertyList.appendChild(card);
// //             });
// //         })
// //         .catch(error => {
// //             console.error('Error fetching properties:', error);
// //         });
// // }