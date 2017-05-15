function test(){
    console.log(typeof arguments);
    console.log(arguments);
}
test(1, 2);
test('a', 'b', 'c');
var map = new Map();
map.set('a', 123);
console.log(typeof map);
console.log(map.get('a'));
console.log(map);
var set = new Set();
set.add(123)
console.log(set);
