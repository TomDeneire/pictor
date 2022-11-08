import requests

INITIAL = "https://iiif.bodleian.ox.ac.uk/iiif/activity/page-0"

url = INITIAL

while url:
    try:
        data = requests.get(url).json()
    except Exception as e:
        print(url)
        break
    for item in data["orderedItems"]:
        if item["object"]["type"] == "Manifest":
            print(item["object"]["id"])
    next_page = data.get("next")
    if next_page:
        url = next_page.get("id")
    else:
        url = None
