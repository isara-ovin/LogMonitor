import logging
import sys
import time
from datetime import datetime, timedelta

file_handler = logging.FileHandler(filename='stream.log')
stdout_handler = logging.StreamHandler(stream=sys.stdout)
handlers = [file_handler, stdout_handler]

logging.basicConfig(
                format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                datefmt='%Y-%m-%d %H:%M:%S',
                handlers=handlers,
                level=logging.DEBUG)


now = datetime.now()
lapsed = now + timedelta(seconds=5)

while datetime.now() < lapsed:
    i = 1
    logging.debug(f"Script running debug {i}")
    logging.info(f"Script running info {i}")
    logging.error(f"Script running error {i}")
    i = i + 1
    time.sleep(0.2)