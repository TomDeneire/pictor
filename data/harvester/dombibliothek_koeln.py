import lxml.etree
import urllib.request
import urllib.error

INITIAL = "https://digital.dombibliothek-koeln.de/hs/oai?verb=ListRecords&set=ddbkha&metadataPrefix=oai_dc"
MANIFEST_URL = "https://digital.dombibliothek-koeln.de/hs/i3f/v20/{id}/manifest"
OAI_XML_NS = r"{http://www.openarchives.org/OAI/2.0/}"  # raw string is safer

url = INITIAL

while url:
    try:
        with urllib.request.urlopen(url) as query:
            response = query.read()
    except (urllib.error.HTTPError, urllib.error.URLError) as err:
        print(url, err)
        break
    root = lxml.etree.fromstring(response)  # type: ignore
    for identifier in root.iter(f"{OAI_XML_NS}identifier"):
        iiifid = identifier.text.partition("oai:digital.dombibliothek-koeln.de:")[2]
        manifest = MANIFEST_URL.replace("{id}", iiifid)
        print(manifest)
    resumptionToken = None
    for token in root.iter(f"{OAI_XML_NS}resumptionToken"):
        resumptionToken = token.text
    if resumptionToken:
        url = INITIAL.partition("&set")[0] + "&resumptionToken=" + resumptionToken
    else:
        url = None
