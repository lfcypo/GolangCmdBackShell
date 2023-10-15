import requests
import base64
import json

while True:

    shell = input("shell: ->   ")
    shell = base64.b64encode(shell.encode("utf-8")).decode("utf-8")

    response = requests.post(f"http://10.78.11.94:20086/run_shell?shell={shell}")
    if response.status_code != 200:
        print("Faild")
        continue

    if str(json.loads(response.text)["code"]) != "200":
        print(str(json.loads(response.text)["msg"]))
        continue

    print(base64.b64decode(str(json.loads(response.text)["data"]).encode("gb2312")).decode("gb2312"))
    continue
