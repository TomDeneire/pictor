#!/usr/bin/env python
# coding: utf-8

# In[1]:


import requests
from lxml import etree


# In[63]:


NSMAP = {
    'zs' : "http://www.loc.gov/zing/srw/",
    'pica' : "info:srw/schema/5/picaXML-v1.0"
}


# In[52]:


def sru(start = 1, maximum = 100):
    endpoint = "http://sru.gbv.de/opac-de-204"

    params = {
        "version" : "1.1",
        "operation" : "searchRetrieve",
        "query" : "pica.BBG=O* and pica.ALL=resolveriaispk*",
        "maximumRecords" : maximum,
        "recordSchema" : "picaxml",
        "startRecord" : start
    }
    res = requests.get(endpoint, params = params)
    return res


# In[69]:


# Get hits
res = sru(1,1)
tree = etree.fromstring(res.content)
hits = tree.find(".//zs:numberOfRecords", NSMAP).text
hits = int(hits)


# In[70]:


with open("iai.txt", "w") as OUT:
    for i in range(1, hits + 1, 100):
        try:
            res = sru(i,100)
            print(res.url)
            tree = etree.fromstring(res.content)
            for rec in tree.findall('.//pica:record', NSMAP):
                ppn = rec.find("pica:datafield[@tag='003@']/pica:subfield[@code='0']", NSMAP)
                ppn = ppn.text
                urls = rec.findall('pica:datafield[@tag="017C"]/pica:subfield[@code="u"]', NSMAP)
                urls = [u.text for u in urls]
                print(f"https://digital.iai.spk-berlin.de/viewer/api/v1/records/{ppn}/manifest/", file = OUT)
        except Exception as e:
            print(f"{e} at iteration i = {i}")


# In[ ]:




