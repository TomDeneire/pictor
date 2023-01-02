import requests

START = "https://data.ucd.ie/api/img/collections/"

url = START

collections = requests.get(START).json()
for collection in collections["collections"]:
    url = collection["@id"]
    manifests = requests.get(url).json()
    for manifest in manifests["manifests"]:
        print(manifest["@id"])

