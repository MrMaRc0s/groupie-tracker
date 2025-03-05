document.addEventListener("DOMContentLoaded", function () {
    // Select all artist cards
    const artistCards = document.querySelectorAll(".artist-cards a");
    const overlays = document.querySelectorAll(".overlay");

    // Open overlay on click
    artistCards.forEach((card) => {
        card.addEventListener("click", function (event) {
            event.preventDefault(); // Prevent the default link behavior
            const artistId = this.getAttribute("href").substring(7); // Extract ID from href
            document.getElementById(`artist-${artistId}`).classList.add("show");
        });
    });

    // Close overlay on clicking close button or outside the content
    overlays.forEach((overlay) => {
        overlay.addEventListener("click", function (event) {
            if (
                event.target.classList.contains("close-button") || // Click close button
                event.target.classList.contains("overlay") // Click outside content
            ) {
                this.classList.remove("show");
                // Extract artistID from overlay's id (assuming id format "artist-{{.ID}}")
                const artistID = this.id.substring(7);
                stopTrack(artistID);
            }
        });
    });

    // button to show filter
    const minDateRange = document.getElementById("min-date-range");
    const minDateValue = document.getElementById("min-date-value");
  
    if (minDateRange && minDateValue) {
        minDateRange.addEventListener("input", function () {
            minDateValue.textContent = this.value;
        });
  
        // Set initial value on page load
        minDateValue.textContent = minDateRange.value;
    }
});

// SEARCH BAR
const searchBar = document.getElementById('searchBar');
const suggestionsDiv = document.getElementById('suggestions');
const resultsDiv = document.getElementById('results');

let debounceTimer;
searchBar.addEventListener('keyup', function() {
  clearTimeout(debounceTimer);
  debounceTimer = setTimeout(() => {
    const query = searchBar.value;

    // Fetch suggestions.
    fetch(`/suggestions?query=${encodeURIComponent(query)}`)
      .then(res => res.json())
      .then(suggestions => {
        suggestionsDiv.innerHTML = '';
        suggestions.forEach(item => {
          const el = document.createElement('div');
          el.textContent = `${item.value} - ${item.type}`;
          suggestionsDiv.appendChild(el);
        });
      });

    // Fetch filtered artist results.
    fetch(`/search?query=${encodeURIComponent(query)}`)
      .then(res => res.json())
      .then(artists => {
        resultsDiv.innerHTML = '';
        artists.forEach(artist => {
          const el = document.createElement('div');
          el.textContent = artist.name;
          resultsDiv.appendChild(el);
        });
      });
  }, 300); // 300ms debounce delay.
});

// MUSIC
function playTrack(artistID) {
  // Select the audio element by concatenating the prefix with the artist's ID
  console.log("Playing track for artist", artistID);
  var audioElement = document.getElementById(`audio-${artistID}`);
  if (audioElement) {
    // Optionally, you can set volume or other properties here
    audioElement.play().catch(function(error) {
      console.error("Playback failed:", error);
    });
  } else {
    console.error("Audio element not found for artist", artistID);
  }
}

function stopTrack(artistID) {
  var overlay = document.getElementById(`artist-${artistID}`);
  if (overlay && overlay.classList.contains("show")) {
    return;
  }
  var audioElement = document.getElementById(`audio-${artistID}`);
  if (audioElement) {
  audioElement.pause();
  audioElement.currentTime = 0; // Reset the track to the beginning
  }
}

function showFilter() {
  const filterDiv = document.getElementById('filter');
  if (filterDiv.style.display === 'none' || filterDiv.style.display === '') {
      filterDiv.style.display = 'block';
  } else {
      filterDiv.style.display = 'none';
  }
}