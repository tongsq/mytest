'use strict';
var fs = require('fs');
var stream = fs.createReadStream('simple.txt', 'utf-8');
stream.on('data', (chunk)=>{
    console.log('DATA: '+chunk);
});
stream.on('end', ()=>{
    console.log('the stream is end');
});
stream.on('error', (err)=>{
    console.log('Error: '+err);
});

var wstream = fs.createWriteStream('out.txt');
wstream.write(new Buffer(`以二进制形式写入文件..\n`, 'utf-8'));
wstream.write(new Buffer('End..', 'utf-8'));
wstream.end(); 
