import requests

INITIAL = "https://digital.library.villanova.edu/Collection/vudl:3/IIIF"


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
