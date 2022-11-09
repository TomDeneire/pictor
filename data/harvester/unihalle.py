# initial code by https://github.com/alexander-winkler

import requests
from lxml import etree


def resTok(tree):
    """
    Parses an etree object and returns the resumptionToken
    """
    resumptionToken = tree.find(
        ".//{http://www.openarchives.org/OAI/2.0/}resumptionToken"
    ).text
    return resumptionToken


def manifestURLs(tree) -> list:
    """
    Extracts the iiif URLs from an etree objecte and returns them as a list
    """
    return [x.text for x in tree.findall(".//{http://dfg-viewer.de/}iiif")]


def addManifestsFromTree(tree):
    """
    Extracts the URLs of the IIIF-manifests from an ElementTree and adds them to a list
    """
    tmpManifests = manifestURLs(tree)
    for manifest in tmpManifests:
        print(manifest)


manifests = []
url = "https://digitale.bibliothek.uni-halle.de/oai"
params = {
    "verb": "ListRecords",
    "metadataPrefix": "mets",
}

# First run
res = requests.get(url, params)
tree = etree.fromstring(res.content)
addManifestsFromTree(tree)
resumptionToken = resTok(tree)
params["resumptionToken"] = resumptionToken
del params[
    "metadataPrefix"
]  # ResumptionToken is only argument, so this one has to be deleted

# 2nd to nth run to collect manifest URLS
while resumptionToken:
    res = requests.get(url, params)
    tree = etree.fromstring(res.content)
    addManifestsFromTree(tree)
    params["resumptionToken"] = resumptionToken
    resumptionToken = resTok(tree)
