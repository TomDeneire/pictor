import requests
from lxml import etree
from urllib.parse import urlparse
import sys




def getIDs(tree) -> list:
    IDs = []
    for x in tree.findall('.//{http://www.openarchives.org/OAI/2.0/}identifier'):
        path = urlparse(x.text).path
        path = "/".join(path.split('/')[2:])
        IDs.append(path)
    return IDs

def resTok(tree) -> str:
    try:
        resumptionToken = tree.find('.//{*}resumptionToken').text
    except:
        resumptionToken = None
    return resumptionToken

def processOAI(res):
    tree = etree.fromstring(res.content)
    tmpIDs = getIDs(tree)
    for ID in tmpIDs:
        print(f"https://iiif.europeana.eu/presentation/{ID}/manifest")
    global resumptionToken
    resumptionToken = resTok(tree)


# First run
res = requests.get("https://api.europeana.eu/oai/record?verb=ListIdentifiers&metadataPrefix=edm")

processOAI(res)

# Iterate till end
while resumptionToken:
    res = requests.get(f"https://api.europeana.eu/oai/record?verb=ListIdentifiers&resumptionToken={resumptionToken}")
    print(res.url, file = sys.stderr)
    processOAI(res)