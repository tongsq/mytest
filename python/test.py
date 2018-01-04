print('1027 * 768 =', 1024*768)
print('''hello
world"
...''')
print(r"hello\"world\"")
def myabs(num):
	if num >= 0:
		return num
	else:
		return -num
print('myabs(-10) =', myabs(-10))
def sum(*nums):
	sum = 0
	for num in nums:
		sum += num
	return sum
print('sum(1,2,3) =', sum(1,2,3))

def fact(n):
	if n == 1:
		return 1
	return n * fact(n-1)
print('fact(100):', fact(100))
def newfact(n):
	return fact_iter(n, 1)
def fact_iter(num, product):
	if num == 1:
		return product
	return fact_iter(num - 1, num * product)
#print('fact(1000):', fact(1000))
#print('newfact(1000):', newfact(1000))
L = [0,1,2,3,4,5,6]
print('L[0:3] :', L[0:3])
print('L[:3] :', L[:3])
print('L[:6:2]', L[:6:2])
for key,value in enumerate(L):
	print('key =',key,' value=',value,'',end='')
D = {'a':111,'b':222,'c':333}
for k,v in D.items():
	print('k =',k,'v =',v)
for k in D:
	print('k =', k)
for v in D.values():
	print('v =', v)
print([x*x for x in range(0,3)])
#generator
g = (x*x for x in range(0,3))
for value in g:
	print('%s,'%(value), end='')
def fib(max):
	n,a,b = 0,0,1
	while(n<max):
		yield b
		a,b = b,a+b
		n = n + 1
	return 'done'
f = fib(5)
while True:
	try:
		x = next(f)
		print('f:',x)
	except StopIteration as e:
		print('Generator return value:', e.value)
		break
def f2(num):
	return num * 2
print(list(map(f2, [1,2,3,4,5])))
def f3(a, b):
	print(a,b)
	return 10 * a + b
from functools import reduce
print(reduce(f3, [1,2,3,4,5]))
def is_odd(*nums):
	return list(filter(lambda x:x%2==0, nums))
print(is_odd(1,2,3,4,5,6))
