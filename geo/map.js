window.initMap = function () {
    const artistArray = window.artistsData;
    if (!artistArray || artistArray.length === 0) {
        console.error("No artist data found");
        return;
    }

    const artist = artistArray[0];

    if (!artist || !artist.Locations || artist.Locations.length === 0) {
        console.error("No locations found for artist");
        return;
    }

    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 4,
        center: { lat: 48.8566, lng: 2.3522 }, // Default to Paris
    });

    for (let loc of artist.Locations) {
        if (loc.lat && loc.lng) {
            new google.maps.Marker({
                position: { lat: loc.lat, lng: loc.lng },
                map: map,
                title: loc.name
            });
        }
    }
};

