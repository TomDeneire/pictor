import requests

SPARQL = "https://data.getty.edu/museum/collection/sparql"
QUERY = """PREFIX crm: <http://www.cidoc-crm.org/cidoc-crm/>
        SELECT ?iiif WHERE {?item crm:P129i_is_subject_of ?iiif}"""

result = requests.get(SPARQL, params={"format": "json", "query": QUERY}).json()

for item in result["results"]["bindings"]:
    link = item["iiif"]["value"]
    if "iiif/manifest" in link:
        print(link)
