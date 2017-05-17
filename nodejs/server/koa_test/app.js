'use strict';
const Koa = require('koa');
const router = require('koa-router')();
const app = new Koa();
const bodyParser = require('koa-bodyparser');
app.use(async (ctx, next)=>{
    let stime = new Date().getTime();
	console.log(ctx.request.body);
    await next();
    var time = new Date().getTime() - stime;
    console.log(`Process ${ctx.request.method} ${ctx.request.url}...${time}ms`);
});
router.get(`/hello/:name`, async(ctx, next)=>{
    var name = ctx.params.name;
    ctx.response.body = `<h1>Hello, ${name}!</h1>`;
});
router.get(`/`, async (ctx, next)=>{
    ctx.response.body = `<h1>Index</h1>
	<form action="/signin" method="post"/>
	<p>Name: <input name="name" value="koa"></p>
	<p>Passwd: <input name="passwd" type="passwd" value="123456"/></p>
	<p><input type="submit" value="submit"></p>
</form>`;
});
router.post('/signin', async (ctx, next)=>{
    var name = ctx.request.body.name || '';
	passwd = ctx.request.body.passwd || '';
	if (name === 'koa' && passwd === '123456'){
	    ctx.response.body = `<h1>Welcome, ${name}!</h1>`;
	}else{
	    ctx.response.body = `<h1>Login failed!</h1>
	    <p><a href="/"> try again</a></p>`;
	}
});
const controller = require('./controller');
app.use(bodyParser());
const number = 100;
//app.use(router.routes());
app.use(controller());
app.listen(8080);
console.log(`app start at port 8080...\n`);
