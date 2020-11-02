import logging
import sys
formatmsg='%(asctime)s:%(levelname)s:%(filename)s:%(lineno)s:%(message)s'
logging.basicConfig(format=formatmsg, filename='mvp.log', level=logging.DEBUG)

# logger and stramHanler
logger = logging.getLogger('mvp')
ch = logging.StreamHandler(sys.stdout)
ch.setLevel(logging.DEBUG)
ch.setFormatter(logging.Formatter(formatmsg))
logger.addHandler(ch)

logger.debug('error message')
logger.debug({'err':'error message'})
