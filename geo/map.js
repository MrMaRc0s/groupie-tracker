window.initMap = function () {
    const artist = window.artistsData; 
    if (!artist || !artist.Locations) {
        console.error("No artist or locations found");
        return;
    }

    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 4,
        center: { lat: 48.8566, lng: 2.3522 }, // Default to Paris
    });

    // Loop through locations and add markers
    for (let location of artist.Locations) {
        const { latitude, longitude, name } = location;
        if (latitude && longitude) {
            new google.maps.Marker({
                position: { lat: latitude, lng: longitude },
                map: map,
                title: name
            });
        } else {
            console.warn("Missing coordinates for location:", location);
        }
    }
}
