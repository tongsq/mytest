var fn_index = async (ctx, next)=>{
    ctx.response.body = `<h1>Index</h1>
        <form action="/signin" method="post"/>
        <p>Name: <input name="name" value="koa"/></p>
        <p>Passwd: <input name="passwd" type="password" value="123456"></p>
        <p><input type="submit" value="submit"></p>
	</form>`;
};

var fn_signin = async (ctx, next)=>{
    var name = ctx.request.body.name || '',
        passwd = ctx.request.body.passwd || '';
    if (name === `koa` && passwd === `123456`){
        ctx.response.body = `<h1>Welcome, ${name}</h1>`;
    }else{
        ctx.response.body = `<h1>Login failed!</h1><p><a href="/">try again</a></p>`;
    }
}
module.exports = {
    'GET /' : fn_index,
    'POST /signin' : fn_signin
};
