const url = 'http://192.168.2.254/ws/NeMo/Intf/lan:getMIBs';
const options = {
  method: 'POST',
  headers: {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:123.0) Gecko/20100101 Firefox/123.0',
    Accept: '*/*',
    'Accept-Language': 'nl,en-US;q=0.7,en;q=0.3',
    'Accept-Encoding': 'gzip, deflate',
    'Content-Type': 'application/x-sah-ws-4-call+js; charset=utf-8',
    Authorization: 'X-Sah-Login',
    Origin: 'http://192.168.2.254',
    Connection: 'keep-alive',
    Referer: 'http://192.168.2.254/'
  },
  body: '{"service":"sah.Device.Information","method":"createContext","parameters":{"applicationName":"webui","username":"admin","password":"Y~XF2w}t3"}}'
};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
