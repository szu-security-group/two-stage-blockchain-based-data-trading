# -*- coding: utf-8 -*-
from functools import reduce

def count_challenge(n):
 #upper bound for challenge lentgh
    count = int(4 * n ** 0.5)
    blocks_broken = int (n * 0.05)
    temp = [(n - blocks_broken - i) / (n - i) for i in range(count)]
    for i in range(count):
        prob_fail = reduce(lambda x, y: x * y, temp[0:i+1])
        if prob_fail <= 0.01:
            return [i, 1 - prob_fail]

#n: block size 1MB
#n100k: block size 100KB

#Total number of blocks in 1MB chunks
#Modify the n array to suit your needs
n = [161, 385, 1.14 * 1024, 2.41 * 1024, 5 * 1024, 10 * 1024, 15 * 1024]
len_n = len(n)
#Total number of blocks in 100KB chunks
n100k = [n[i] * 10.24 for i in range(len_n)]

size = [count_challenge(i) for i in n]
print ("challenge size is:" + str(size))

size_100k = [count_challenge(i) for i in n100k]
print ("challenge size_100k is:" + str(size_100k))
