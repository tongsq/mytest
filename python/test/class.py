class animal:
	public = 'animal'
	__name = ''
	#__slots__ = ('name')
	#name = ''
	def __init__(self, name):
		self.__name = name
		self.name = name
	def say(self):
		print('parent:my name is {0}'.format(self.name))
		print('parent:my __name is {0}'.format(self.__name))

class person(animal):
	#__slots__ = ('name', 'age')
	#age = 0;
	def __init__(self, name, age):
		self.__name = 'aa'
		self.name = 'bb'
		self.age = age
		animal.__init__(self, name)
	def say(self):
		animal.say(self)
		print('child: my __name is {0}'.format(self.__name))
		print('child: my name is {0}'.format(self.name))
	def sayage(self):
		print('child: my age is {0:d}'.format(self.age))
	@property
	def color(self):
		return self._color
	@color.setter
	def color(self, value):
		if (value>0 and value < 200):
			self._color = value
		else:
			print('error value of color')
p = person('tom', 18)
p.say()
p.sayage()
print(animal.public)
print(p.public)
animal.public = 'testttt'
print(person.public)
print(p.public)
def sayage2(self):
	print('sayage bind by extrai age is %d' % (self.age))
animal.sayage2 = sayage2
p.sayage2()
from types import MethodType
sayage3 = sayage2
p.sayage3 = MethodType(sayage3, p)
p.sayage3()
p.name = 'aa'
p.name2 = 'bb'
print(p.name2)
p.color = 88
print('p.color is %d' % (p.color))
p.color = 300
