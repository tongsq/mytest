'use strict';
var fs = require('fs');
fs.readFile('simple.txt', 'utf-8', (err, data)=>{
    if (err)
	console.log(err);
    else
	console.log(data);
});
fs.readFile('simple.txt', (err, data)=>{
    if (err)
	console.log(err);
    else
    {
	console.log(data);
	console.log(data.length + ' bytes');
	let text = data.toString('utf-8');
	console.log(text);
    }	
});
var data = "hello nodejs\n";
fs.writeFile('output.txt', data, (err)=>{
    if (err) console.log(err);
});
