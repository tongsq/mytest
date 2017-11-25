import datetime
now = datetime.date.today()
str = now.strftime('%y-%m-%d')
print('now is {0}, {1}'.format(now, str))

birthday = datetime.date(1993, 9, 10)
age = now - birthday
print('days is {0:d}, age is {1:d}'.format(age.days, int(age.days/365)))

import zlib
s = b'hello world hello world'
print('len is {}'.format(len(s)))
t = zlib.compress(s)
print('len compress is {0:d}'.format(len(t)))
print(zlib.crc32(t), zlib.decompress(t))
