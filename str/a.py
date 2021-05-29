import requests
import time


def isOsmHostOk(host='http://prod-osm1959.hdmap.momenta.works'):
    isOk = False
    for i in range(12):
        time.sleep(10)
        if 'Welcome to OpenStreetMap' in requests.get(host).text:
            isOk = True
            break
    return isOk
print(isOsmHostOk())
