f = open('a.txt', 'r+')
f.write('hello world')
print(f.readline())
for line in f:
	print(line)
print(f)
f.close()
print(f)
with open('a.txt') as fs:
	for line in fs:
		print(line)
print(fs)
