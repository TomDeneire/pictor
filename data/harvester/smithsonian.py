from urllib.request import urlopen

INITIAL = "https://collections.si.edu/search/results.htm?q=&iiif.enabled=true"

url = INITIAL

with urlopen(url) as reader:
    content = reader.read()
    print(content)
