import threading
from threading import Condition, Lock
import time
import random
queue = []
condition = Condition()

def produce():
    nums = range(5) #Will create the list [0, 1, 2, 3, 4]
    global queue
    while True:
        num = random.choice(nums) #choose a random number from nums
        condition.acquire()
        queue.append(num)
        print("Produced", num, "queue size", len(queue))
        condition.notify()
        condition.release()
        time.sleep(10)

def consume():
    global queue
    while True:
        condition.acquire()
        if not queue:
            print ("Nothing in queue, Wait")
            condition.wait()
        num = queue.pop(0)
        print("Consumed", num)
        condition.release()
        time.sleep(1)
        
p=threading.Thread(target=produce)
c=threading.Thread(target=consume)
p.start()
c.start()
p.join()
c.join()
time.sleep(1)
print("Finished")