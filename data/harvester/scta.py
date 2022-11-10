import requests

INITIAL = "https://scta.info/iiif/authors/collection"


def parse_collection(collection):
    data = requests.get(collection).json()
    if "manifests" in data:
        for manifest in data["manifests"]:
            if "iiif" in manifest["@id"]:
                print(manifest["@id"])
    elif "collections" in data:
        for col in data["collections"]:
            parse_collection(col["@id"])
    else:
        return
    return


parse_collection(INITIAL)
