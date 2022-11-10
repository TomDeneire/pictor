import requests

INITIAL = "https://api.digitale-sammlungen.de/iiif/presentation/v2/collection/top?cursor=initial"

url = INITIAL

while url:
    try:
        manifests = requests.get(url).json()
    except Exception as e:
        print(url, e)
        break
    for manifest in manifests["manifests"]:
        print(manifest["@id"])
    next_page = manifests.get("next")
    if next_page:
        url = next_page
    else:
        url = None
