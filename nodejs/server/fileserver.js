'use strict';
var fs = require('fs');
var path = require('path');
var url = require('url');
var http = require('http');
var root = path.resolve(process.argv[2] || `.`);
console.log('root is ' + root);

var server = http.createServer((request, response)=>{
    var pathname = url.parse(request.url).pathname;
    var filepath = path.join(root, pathname);
    fs.stat(filepath, (err, stats)=>{
	if (!err && stats.isFile()){
	    console.log('200' + request.url);
	    console.log(request.connection.remoteAddress);
	    response.writeHead(`200`, {'Content-Type': 'text/html;charset=UTF-8'});
	    let fstream = fs.createReadStream(filepath).pipe(response);
	}else{
	    console.log('404' + request.url);
	    response.writeHead(`404`);
	    response.end('404 Not Found!');
	}
    });
});
server.listen('8080');
console.log(`server start at 8080\n`);
