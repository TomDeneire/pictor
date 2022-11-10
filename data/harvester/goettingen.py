import requests
from lxml import etree


def opac(start: int):
    opac = "http://sru.gbv.de/opac-de-7"
    params = {
        "version": "1.1",
        "operation": "searchRetrieve",
        "query": "pica.url = resolver.sub.unigoettingen.depurlppn*",
        "maximumRecords": "100",
        "recordSchema": "picaxml",
        "startRecord": start,
    }
    res = requests.get(opac, params=params)
    tree = etree.fromstring(res.content)
    return tree


hits = opac(1).find(".//{http://www.loc.gov/zing/srw/}numberOfRecords").text
hits = int(hits)

manifests = []
for i in range(1, hits + 1, 100):
    tree = opac(i)
    for ppn in tree.findall(
        './/{info:srw/schema/5/picaXML-v1.0}datafield[@tag="003@"]/{info:srw/schema/5/picaXML-v1.0}subfield[@code="0"]'
    ):
        manifest = f"https://manifests.sub.uni-goettingen.de/iiif/presentation/PPN{ppn.text}/manifest"
        print(manifest)
