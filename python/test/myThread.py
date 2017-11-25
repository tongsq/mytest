import threading
import time

exitFlag = 0

class myThread (threading.Thread):
	def __init__(self, threadID, name, counter):
		threading.Thread.__init__(self)
		self.threadID = threadID
		self.name = name
		self.counter = counter
	
	def run(self):
		print('start thread: ' + self.name)
		print_time(self, 5)
		print('exit thread: ' + self.name)

def print_time(self, counter):
	while counter:
		if exitFlag:
			exit()
			break
		time.sleep(self.counter)
		print('%s: %s' % (self.name, time.ctime(time.time())))
		counter -= 1

thread1 = myThread(1, 'Thread-1', 1)
thread2 = myThread(2, 'thread2', 2)
print(dir(threading.Thread))
#exit()
thread1.start()
thread2.start()
thread1.join()
#thread2.join()
print('main process exit')	
exitFlag = 1
exit()
