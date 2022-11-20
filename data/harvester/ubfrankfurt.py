import requests
from lxml import etree


def getManifests(tree) -> list:
    manifests = []
    for x in tree.findall('.//{http://dfg-viewer.de/}iiif'):
        manifests.append(x.text)
    return manifests

def resTok(tree) -> str:
    try:
        resumptionToken = tree.find('.//{*}resumptionToken').text
    except:
        resumptionToken = None
    return resumptionToken

def processOAI(res):
    tree = etree.fromstring(res.content)
    tmpManifests = getManifests(tree)
    for manifest in tmpManifests:
        print(manifest)
    global resumptionToken
    resumptionToken = resTok(tree)


url = "http://digitale-sammlungen.ulb.uni-bonn.de/oai/"


# First run
res = requests.get(url, params = {
    'verb' : 'ListRecords',
    'metadataPrefix' : "mets"
})
processOAI(res)

# Iterate till end
while resumptionToken:
    res = requests.get(url, params = {'verb' : 'ListRecords', 'resumptionToken' : resumptionToken })
    processOAI(res)