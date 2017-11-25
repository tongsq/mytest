from collections import deque
queue = deque(['aa','bb','cc'])
queue.append('dd')
print(queue.popleft())
print(queue)

vec = [2,4,6]
vec3 = [[x,x*3] for x in vec if x>2]
print(vec3)
question = ['aa?','bb?']
answer = ['a', 'b']
for que,ans in zip(question, answer):
	print('question:{0} answer:{1}'.format(que,ans))

knights = {'gallahad': 'the pure', 'robin': 'the brave'}
for k, v in knights.items():
	print(k, v)
for i, v in enumerate(['tic', 'tac', 'toe']):
	print(i, v)
