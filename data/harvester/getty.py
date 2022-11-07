import requests

SPARQL = "https://data.getty.edu/museum/collection/sparql"
QUERY = """PREFIX crm: <http://www.cidoc-crm.org/cidoc-crm/>
SELECT ?iiif WHERE {
  ?item crm:P129i_is_subject_of ?iiif .
  FILTER(CONTAINS(STR(?iiif),'iiif/manifest'))
}"""

result = requests.get(SPARQL, params={"format": "json", "query": QUERY}).json()

for item in result["results"]["bindings"]:
    link = item["iiif"]["value"]
    print(link)
