
import requests

INITIAL = "https://iiif.wellcomecollection.org/presentation/v2/collections"


def parse_collection(collection):
    data = requests.get(collection).json()
    if "collections" in data:
        for col in data["collections"]:
            parse_collection(col["@id"])
    elif "members" in data:
        for memb in data["members"]:
            if memb["@type"] == "sc:Collection":
                identifier = memb["@id"]
                if identifier == collection:
                    continue
                parse_collection(memb["@id"])
            elif memb["@type"] == "sc:Manifest":
                print(memb["@id"])
    return


parse_collection(INITIAL)
