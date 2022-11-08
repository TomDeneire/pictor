import requests
import time

INITIAL = "https://lbiiif.riksarkivet.se/collection/riksarkivet"


def parse_collection(collection):
    # Timeout helps, but still this runs into requests.exceptions.ConnectionError:
    # HTTPSConnectionPool(host='lbiiif.riksarkivet.se', port=443):
    # Max retries exceeded with url: /collection/arkiv/nsCHjheTGawE8aIsztfLo6
    time.sleep(5)
    data = requests.get(collection).json()
    if "items" in data:
        for item in data["items"]:
            if item["type"] == "Manifest":
                print(item["id"])
            elif item["type"] == "Collection":
                for col in data["items"]:
                    parse_collection(col["id"])
    return


parse_collection(INITIAL)
