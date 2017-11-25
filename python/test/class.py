class animal:
	__name = ''
	def __init__(self, name):
		self.__name = name
	def say(self):
		print('my name is {0}'.format(self.__name))

class person(animal):
	age = 0;
	def __init__(self, name, age):
		self.name = name
		self.age = age
		animal.__init__(self, name)
	def sayage(self):
		print('my age is {0:d}'.format(self.age))

p = person('tom', 18)
p.say()
p.sayage()
