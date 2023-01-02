// Constants

const Universal = "https://uv-v4.netlify.app/#?c=&m=&s=&cv=&manifest="
const Mirador = "https://projectmirador.org/embed/?iiif-content="
const Clover = "https://samvera-labs.github.io/clover-iiif/?iiif-content="

// Load databases

var page = document.getElementById("page").innerHTML;
document.getElementById("page").innerHTML = "Loading...";

var resid = await fetch('https://tomdeneire.github.io/pictor/identifiers.json')
const identifiers = await resid.json()
var resmeta = await fetch('https://tomdeneire.github.io/pictor/metadata.json')
const metadata = await resmeta.json()
var resindex = await fetch('https://tomdeneire.github.io/pictor/index.json')
const index = await resindex.json()

document.getElementById("page").innerHTML = page;
document.getElementById("total").innerHTML = Object.keys(identifiers).length;

// Extract random suggestions

var suggestions = "";
var numberOfKeywords = Object.keys(index).length;
for (let i = 0; i < 20; i++) {
    let random = Math.floor(Math.random() * numberOfKeywords);
    let keyword = Object.keys(index)[random];
    let a = `<a onclick="document.getElementById('search').value = '${keyword}'; submit()">${keyword}</a> `;
    suggestions = suggestions + a;
}
document.getElementById("suggestions").innerHTML = suggestions

// Enable enter

var input = document.getElementById("search");
input.addEventListener("keypress", function(event) {
    // If the user presses the "Enter" key on the keyboard
    if (event.key === "Enter") {
        event.preventDefault();
        // Trigger the button element with a click
        document.getElementById("submit").click();
    }
});

// Collections

let collections = new Set();
Object.values(identifiers).forEach(url => {
    if (url.startsWith("http")) {
        let components = url.split("/");
        collections.add(`${components[0]}//${components[1]}${components[2]}`);
    }
})
let cols_sorted = Array.from(collections)
cols_sorted.sort()

var select_box = document.getElementById('collections');
for (var i = 0; i < cols_sorted.length; i++) {
    var option = document.createElement('option');
    option.innerHTML = cols_sorted[i];
    option.value = cols_sorted[i];
    select_box.appendChild(option);
}

// Help functions

function onlyUnique(value, index, self) {
    return self.indexOf(value) === index;
}

function onlySelected(value, index, self) {
    return self.indexOf(value) === index;
}

// Submit function

window.submit = function() {
    document.getElementById("result").innerHTML = "";
    // Validate search input
    let search = document.getElementById("search").value;
    search = search.toLowerCase();
    if (search == "") {
        return
    }
    // Perform search
    let searches = search.split("+");
    let result = [];
    searches.forEach(term => {
        term = term.trim()
        let termResult = index[term];
        if (termResult != null) {
            result = result.concat(termResult);
        };
    });
    // Apply "+"
    if (searches.length > 1) {
        searches.forEach(term => {
            term = term.trim();
            let termResult = index[term];
            if (termResult != null) {
                result = result.filter(x => termResult.includes(x));
            } else {
                result = [];
            };
        });
        result = result.filter(onlyUnique);
    };
    let collection = document.getElementById("collections").value
    if (collection != "all") {
        result = result.filter(hash => (identifiers[hash]).startsWith(collection));
    };
    // Fill result template
    let html = "<table>";
    if (result != null) {
        let size = Object.keys(result).length;
        html = `${html}<tr><td></td><td>${size} results</td></tr>`
        result.forEach(hash => {
            let url = identifiers[hash];
            let img = metadata[hash]["T"];
            let full_img = img;
            if (img.endsWith("/full/0/default.jpg")) {
                img = img.replaceAll("/full/0/default.jpg", "/150,/0/default.jpg");
            }
            let a = `<a target="_blank" href="${full_img}">
            <img src="${img}" alt="thumbnail" width="150"></a>`;
            let desc = metadata[hash]["L"] + "<p>" + `<a target="_blank" href="${url}">${url}</a>`;
            let viewers = `<p><a target="_blank" href="${Universal + url}">View with Universal</a>
            <br><a target="_blank" href="${Mirador + url}">View with Mirador</a>
            <br><a target="_blank" href="${Clover + url}">View with Clover</a>`
            html = `${html}<tr><td>${a}</td><td>${desc + viewers}</td></tr>`;
        })

    } else { html = `${html}<tr><td></td><td>0 results</td></tr>` };
    document.getElementById("result").innerHTML = html + "</table>";
}

// Set focus

window.onload = function() {
    document.getElementById("search").focus();
}
