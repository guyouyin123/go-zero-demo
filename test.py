import json

import requests


def user_main1():
    url = "http://127.0.0.1:8888/dev/v1/user/info"
    data = {
        "userId":1
    }
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.post(url,headers=headers,data=json.dumps(data))
    print(res.text)

def user_rpc_main2():
    url = "http://127.0.0.1:8888/dev/v1/rpc/user/info"
    data = {
        "userId":2
    }
    headers = {
        'Content-Type': 'application/json'
    }


    res = requests.post(url,headers=headers,data=json.dumps(data))
    print(res.text)

def main2():
    url = "http://127.0.0.1:8888/dev/v1/bike/info?bikeId=1"
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.get(url,headers=headers)
    print(res.text)

def main3():
    url = "http://127.0.0.1:8888/dev/v1/bike/insert"
    data = {
        "bikeId":110,
        "bikeName":"R01车型"
    }
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.post(url,headers=headers,data=json.dumps(data))
    print(res.text)





if __name__ == '__main__':
    user_rpc_main2()