import http.client

conn = http.client.HTTPConnection("modem.uranus.aawa.nl")

payload = '{"service": "sah.Device.Information","method": "createContext","parameters": {"applicationName": "webui","username": "admin","password": "Y~X\\"F2w}t3"}}'
payload = '{"service":"sah.Device.Information","method":"createContext","parameters":{"applicationName":"webui","username":"admin","password":"Y~X\\"F2w}t3"}}'

headers = {
    "Content-Type": "application/x-sah-ws-4-call+js; charset=utf-8",
    "Authorization": "X-Sah-Login",
}

conn.request("POST", "/ws/NeMo/Intf/lan:getMIBs", payload, headers)

res = conn.getresponse()
data = res.read()

print(res.getheaders())

print(data.decode("utf-8"))
