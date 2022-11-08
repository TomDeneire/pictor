# import requests
# import yaml
#
# URL = (
#    "https://github.com/allmaps/iiif-map-collections/blob/main/iiif-map-collections.yml"
# )
#
# YAML_FILE = "allmaps.yml"
#
# data = requests.get(URL)
# with open(YAML_FILE, "w") as writer:
#    writer.write(data.text)
#
# with open(YAML_FILE, "r") as stream:
#    try:
#        print(yaml.safe_load(stream))
#    except yaml.YAMLError as exc:
#        print(exc)

# parsing error in yaml! instead ->

# curl -O https://github.com/allmaps/iiif-map-collections/blob/main/iiif-map-collections.yml
# grep https.*manifest *.yml
