class Fib(object):
	def __init__(self):
		self.a, self.b = 0, 1

	def __iter__(self):
		return self

	def __next__(self):
		self.a, self.b = self.b, self.a + self.b
		if self.a > 100000:
			raise StopIteration()
		return self.a
	def __getitem__(self, n):
		if isinstance(n, int):
			a,b = 1, 1
			for x in range(n):
				a, b = b, a + b
			return a
		if isinstance(n, slice):
			start = n.start
			stop = n.stop
			if start is None:
				start = 0
			a, b = 1, 1
			L = []
			for x in range(stop):
				if x >= start:
					L.append(a)
				a, b = b, a + b
			return L
	def __getattr__(self, attr):
		if attr == 'score':
			return 90
		raise AttributeError('Fib object has no attribute {0}'.format(attr))
if __name__ == '__main__':
	fib = Fib()
	for value in fib:
		if (value<100):
			print(value, end=',')
		else:
			break
	print('Fib[3] = ',fib[3])
	print('Fib[1:4] = ', fib[1:4])
	print('Fib.score = ', fib.score)
	print('Fib.score2', fib.score2)
