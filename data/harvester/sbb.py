import requests
from lxml import etree

NSMAP = {
    'zs': "http://www.loc.gov/zing/srw/",
    'pica': "info:srw/schema/5/picaXML-v1.0"
}


def sbb(start=1, maximum=100):
    endpoint = "http://sru.k10plus.de/opac-de-1"

    params = {
        "version": "1.1",
        "operation": "searchRetrieve",
        "query": "pica.XBBG=O* and pica.XALL=resolverstaatsbibliothek*",
        "maximumRecords": maximum,
        "recordSchema": "picaxml",
        "startRecord": start
    }
    res = requests.get(endpoint, params=params)
    return res


# Get hits
res = sbb(1, 1)
tree = etree.fromstring(res.content)
hits = tree.find(".//zs:numberOfRecords", NSMAP).text
hits = int(hits)


with open("sbb.txt", "w") as OUT:
    for i in range(1, hits + 1, 100):
        try:
            res = sbb(i, 100)
            tree = etree.fromstring(res.content)
            for rec in tree.findall('.//pica:record', NSMAP):
                ppn = rec.find("pica:datafield[@tag='003@']/pica:subfield[@code='0']", NSMAP)
                ppn = ppn.text
                urls = rec.findall('pica:datafield[@tag="017C"]/pica:subfield[@code="u"]', NSMAP)
                urls = [u.text for u in urls]
                print(f"https://content.staatsbibliothek-berlin.de/dc/{ppn}/manifest", file=OUT)
        except Exception as e:
            print(f"{e} at iteration i = {i}")

