// var http = require('http');
// var url = 'http://www.baidu.com';
// http.get(url,function(res){
// 	var html = '';
// 	res.on('data',function(data){
// 		html += data;
// 	});
// 	res.on('end',function(){
// 		console.log(html);
// 	});
// }).on('error',function(){
// 	console.log('获取失败');
// });
var EventEmitter = require('events').EventEmitter;
var life = new EventEmitter();
life.setMaxListeners(11);
life.on('out',function(who){
	console.log('请',who,'滚出去');
});
function water(who) {
	console.log('请',who,'喝水');
}
life.on('out',water);
life.emit('out','小明');
life.removeListener('out',water);
console.log('监听函数数量',life.listeners('out').length);
life.removeAllListeners('out');
console.log(EventEmitter.listenerCount(life,'out'));