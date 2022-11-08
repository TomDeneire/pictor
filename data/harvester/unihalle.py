import lxml.etree
import urllib.request
import urllib.error

INITIAL = (
    "https://digitale.bibliothek.uni-halle.de/oai?verb=ListRecords&metadataPrefix=mets"
)
DV_NS = r"{http://dfg-viewer.de/}"  # raw string is safer
OAI_NS = r"{http://www.openarchives.org/OAI/2.0/}"  # raw string is safer

url = INITIAL

while url:
    try:
        # unihalle tends to timeout rather quickly
        with urllib.request.urlopen(url, timeout=10) as query:
            response = query.read()
    except (urllib.error.HTTPError, urllib.error.URLError) as err:
        print(url, err)
        break
    root = lxml.etree.fromstring(response)  # type: ignore
    for identifier in root.iter(f"{DV_NS}iiif"):
        iiifid = identifier.text
        print(iiifid)
    resumptionToken = None
    for token in root.iter(f"{OAI_NS}resumptionToken"):
        resumptionToken = token.text
    if resumptionToken:
        url = INITIAL.replace(
            "&metadataPrefix=mets", f"&resumptionToken={resumptionToken}"
        )
    else:
        url = None
