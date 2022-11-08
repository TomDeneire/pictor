import requests

INITIAL = "https://sharedcanvas.be/IIIF/mmmonk/discover?limit=50&q=&start=0"

url = INITIAL

while url:
    try:
        manifests = requests.get(url).json()
    except Exception as e:
        print(url)
        break
    if "manifests" in manifests:
        for manifest in manifests["manifests"]:
            print(manifest["@id"])
    next_page = manifests.get("next")
    if next_page:
        url = next_page
    else:
        url = None
