#!/usr/bin/python3
import os, time, socket
import logging, logging.config
from urllib.request import urlopen
import json, daemon
import subprocess

def getHostInfo():
     try:
       host_name = socket.gethostname()
       host_ip = socket.gethostbyname(host_name)
       return host_name
     except:
       print("Unable to get host name")

host=getHostInfo()
serverUrl="http://192.168.10.10/myagent"

#This url is for pulling the list of commands from host
cmdurl=("%s/getcmds.php?hostname=%s" % (serverUrl,host))

#This url is for updating the server about finished job
seturl=("%s/setcmds.php?hostname=%s" % (serverUrl,host))

formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
logger = logging.getLogger('myagent')
logger.setLevel(logging.INFO)

fh = logging.FileHandler("/home/agent/myagent.log")
fh.setFormatter(formatter)
logger.addHandler(fh)
logger.info('Starting the agent ...')

def main_program():
    logger.info('Agent started ...')
    while True:
       response = urlopen(cmdurl)
       data = json.loads(response.read())
       for x in data:
              seq=x['cmdseq']
              cmd=x['cmd']
              cmdstr='sudo %s' % (cmd)
              r = subprocess.run(cmdstr.split(),stdout=subprocess.PIPE,stderr=subprocess.PIPE)
              p = r.returncode
              err=r.stderr
              our=r.stdout
              #p=os.system(cmdstr)
              if (p==0):
                    urlopen(seturl+"&cmdseq="+seq)
                    logger.info("Command successfully executed"+cmd)
                    logger.info(out)
                    logger.info(err)
              else:
                    logger.error("Error executing the command"+cmd)
       time.sleep(1)

with daemon.DaemonContext(file_preserve = [ fh.stream ]):
main_program()