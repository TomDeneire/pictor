import requests

FRICK = "https://collections.frick.org/apis/iiif/presentation/v2/objects"


with open("frick.txt", "a") as writer:

    for i in range(1, 1000000):
        url = f"{FRICK}-{i}/manifest"
        try:
            res = requests.get(url, timeout=10)
            test = res.status_code
            if test == 200:
                writer.write(url + "\n")
                print(url)
        except Exception:
            pass
