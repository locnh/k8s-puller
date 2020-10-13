#!/usr/bin/env python3

import schedule
import time
import os

images   = os.getenv('IMAGES').split(',')
interval = int(os.getenv('INTERVAL', 60))

def pull(run_once):
    print("Start pulling ...")
    try:
        for image in images:
            print("Pulling image %s" % image)
            os.system("docker pull %s" % image)
            print()
    except:
        import traceback
        print(traceback.format_exc())
        return schedule.CancelJob

    if run_once == True:
        return schedule.CancelJob

schedule.every().second.do(pull, True)
schedule.every(interval).minutes.do(pull, False)

while True:
    schedule.run_pending()
    time.sleep(1)
