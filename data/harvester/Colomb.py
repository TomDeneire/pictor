#!/usr/bin/env python
# coding: utf-8

# In[38]:


import requests
import sys


# In[28]:


for i in range(1,99):
    toplevel = f"https://babel.banrepcultural.org/iiif/2/p17054coll{i}/manifest.json"


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
            'collections' : res.get('collections')
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

for i in range(1,99):
    toplevel = f"https://babel.banrepcultural.org/iiif/2/p17054coll{i}/manifest.json"
    drilldownIIIF(toplevel)