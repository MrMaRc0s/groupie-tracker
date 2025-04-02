function initMap() {
    const map = new google.maps.Map(document.getElementById('map'), {
        zoom: 2,
        center: { lat: 20, lng: 0 }, // Default center
    });

    // Add markers for all locations
    window.artistsData.forEach(artist => {
        artist.Locations.forEach(location => {
            if (location.Lat && location.Lng) {
                new google.maps.Marker({
                    position: { lat: location.Lat, lng: location.Lng },
                    map: map,
                    title: `${artist.Name} - ${location.Name}`,
                });
            console.log('Artists data:', window.artistsData);
            console.log('Total locations:', window.artistsData.reduce((acc, artist) => acc + artist.Locations.length, 0));
            }
        });
    });
}
