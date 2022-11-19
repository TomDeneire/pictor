import requests

INITIAL = "https://collectiones.manuscriptorium.com"


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
