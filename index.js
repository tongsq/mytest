const http = require('http');

const hostname = '127.0.0.1';
const port = 4747;

http.createServer((req, res) => {
  res.writeHead({'Content-Type':'chare-set="utf-8"'});
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('Hello World!\n');
}).listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
  console.log('the port is ${port}');
  console.log('I have add a dev');
});