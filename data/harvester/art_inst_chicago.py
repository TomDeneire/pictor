# https://api.artic.edu/docs/#iiif-manifests

import requests

url = "https://api.artic.edu/api/v1/artworks"

while url:
    res = requests.get(url, timeout=10)
    for i in res.json().get("data"):
        print(i.get("api_link") + "/manifest.json")
    url = res.json().get("pagination").get("next_url")
