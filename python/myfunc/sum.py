def sum(*nums):
	sum = 0
	for num in nums:
		sum = sum + num
	return sum
if __name__ == '__main__':
	print('sum(1,2,3) =', sum(1,2,3))
