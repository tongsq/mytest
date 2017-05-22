var fn_hello = async (ctx, next)=>{
    var name = ctx.params.name;
    //ctx.response.body = `<h1>Hello, ${name}!</h1>`;
    console.log('render in hello.js');
    ctx.render('hello.html', {name:name});
};

module.exports = {
    'GET /hello/:name' : fn_hello
};
