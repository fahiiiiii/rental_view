// Function to load and populate cities
async function loadCities() {
    // Get the select element
    const destinationSelect = document.getElementById('destinationSelect');
    
    // If options are already loaded (more than just the placeholder), don't reload
    if (destinationSelect.options.length > 1) {
        return;
    }

    try {
        const response = await fetch('/static/data/cities.json');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const cities = await response.json();
        
        // Sort cities alphabetically by city name
        cities.sort((a, b) => a.city_name.localeCompare(b.city_name));
        
        // Add new options
        cities.forEach(city => {
            const option = document.createElement('option');
            option.value = city.city_id;
            option.textContent = `${city.city_name}, ${city.country}`;
            destinationSelect.appendChild(option);
        });
    } catch (error) {
        console.error('Error loading cities:', error);
        const errorOption = new Option('Error loading cities', '');
        errorOption.disabled = true;
        destinationSelect.add(errorOption);
    }
}

// Function to load properties for a selected city
async function loadProperties(cityId) {
    try {
        const response = await fetch(`/v1/property/list?city_id=${encodeURIComponent(cityId)}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        
        if (result.code === 200 && result.data) {
            window.location.href = `/property/list?city_id=${encodeURIComponent(cityId)}`;
        } else {
            alert(result.message || 'Error loading properties');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Error loading properties. Please try again.');
    }
}

// Initialize when document is ready
document.addEventListener('DOMContentLoaded', function() {
    const destinationSelect = document.getElementById('destinationSelect');
    const searchForm = document.querySelector('form');
    
    // Load cities only once on initial page load
    loadCities();
    
    // Handle form submission
    searchForm.addEventListener('submit', function(e) {
        e.preventDefault();
        const selectedCityId = destinationSelect.value;
        
        if (!selectedCityId) {
            alert('Please select a destination');
            return;
        }
        
        loadProperties(selectedCityId);
    });
    
    // Only load cities on focus if they haven't been loaded yet
    destinationSelect.addEventListener('focus', () => {
        if (destinationSelect.options.length <= 1) {
            loadCities();
        }
    });
});
// // Function to load and populate cities
// async function loadCities() {
//     try {
//         const response = await fetch('/static/data/cities.json');
//         if (!response.ok) {
//             throw new Error(`HTTP error! status: ${response.status}`);
//         }
//         const cities = await response.json();
        
//         // Get the select element
//         const destinationSelect = document.getElementById('destinationSelect');
        
//         // Clear existing options except the first one
//         while (destinationSelect.options.length > 1) {
//             destinationSelect.remove(1);
//         }
        
//         // Sort cities alphabetically by city name
//         cities.sort((a, b) => a.city_name.localeCompare(b.city_name));
        
//         // Add new options
//         cities.forEach(city => {
//             const option = document.createElement('option');
//             option.value = city.city_id;
//             option.textContent = `${city.city_name}, ${city.country}`;
//             destinationSelect.appendChild(option);
//         });
//     } catch (error) {
//         console.error('Error loading cities:', error);
//         const destinationSelect = document.getElementById('destinationSelect');
//         const errorOption = new Option('Error loading cities', '');
//         errorOption.disabled = true;
//         destinationSelect.add(errorOption);
//     }
// }

// // Function to load properties for a selected city
// async function loadProperties(cityId) {
//     try {
//         const response = await fetch(`/v1/property/list?city_id=${encodeURIComponent(cityId)}`);
//         if (!response.ok) {
//             throw new Error(`HTTP error! status: ${response.status}`);
//         }
//         const result = await response.json();
        
//         if (result.code === 200 && result.data) {
//             window.location.href = `/property/list?city_id=${encodeURIComponent(cityId)}`;
//         } else {
//             alert(result.message || 'Error loading properties');
//         }
//     } catch (error) {
//         console.error('Error:', error);
//         alert('Error loading properties. Please try again.');
//     }
// }

// // Initialize when document is ready
// document.addEventListener('DOMContentLoaded', function() {
//     const destinationSelect = document.getElementById('destinationSelect');
//     const searchForm = document.querySelector('form');
    
//     // Load cities on initial page load
//     loadCities();
    
//     // Handle form submission
//     searchForm.addEventListener('submit', function(e) {
//         e.preventDefault();
//         const selectedCityId = destinationSelect.value;
        
//         if (!selectedCityId) {
//             alert('Please select a destination');
//             return;
//         }
        
//         loadProperties(selectedCityId);
//     });
    
//     // Optional: Load cities on focus/click as a fallback
//     destinationSelect.addEventListener('focus', loadCities);
//     destinationSelect.addEventListener('click', loadCities);
// });