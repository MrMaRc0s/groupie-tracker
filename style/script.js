document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/artists")
        .then(response => response.json())
        .then(data => {
            const artistsContainer = document.getElementById("artists");
            artistsContainer.innerHTML = ""; // Clear placeholder text

            data.forEach(artist => {
                const artistCard = document.createElement("div");
                artistCard.classList.add("artist-card");
                artistCard.innerHTML = `
                    <h2>${artist.name}</h2>
                    <img src="${artist.image}" alt="${artist.name}"border-radius: 10px;">
                `;
                artistCard.addEventListener("click", () => {
                    showArtistDetails(artist);
                });
                artistsContainer.appendChild(artistCard);
            });
        })
        .catch(error => {
            console.error("Error fetching artists:", error);
            document.getElementById("artists").innerHTML = "<h2>Failed to load artists.</h2>";
        });
});

function showArtistDetails(artist) {
    const modalContent = document.getElementById("modalContent");
    modalContent.innerHTML = `
        <div class="artist-detail-card">
        <img src="${artist.image}" alt="${artist.name}" class="artist-detail-image">
        <div class="artist-detail-info">
            <h2>${artist.name}</h2>
            <p><strong>Members:</strong> ${artist.members.join(", ")}</p>
            ${artist.info ? `<p><strong>Info:</strong> ${artist.info}</p>` : ""}
            <!-- Add more fields here as needed -->
        </div>
        </div>
    `;
  
    // Display the modal and overlay
    document.getElementById("artistModal").classList.add("show");
    document.getElementById("overlay").classList.add("show");
  }

  function hideModal() {
    document.getElementById("artistModal").classList.remove("show");
    document.getElementById("overlay").classList.remove("show");
  }
  
  // Set up event listeners for closing the modal
  document.getElementById("modalClose").addEventListener("click", hideModal);
  document.getElementById("overlay").addEventListener("click", hideModal);