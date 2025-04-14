let currentAudioElement = null; // Keep track of the currently playing audio element
let globalVolume = 1.0;

// VOLUME SLIDER
function showSlider() {
  const slider = document.getElementById("volume-slider");
  const output = document.getElementById("volume-value");

  // Toggle display using a helper function
  toggleDisplay(slider);
  toggleDisplay(output);
}

function toggleDisplay(element) {
  if (element.style.display === "none" || element.style.display === "") {
    element.style.display = "block";
  } else {
    element.style.display = "none";
  }
}

// Update audio volume when slider value changes
document.addEventListener('DOMContentLoaded', function () {
  const volumeSlider = document.getElementById("volume-slider");
  const volumeIcon = document.getElementById("sound-icon");
  const output = document.getElementById("volume-value");

  volumeSlider.addEventListener('input', function() {
    const volume = this.value;
    let iconSrc;

    if (volume === '0') {
      iconSrc = 'style/media/icons/volume-min.png';
    } else if (volume <= 45) {
      iconSrc = 'style/media/icons/volume-1.png';
    } else if (volume <= 85) {
      iconSrc = 'style/media/icons/volume-2.png';
    } else {
      iconSrc = 'style/media/icons/volume-max.png';
    }
    output.textContent = volume; // Update the displayed volume value
    globalVolume = volume / 100; // Update the global volume variable

    const slideRect = volumeSlider.getBoundingClientRect();
    const sliderWidth = slideRect.width;
    const thumbWidth = volumeSlider.offsetHeight;
    const position = ((volume - this.min) / (this.max - this.min));
    const offset = position * (sliderWidth - thumbWidth) + (thumbWidth / 2);

    volumeIcon.src = iconSrc;
    output.style.left = `${offset}px`;
    output.textContent = `${volume}%`;
  });
});

// MUSIC
function playTrack(artistID) {
  console.log("Playing track for artist", artistID);
  const audioElement = document.getElementById(`audio-${artistID}`);
  if (audioElement) {
    currentAudioElement = audioElement; // Set the current audio element
    audioElement.volume = globalVolume; // Apply the global volume
    audioElement.play().catch(function (error) {
      console.error("Playback failed:", error);
    });
  } else {
    console.error("Audio element not found for artist", artistID);
  }
}

function stopTrack(artistID) {
  const overlay = document.getElementById(`artist-${artistID}`);
  if (overlay && overlay.classList.contains("show")) {
    return;
  }
  const audioElement = document.getElementById(`audio-${artistID}`);
  if (audioElement) {
    audioElement.pause();
    audioElement.currentTime = 0; // Reset the track to the beginning
    if (currentAudioElement === audioElement) {
      currentAudioElement = null; // Clear the current audio element
    }
  }
}

function togglePlayPause(artistId) {
  const audioElement = document.getElementById(`audio-${artistId}`);
  const pauseIcon = document.getElementById(`pause-icon-${artistId}`);
  const playIcon = document.getElementById(`play-icon-${artistId}`);

  if (audioElement.paused) {
    audioElement.play();
    pauseIcon.style.display = 'inline';
    playIcon.style.display = 'none';
  } else {
    audioElement.pause();
    pauseIcon.style.display = 'none';
    playIcon.style.display = 'inline';
  }
}

function toggleHeader() {
  const header = document.querySelector('.main-header');
  const btn = document.querySelector('.close-header');
  header.classList.toggle('hidden');

  if (header.classList.contains('hidden')) {
    btn.innerHTML = "&#x25BC;";
  } else {
    btn.innerHTML = "&#x25B2;";
  }
}

// Re-show the header when scrolled to the top.
window.addEventListener("scroll", function() {
  const closeBtn = document.querySelector('.close-header');

  if (window.pageYOffset > 100) {
    closeBtn.style.display = "block";
  } else if (window.pageYOffset === 0) {
    closeBtn.innerHTML = "&#x25B2;";
    closeBtn.style.display = "none";
    document.querySelector('.main-header').classList.remove('hidden');
  }
});

document.addEventListener("DOMContentLoaded", function () {
    // Select all artist cards
    const artistCards = document.getElementById("artists");
    const overlays = document.querySelectorAll(".overlay");
  
    overlays.forEach(overlay => {
      const heading = overlay.querySelector("h2");
      const btn = overlay.querySelector('.pause-btn');

      if (heading.textContent.trim().length > 21) {
        btn.classList.add("big-name");
      }
    });

    // Open overlay on click
    if (artistCards) {
      artistCards.addEventListener("click", function (event) {
        const card = event.target.closest(".card");
        if (card) {
          const artistId = card.getAttribute("data-artist-id");
          const overlay = document.getElementById(`artist-${artistId}`);
          if (overlay) {
            overlay.classList.add("show");
            const pauseIcon = document.getElementById(`pause-icon-${artistId}`);
            const playIcon = document.getElementById(`play-icon-${artistId}`);
            if (pauseIcon && playIcon) {
              pauseIcon.style.display = "inline";
              playIcon.style.display = "none";
            }
          }
        }
      });
    }

    // Close overlay on clicking close button or outside the content
    document.addEventListener("click", function (event) {
      if (
        event.target.classList.contains("close-button") || // Click close button
        event.target.classList.contains("overlay") // Click outside content
      ) {
        const overlay = event.target.closest(".overlay");
        if (overlay) {
          overlay.classList.remove("show");
          // Extract artistID from overlay's id (assuming id format "artist-{{.ID}}")
          const artistID = overlay.id.substring(7);
          stopTrack(artistID);
        }
      }
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
    
    const minAlbumDateRange = document.getElementById("min-album-date-range");
    const minAlbumDateValue = document.getElementById("min-album-date-value");
  
    if (minAlbumDateRange && minAlbumDateValue) {
        minAlbumDateRange.addEventListener("input", function () {
            minAlbumDateValue.textContent = this.value;
        });
  
        // Set initial value on page load
        minAlbumDateValue.textContent = minAlbumDateRange.value;
    }
});

// SEARCH BAR
  // A simple debounce function to avoid making a request on every single keystroke immediately.
  function debounce(func, delay) {
    let timeout;
    return function (...args) {
      clearTimeout(timeout);
      timeout = setTimeout(() => func.apply(this, args), delay);
    };
  }

  document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("search-bar");
    const resultsContainer = document.getElementById("artists"); // Your div container for artists

    // Event listener with debouncing (adjust the delay as needed)
    searchInput.addEventListener("input", debounce(async function () {
      const query = searchInput.value;
      
      try {
        // Update the URL or API endpoint as necessary.
        const response = await fetch(`/homepage?query=${encodeURIComponent(query)}`, {
          headers: {
            'X-Requested-With': 'XMLHttpRequest'
          }
        });
        
        if (response.ok) {
          // You can either expect a JSON response or an HTML fragment.
          // For example, if you're returning HTML:
          const htmlFragment = await response.text();
          resultsContainer.innerHTML = htmlFragment;
        }
      } catch (error) {
        console.error("Error fetching search results: ", error);
      }
    }, 300));  // Delay in ms; here 300ms delay after last keystroke
  });

// Filter
function showFilter() {
  const filter = document.getElementById('filter');
  filter.classList.toggle('show');
}