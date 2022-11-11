
import requests

INITIAL = "https://graph.global/static/data/universes/iiif/stanford.json"


def parse_collection(collection):
    data = requests.get(collection).json()
    for manifest in data["manifests"]:
        print(manifest["@id"])


parse_collection(INITIAL)
