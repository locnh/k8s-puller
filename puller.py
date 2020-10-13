#!/usr/bin/env python3

import schedule
import time
import os
from os import system

images   = os.getenv('IMAGES').split(',')
interval = int(os.getenv('INTERVAL', 60))

def pull():
    print("Pulling images ...")
    for image in images:
        system("""/usr/local/bin/docker pull %s""" % image)

schedule.every(interval).minutes.do(pull)

while True:
    schedule.run_pending()
    time.sleep(1)
