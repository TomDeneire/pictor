#!/usr/bin/env python
# coding: utf-8

# In[38]:


import requests
import sys


# In[28]:


toplevel = "https://cdm16028.contentdm.oclc.org/iiif/p15324coll10/manifest.json"


# In[37]:


def parseIIIF(manifest):
    print(f"Trying to open {manifest}", file = sys.stderr)
    try:
        res = requests.get(manifest, timeout = 10)
        res = res.json()
        return {
            'id' : res.get('@id'),
            'type' : res.get('@type'),
            'manifests' : res.get('manifests'),
            'collections' : res.get('collections'),
            'first' : res.get('first'),
            'next' : res.get('next')
        }
    except Exception as e:
        print(e, file = sys.stderr)
        print(f"PROBLEM with {manifest}", file = sys.stderr)

def drilldownIIIF(manifest):
    print(f"new drilldown with {manifest}", file = sys.stderr)
    iiif = parseIIIF(manifest)
    try:
        for m in iiif.get('manifests'):
            print(m.get('@id'))
    except Exception as e:
        print(e, file = sys.stderr)
    try:
        for c in iiif.get('collections'):
            drilldownIIIF(c.get('@id'))
    except Exception as e:
        print(e, file = sys.stderr)

manifest = parseIIIF(toplevel)
first = manifest.get('first')
print(first)
drilldownIIIF(first)
manifest = parseIIIF(first)
NEXT = manifest.get('next')
print(NEXT)
drilldownIIIF(NEXT)

while NEXT:
    manifest = parseIIIF(NEXT)
    print(NEXT)
    drilldownIIIF(NEXT)
    NEXT = manifest.get('next')