# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
import threading


global i
i = 0

lock = threading.Lock()

def someThreadFunction1():
	global i
	lock.acquire() 	
	for n in range(0,1000000):
		i = i + 1
	lock.release()

def someThreadFunction2():
	global i
	lock.acquire()	
	for m in range(0,999999):
		i = i - 1
	lock.release()


# Potentially useful thing:
# In Python you "import" a global variable, instead of "export"ing it when you declare it
# (This is probably an effort to make you feel bad about typing the word "global")



def main():
	someThread1 = Thread(target = someThreadFunction1, args = (),)
	someThread1.start()
	someThread2 = Thread(target = someThreadFunction2, args = (),)
	someThread2.start()
	

	someThread1.join()
	someThread2.join()
	
	print "i = ",i


main()
