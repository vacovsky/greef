import requests

def read_latest_vals():
    data = requests.get("http://localhost:3000/data").text
    return data
