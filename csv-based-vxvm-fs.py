#!/bin/usr/python

import csv
import logging
import subprocess
import sys
import time

dtnow = time.strftime("%Y-%m-%d_%H_%M_%S")
fname = 'xyz.logs_'+dtnow
logging.basicConfig(filename=fname, filemode='w', level=logging.DEBUG, format='%(levelname)s %(asctime)s:%(message)s')
