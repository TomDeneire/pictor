manifests = set()
with open("BIIIF.zwr", "r") as reader:
    for line in reader:
        if not line.startswith("^BIIIF"):
            continue
        identifier = line.split(",")[1].strip('"')
        if ":" not in identifier:
            manifest = f"https://anet.be/iiif/{identifier}/manifest"
            manifests.add(manifest)

for manifest in manifests:
    print(manifest)
