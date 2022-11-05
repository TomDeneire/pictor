# import requests

# SPARQL = "http://sparql.europeana.eu/"
# QUERY = """PREFIX edm: <http://www.europeana.eu/schemas/edm/>
# PREFIX ore: <http://www.openarchives.org/ore/terms/>
# PREFIX dcterms: <http://purl.org/dc/terms/>

# SELECT ?iiif
# WHERE {
# ?object ore:proxyIn ?local_aggr .
# ?local dcterms:isReferencedBy ?iiif .
# }"""

# result = requests.get(SPARQL, params={"format": "json", "query": QUERY}).json()

# for item in result["results"]["bindings"]:
#     link = item["url"]["value"]
#     # print(link)
#     if "iiif" in link:
#         print(link)
