#!/usr/bin/env python
# coding: utf-8

# In[53]:


import requests
from lxml import etree
from io import StringIO
from urllib.parse import urljoin
import sys


# In[ ]:


pageURL = "http://iiif.sld.cu/sitemapcolletion"


# In[44]:


def pagelinks(pageURL):
    html = requests.get(pageURL).content
    parser = etree.HTMLParser()
    tree   = etree.fromstring(html, parser)
    links = []
    for l in tree.findall('.//a[@href]'):
        try:
            href = l.attrib.get('href')
            if href.startswith("/coleccion"):
                links.append(urljoin(pageURL,href))
        except Exception as e:
            print(e, file = sys.stderr)
    return links


# In[55]:


def manifests(pageURL):
    html = requests.get(pageURL).content
    parser = etree.HTMLParser()
    tree   = etree.fromstring(html, parser)
    for l in tree.findall('.//a[@href]'):
        try:
            href = l.attrib.get('href')
            if "manifest" in href:
                print(urljoin(pageURL,href))
        except Exception as e:
            print(e, file = sys.stderr)



# In[14]:

links = pagelinks(pageURL)

for l in links:
    manifests(l)