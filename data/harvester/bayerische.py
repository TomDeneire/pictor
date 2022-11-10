import requests
import datetime

INITIAL = "https://api.digitale-sammlungen.de/iiif/presentation/v2/collection/top?cursor=initial"

url = INITIAL

while url:
    try:
        manifests = requests.get(url, timeout=10).json()
    except Exception as e:
        print(url, e, datetime.datetime.now())
        break
    for manifest in manifests["manifests"]:
        print(manifest["@id"])
    next_page = manifests.get("next")
    if next_page:
        url = next_page
    else:
        url = None
