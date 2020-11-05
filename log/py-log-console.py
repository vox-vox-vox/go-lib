import logging
import sys
import json

def getLogger(name='root', level=logging.DEBUG):
    formatmsg='%(asctime)s\t%(levelname)s\t%(name)s\t%(pathname)s:%(lineno)s\t%(message)s'
    datefmt='%Y-%m-%dT%I:%M:%S%z'
    logging.basicConfig(format=formatmsg, datefmt=datefmt, level=level)

    logger = logging.getLogger(name)
    #ch = logging.StreamHandler(sys.stdout)
    #ch.setLevel(level)
    #ch.setFormatter(logging.Formatter(formatmsg, ), )
    #logger.addHandler(ch)
    return logger


logger = getLogger('translog') # 实例名
log = {
    # 日志类型 1.数据格式错误　DATAERROR 2.网络服务错误    NETWORKERROR 3.计算错误             COMPUTEERROR
    'type': "DATAERROR", 
    # 日志消息
    'msg':'M1903-SKU4-V03-2006-0064_2020_08_27_13_29_11 　starts preparing file',

    # 其它：自定义
    "pod_name":'detection-partition-64c469f6c8-q9zlb',
    'module_name':'translog-A[73]',
}

logger.info(json.dumps(log))
#------------------------------
#2020-11-05T08:11:51+0800	INFO	translog	py-log-console.py:30	{"type": "DATAERROR", "msg": "M1903-SKU4-V03-2006-0064_2020_08_27_13_29_11 \u3000starts preparing file", "pod_name": "detection-partition-64c469f6c8-q9zlb", "module_name": "translog-A[73]"}
