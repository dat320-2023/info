import threading
from threading import Condition
import time
import random
queue = []
condition = Condition()

def produce():
    nums = range(5) #Will create the list [0, 1, 2, 3, 4]
    global queue
    MAX_NUM = 10
    
    while True:
        num = random.choice(nums) #choose a random number from nums
        condition.acquire()
        if len(queue) == MAX_NUM:
            print ("Full buffer")
            condition.wait()
            print ("Buffer ready, Consumer notified the producer")       
        queue.append(num)
        print("Produced", num, "queue size", len(queue))
        condition.notify()
        condition.release()
        time.sleep(1)

def consume():
    global queue
    while True:
        condition.acquire()
        if not queue:
            print ("Nothing in queue, Wait")
            condition.wait()
        num = queue.pop(0)
        print("Consumed", num)
        condition.notify()
        condition.release()
        time.sleep(10)
        
p=threading.Thread(target=produce)
c=threading.Thread(target=consume)
p.start()
c.start()
p.join()
c.join()
time.sleep(1)
print("Finished")