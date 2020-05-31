class Chain(object):

	def __init__(self, path=''):
		self._path = path

	def __getattr__(self, path):
		return Chain('%s/%s' % (self._path, path))

	def __str__(self):
		return self._path
	
	def __call__(self, name=''):
		if name != '':
			rindex = self._path.rfind('/')
			return Chain('%s/%s' % (self._path[:rindex], name))
			list = self._path.split('/')[:-1]
			return Chain('%s/%s' % ('/'.join(list), name))
		else:
			return Chain('%s/%s' % (self._path, name))
	__repr__ = __str__

	

if (__name__ == '__main__'):
	print(Chain().status.user.timeline.list)
	print(Chain().status.user('jack').timeline.list) 
