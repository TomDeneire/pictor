// Load constants

var page = document.getElementById("page").innerHTML;
document.getElementById("page").innerHTML = "Loading...";

var resid = await fetch("identifiers.json");
const identifiers = await resid.json();
var resmeta = await fetch("metadata.json");
const metadata = await resmeta.json();
var resindex = await fetch("index.json");
const index = await resindex.json();

// Functions

function getSelectedText(elementId) {
  var elt = document.getElementById(elementId);
  if (elt.selectedIndex == -1) return null;
  return elt.options[elt.selectedIndex].text;
}

function extractRandomSuggestionsfromIndex() {
  var suggestions = "";
  var numberOfKeywords = Object.keys(index).length;
  for (let i = 0; i < 20; i++) {
    let random = Math.floor(Math.random() * numberOfKeywords);
    let keyword = Object.keys(index)[random];
    let a = `<a onclick="document.getElementById('search').value = '${keyword}'; submit()">${keyword}</a> `;
    suggestions = suggestions + a;
  }
  return suggestions;
}

function getCollections() {
  let collections = new Set();
  Object.values(identifiers).forEach((url) => {
    if (url.startsWith("http")) {
      let components = url.split("/");
      collections.add(`${components[0]}//${components[1]}${components[2]}`);
    }
  });
  let cols_sorted = Array.from(collections);
  return cols_sorted.sort();
}

function onlyUnique(value, index, self) {
  return self.indexOf(value) === index;
}

function makeThumbnail(image) {
  const full_image = image;
  if (image.endsWith("/full/0/default.jpg")) {
    image = image.replaceAll("/full/0/default.jpg", "/150,/0/default.jpg");
  }
  return `<a target="_blank" href="${full_image}">
          <img src="${image}" alt="thumbnail" width="150"></a>`;
}

function makeManifestLink(manifest) {
  return `<a target="_blank" href="${manifest}">${manifest}</a>`;
}

function makeViewerLink(manifest) {
  const viewer = document.getElementById("viewer").value;
  const viewerName = getSelectedText("viewer");
  return `<p><a target="_blank" href="${viewer + manifest}">
    View with ${viewerName}</a><br>`;
}

// Build start page

document.getElementById("page").innerHTML = page;
document.getElementById("total").innerHTML = Object.keys(identifiers).length;
document.getElementById("suggestions").innerHTML =
  extractRandomSuggestionsfromIndex(index);

// Enable submit on enter

const input = document.getElementById("search");
input.addEventListener("keypress", function (event) {
  // If the user presses the "Enter" key on the keyboard
  if (event.key === "Enter") {
    event.preventDefault();
    // Trigger the button element with a click
    document.getElementById("submit").click();
  }
});

// Build collections selectbox

const collections = getCollections();
var select_box = document.getElementById("collections");
for (var i = 0; i < collections.length; i++) {
  var option = document.createElement("option");
  option.innerHTML = collections[i];
  option.value = collections[i];
  select_box.appendChild(option);
}

// Set focus

document.getElementById("search").focus();

// Submit function

window.submit = function () {
  document.getElementById("result").innerHTML = "";

  // Validate search input
  let search = document.getElementById("search").value;
  search = search.toLowerCase();
  if (search == "") {
    return;
  }

  // Perform search
  const searches = search.split("+");
  let result = [];
  searches.forEach((term) => {
    term = term.trim();
    let termResult = index[term];
    if (termResult != null) {
      result = result.concat(termResult);
    }
  });

  // Apply "+"
  if (searches.length > 1) {
    searches.forEach((term) => {
      term = term.trim();
      let termResult = index[term];
      if (termResult != null) {
        result = result.filter((hash) => termResult.includes(hash));
      } else {
        result = [];
      }
    });
    result = result.filter(onlyUnique);
  }

  // Filter on collections
  const collection = document.getElementById("collections").value;
  if (collection != "all") {
    result = result.filter((hash) => identifiers[hash].startsWith(collection));
  }

  // Fill result template
  let html = "<div>";
  if (result.length > 0) {
    const size = Object.keys(result).length;
    html += `${size} results</div>`;
    html += "<div>";

    result.forEach((hash, index) => {
      const manifest = identifiers[hash];
      const thumbnail = makeThumbnail(metadata[hash]["T"]);
      const title = metadata[hash]["L"];
      const manifestLink = makeManifestLink(manifest);
      const viewerLink = makeViewerLink(manifest);
      const card = `<div class="card">
                      <div class="card-body">
                        <h5 class="card-title">${title}</h5>
                        <p class="card-title">${manifestLink}</p>
                        <p class="card-text">${thumbnail}</p>
                        <p class="card-text">${viewerLink}</p>
                      </div>
                    </div>`;
      if (index % 2 === 0) {
        html += `<div class="row"><div class="col-sm-6 mb-3 mb-sm-0">${card}`;
      } else {
        html += `<div class="col-sm-6">${card}</div></div>`;
      }
      html += "</div>";
    });
  } else {
    html += `0 results`;
  }

  document.getElementById("result").innerHTML = html;
};
