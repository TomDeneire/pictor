import internetarchive

#  search = internetarchive.search_items('identifier:*')  # all (25M+ identifiers)
search = internetarchive.search_items('collection:image')

for index, result in enumerate(search):
    print("https://iiif.archivelab.org/iiif/" + result["identifier"] + "/manifest.json")
