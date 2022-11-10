"""
Doesn't work yet...
"""

import lxml.etree
import urllib.request
import urllib.error

INITIAL = "https://api.europeana.eu/oai/record?verb=ListIdentifiers&metadataPrefix=edm"
GETURL = (
    "https://api.europeana.eu/oai/record?verb=GetRecord&metadataPrefix=edm&identifier="
)
OAI_XML_NS = r"{http://www.openarchives.org/OAI/2.0/}"  # raw string is safer
OAI_DC_NS = r"{http://purl.org/dc/terms/}"


def get_manifest(identifier: str) -> str:
    url = GETURL + identifier
    try:
        with urllib.request.urlopen(url, timeout=30) as query:
            response = query.read()
    except (urllib.error.HTTPError, urllib.error.URLError) as err:
        print(url, err)
        exit(1)
    root = lxml.etree.fromstring(response)  # type: ignore
    for iiif in root.iter(f"{OAI_DC_NS}isReferencedBy"):
        print(iiif.text)
    return ""


url = INITIAL

while url:
    try:
        with urllib.request.urlopen(url, timeout=30) as query:
            response = query.read()
    except (urllib.error.HTTPError, urllib.error.URLError) as err:
        print(url, err)
        break
    root = lxml.etree.fromstring(response)  # type: ignore
    for identifier in root.iter(f"{OAI_XML_NS}identifier"):
        record = identifier.text
        get_manifest(record)
    resumptionToken = None
    for token in root.iter(f"{OAI_XML_NS}resumptionToken"):
        resumptionToken = token.text
    if resumptionToken:
        token = "&resumptionToken=" + resumptionToken
        url = INITIAL.replace("&metadataPrefix=edm", token)
    else:
        url = None
