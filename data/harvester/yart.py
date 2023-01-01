import requests
import re
from urllib import parse

PREFIX = "https://artgallery.yale.edu"

URLS = set()
MANIFESTS = set()


def find_manifest(url: str) -> None:
    URLS.add(url)
    try:
        page = requests.get(url)

        # oh boy, parsing html with regex :-)
        hrefs = re.findall(r'href=".*?"', page.text)
        for href in hrefs:
            href = parse.unquote(href)
            if "manifest=https://manifests.collections.yale.edu" in href:
                manifest = href.partition("manifest=")[2].partition("&canvas")[0].strip('"')
                if manifest not in MANIFESTS:
                    MANIFESTS.add(manifest)
                    print(manifest)
    except Exception:
        return


for i in range(0, 18402):
    url = f"{PREFIX}/collection/search?page={i}"
    try:
        page = requests.get(url)
        # oh boy, parsing html with regex :-)
        hrefs = re.findall(r'href=".*?"', page.text)
        for href in hrefs:
            if "collections/objects" in href:
                record_url = PREFIX + href.replace('href="', "").strip('"')
                if record_url not in URLS:
                    find_manifest(record_url)

    except Exception:
        continue


