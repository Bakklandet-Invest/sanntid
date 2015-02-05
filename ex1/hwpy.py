from threading import Thread

i = 0

def someThreadFunction1():
#    i = i + 1

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
    global i
    for n in range(0,1000000):
    	i = i + 1

def someThreadFunction2():
#    i = i - 1

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
    global i

    for a in range(0,1000000):
        i = i - 1


def main():
    someThread1 = Thread(target = someThreadFunction1, args = (),)
    someThread1.start()
    someThread2 = Thread(target = someThreadFunction2, args = (),)
    someThread2.start()
    
    someThread1.join()    
    someThread2.join()

    print(i)


main()
