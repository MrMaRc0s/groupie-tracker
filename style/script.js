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
            }
        });
    });
});



// DEN LEITOURGEI !!!

let pauseTimeout;
function playTrack(artistID) {
    // Select the audio element by concatenating the prefix with the artist's ID
    console.log("Playing track for artist", artistID);
    clearTimeout(pauseTimeout);
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
    pauseTimeout = setTimeout(function() {
        var audioElement = document.getElementById(`audio-${artistID}`);
        if (audioElement) {
        audioElement.pause();
        audioElement.currentTime = 0; // Reset the track to the beginning
        }
    }, 100)
  }