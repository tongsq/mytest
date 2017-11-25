print('hello world\n', end='')
import sys
sys.stdout.write('hello world 2\n')
for i in sys.argv:
	print(i)

list = ['a','b','c','d','e']
print(list[1:3])

tuple = (1, 2, 3, 'aa', 'bb')
print(tuple)
print(tuple[3:])
#set
student = {'tom', 'tim', 'tom', 'rose'}
print(student)
if ('rose' in student):
	print('rose in student')
else:
	print('rose not in student')
#dict
dict = {'name':'tom', 'age':18}
print(dict)

print ("我叫 %s 今年 %d 岁!" % ('小明', 10))
for i in range(5):
	print(i)

def fib(n):
	a, b, count = 0, 1, 0
	while count<n:
		yield a
		a, b = b, a+b
		count+=1
nums = fib(10)
for num in nums:
#	pass
	print(num, end=' ')
for num in nums:
	print(num, end=',')
while True:
	try:
		print (next(nums), end=" ")
	except StopIteration:
		print ('iter end\n')
		break


