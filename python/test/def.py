def showPerson(name, age=20, *vartuple):
	print('name: ', name)
	print('age: ', age)
	for var in vartuple:
		print(var)
showPerson(age=18, name='tom')
showPerson('tom', 18, 'aaa')

#匿名函数
sum = lambda arg1,arg2:arg1+arg2
print('1 + 2 = ', sum(1, 2))
#nonlocal  global
Flag = 1
def showflag():
#	print(Flag)
#	Flag = 2
	print(Flag)

showflag()
print(Flag)
