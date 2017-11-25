import urllib.request
for line in urllib.request.urlopen('http://172.19.80.83:81/verify?encoding=utf-8'):
	line = line.decode('utf-8')
	print(line)
