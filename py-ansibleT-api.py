#!/bin/usr/python

#
# This is basic code just started in search of python app that can communicate with AnsibleTower
# One must have python installed with towerlib module as well. Here pip or easy_install should help.
# pip install towerlib --user   --->  one can try anaconda prompt and Spyder as per choice
#

from towerlib import Tower

#to connect with ansible tower using valid authentication
tower = Tower('xyz-ansibleT.com', 'Username', 'password', secure=True, ssl_verify=False)

#to print list of servers in the inventory
for host in tower.hosts:
    print(host.name)
#
