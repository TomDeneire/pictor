# Info on Art Institute Chicago API at https://api.artic.edu/docs/#iiif-image-api
import requests

with open("art_inst_chicago.txt", "w") as OUT:
    url = "https://api.artic.edu/api/v1/artworks"

    res = requests.get(url, timeout=10)
    iiif_url = res.json().get('config').get('iiif_url')
    for i in res.json().get('data'):
        print(f"{iiif_url}/{i.get('image_id')}", file=OUT)
    next_url = res.json().get('pagination').get('next_url')

    while next_url:
        print(next_url)
        res = requests.get(next_url, timeout=10)
        iiif_url = res.json().get('config').get('iiif_url')
        for i in res.json().get('data'):
            print(f"{iiif_url}/{i.get('image_id')}", file=OUT)

            next_url = res.json().get('pagination').get('next_url')
