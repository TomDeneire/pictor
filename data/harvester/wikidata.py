import requests

SPARQL = "https://query.wikidata.org/sparql"
QUERY = "SELECT distinct ?iiif WHERE { ?item wdt:P6108 ?iiif . }"

result = requests.get(SPARQL, params={"format": "json", "query": QUERY}).json()

for item in result["results"]["bindings"]:
    print(item["iiif"]["value"])
