import requests
from requests.auth import HTTPBasicAuth
import json

# dev_api = "https://siterightapi-dev.lp-ase.appserviceenvironment.net/token"

# xform_headers = {'Content-Type': 'application/x-www-form-urlencoded'}
# payload = {
#     'grant_type': 'password', 
#     'username': 'jaron.jones@elmlocating.com',
#     'password': 'Zyzz1234!'
# }

# #Get the token
# b_token_dev = requests.post(url=dev_api, headers=xform_headers, data=payload)

# #Parse the reqeust to get the access token from this shit
# print (b_token_dev)

# #Get the token
# token = b_token_dev['access_token']
response = [{'access_token':'12345'}]
token = ''
try:
    for things in response:
        # if access_token key exists
        if things['access_token']:
            print('found token')
            token = things['access_token']
        else:
            raise Exception('didnt find no access token mang')
except Exception as e:
    print(error)

print(token)