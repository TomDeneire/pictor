import requests

SPARQL = "https://www.goudatijdmachine.nl/sparql/repositories/gtm"
QUERY = """SELECT ?iiif WHERE {
?s a <https://schema.org/ImageObject> .
    ?s <http://omeka.org/s/vocabs/o#id> ?id .
    BIND(CONCAT('https://www.goudatijdmachine.nl/data/iiif/2/',STR(?id),'/manifest') AS ?iiif)
}"""

result = requests.get(SPARQL, params={"query": QUERY})


for line in result.text.split("\n"):
    if line == "":
        continue
    print(line)
