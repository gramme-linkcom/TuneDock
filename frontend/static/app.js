const trackList = document.getElementById("trackList");
const audioPlayer = document.getElementById("audioPlayer");
const playerTitle = document.getElementById("playerTitle");
const playerMeta = document.getElementById("playerMeta");
const playerCover = document.getElementById("playerCover");
const libraryUpdateBtn = document.getElementById("libraryUpdateBtn");

async function loadTracks() {
  const res = await fetch("/api/tracks");
  if (!res.ok) {
    trackList.innerHTML = "<p>Failed to load tracks.</p>";
    return;
  }

  const tracks = await res.json();
  renderTracks(tracks);
}

function renderTracks(tracks) {
  trackList.innerHTML = "";

  for (const track of tracks) {
    const row = document.createElement("div");
    row.className = "track-row";

    const cover = document.createElement("img");
    cover.className = "track-cover";
    cover.src = `/api/tracks/data/cover?id=${encodeURIComponent(track.id)}`;
    cover.onerror = () => {
      cover.src = "";
    };

    const info = document.createElement("div");

    const title = document.createElement("div");
    title.className = "track-title";
    title.textContent = track.title;

    const codec = document.createElement("div");
    codec.className = "track-codec";

    const meta = [];
    if (track.codec) meta.push(track.codec.toUpperCase());
    if (track.sample_rate) meta.push(`${track.sample_rate}Hz`);
    if (track.bitrate) meta.push(`${track.bitrate}kbps`);

    codec.textContent = meta.join(" / ");

    info.appendChild(title);
    info.appendChild(codec);

    const badge = document.createElement("div");
    if (track.is_hires) {
      badge.className = "badge";
      badge.textContent = "Hi-Res";
    }

    row.appendChild(cover);
    row.appendChild(info);
    row.appendChild(badge);

    row.addEventListener("click", () => {
      playTrack(track);
    });

    trackList.appendChild(row);
  }
}

function playTrack(track) {
  audioPlayer.src = `/api/stream?id=${encodeURIComponent(track.id)}`;
  audioPlayer.play();

  playerTitle.textContent = track.title;
  playerMeta.textContent = track.codec ? track.codec.toUpperCase() : "Unknown codec";
  playerCover.src = `/api/tracks/data/cover?id=${encodeURIComponent(track.id)}`;
}

libraryUpdateBtn.addEventListener("click", async () => {
  libraryUpdateBtn.disabled = true;
  libraryUpdateBtn.textContent = "Updating...";

  try {
    const res = await fetch("/library/update", {
      method: "GET",
    });

    if (!res.ok) {
      alert("Library update failed.");
      return;
    }

    await loadTracks();
  } finally {
    libraryUpdateBtn.disabled = false;
    libraryUpdateBtn.textContent = "Library Update";
  }
});

loadTracks();
