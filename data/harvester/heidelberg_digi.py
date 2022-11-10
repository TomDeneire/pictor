import requests
from lxml import etree

# Info on APIs: https://www.ub.uni-heidelberg.de/helios/kataloge/datenschnittstellen.html 


# initialize
URL = "https://digi.ub.uni-heidelberg.de/cgi-bin/digioai.cgi?verb=ListRecords&metadataPrefix=oai_dc"
res = requests.get(URL)
tree = etree.fromstring(res.content)
manifests = [x.text for x in tree.findall('.//{*}hasFormat')]
resToken = tree.find('.//{*}resumptionToken').text
while resToken:
    print(resToken)
    url = f"https://digi.ub.uni-heidelberg.de/cgi-bin/digioai.cgi?verb=ListRecords&resumptionToken={resToken}"
    res = requests.get(url)
    tree = etree.fromstring(res.content)
    manifests.extend([x.text for x in tree.findall('.//{*}hasFormat')])
    try:
        resToken = tree.find('.//{*}resumptionToken').text
    except:
        resToken = None

with open('heidelberg_digi.txt','w') as OUT:
    for m in set(manifests):
        print(m, file = OUT)