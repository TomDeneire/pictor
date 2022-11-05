import requests

INITIAL = "https://iiif.library.utoronto.ca/presentation/v2/collections/utl:root"


def parse_collection(collection):
    data = requests.get(collection).json()
    if "manifests" in data:
        for manifest in data["manifests"]:
            print(manifest["@id"])
    if "collections" in data:
        for col in data["collections"]:
            parse_collection(col["@id"])
    return


parse_collection(INITIAL)
