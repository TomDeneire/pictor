# Docu on Harvard-API at https://github.com/harvardartmuseums/api-docs

import requests

apikey = "" # to be obtained via https://docs.google.com/forms/d/1Fe1H4nOhFkrLpaeBpLAnSrIMYvcAxnYWm0IU9a6IkFA/viewform


res = requests.get(f"https://api.harvardartmuseums.org/Object?apikey={apikey}", timeout = 10)
NEXT = res.json().get('info').get('next')
recIDs = [x.get('id') for x in res.json().get('records')]

while NEXT:
    res = requests.get(NEXT, timeout = 10)
    NEXT = res.json().get('info').get('next')
    recIDs.extend([x.get('id') for x in res.json().get('records')])

with open("harvard.txt", "w") as OUT:    
    for r in recIDs:
        print(f"https://iiif.harvardartmuseums.org/manifests/object/{r}", file = OUT)