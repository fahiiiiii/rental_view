// static/js/property-details.js
function loadPropertyDetails() {
    const urlParams = new URLSearchParams(window.location.search);
    const propertyId = urlParams.get('property_id');

    $.ajax({
        url: '/v1/property/details',
        method: 'GET',
        data: { property_id: propertyId },
        success: function(response) {
            if (response.code === 200) {
                const details = response.data;
                displayPropertyDetails(details);
            }
        },
        error: function(xhr, status, error) {
            console.error('Error fetching property details:', error);
        }
    });
}

function displayPropertyDetails(details) {
    const property = details.property;
    
    $('#location').text(details.location);
    $('#propertyName').text(property.name);
    $('#propertyTitle').text(property.name);
    
    $('.review-score').text(property.review_score);
    $('.review-count').text(`${property.review_count} Reviews`);
    $('.review-word').text(property.review_score_word);
    
    $('.description').text(property.description);
    
    const amenitiesList = property.amenities.map(amenity => 
        `<div class="amenity-item">${amenity}</div>`
    ).join('');
    $('.amenities').html(amenitiesList);
}

$(document).ready(function() {
    loadPropertyDetails();
});