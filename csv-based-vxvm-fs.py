#!/bin/usr/python

import csv
import logging
import subprocess
import sys
import time

dtnow = time.strftime("%Y-%m-%d_%H_%M_%S")
fname = 'xyz.logs_'+dtnow
logging.basicConfig(filename=fname, filemode='w', level=logging.DEBUG, format='%(levelname)s %(asctime)s:%(message)s')

def file_validator(file, header):
    if not file.endswith(.csv):
        print "File format mis-match"
        return False
    with open(file, 'r') as f:
        reader = csv.reader(f)
        head = reader.next()  # Import only Header
        if (head == header):
            return True
        else:
            print ('Header mis-match')
            return False

def file_processing(csvfile):
    with open(csvfile, 'r') as csv_file:
        csv_reader = csv.DictReader(csv_file)
        for row in csv_reader:
            csv_reader = csv.DictReader(csv_file)
            # print rows

            DG = row['DiskGroup']
            Vol = row['Volume']
            IN = row['FSSIZE']
            disks = row['Disks']
            DI = row['DiskIN']
            VX = row['Filesystem']

            # initialize Disks

            if not DI:
                break
            initializeDG = '/etc/vx/bin/vxdisksetup -i' + ' ' + DI
            # print initializeDG
            with open('InitializeDG.txt', 'a') as fw:
                fw.write(InitializeDG)
                fw.write('\n')

            # print CreateDG
            if not DG:
                break
            if len(disks.strip()) > 0:
                CreateDG = 'vxdg init ' + ' ' + DG + ' ' + disks
                with open('CreateDG.txt', 'a') as fw:
                    fw.write(CreateDG)
                    fw.write('\n')

            # Create Volumes
            if not DG:
                break
            CreateVol = 'vxassist -g ' + DG + ' ' + 'make' + ' ' + Vol + ' ' +''+ IN + 'G'
            # print CreateVol
            with open('CreateVol.txt', 'a') as fw:
                fw.write(CreateVol)
                fw.write('\n')

            # Create FS
            if not DG:
                break
            VisibleFS = 'mkfs.vxfs' + '/dev/vx/rdsk' + DG + '/' + Vol
            # print VisibleFS
            with open('VisibleFS.txt', 'a') as fw:
                fw.write(VisibleFS)
                fw.write('\n')

            # Entry to fstab
            fstab = '/dev/vx/dsk' + DG + '/' + Vol
            Putfstab = fstab + ' ' + VX + ' ' + 'vxfs  _netdev    0  0'

            if not DG:
                break
            else
                with open('Putfstab.txt', 'a') as fw:
                    fw.write(Putfstab)
                    fw.write('\n')

    csv_file.close()

    ### Command execution starts here ###

    def initializing_disk_groups(file):
        with open(file, 'r') as f:
            lines = f.readlines()
            logging.info("Initializing the Disks")

            for cmd in lines:
                ID = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
                out, err = ID.communicate()
                if ID.returncode == 0
                    logging.info(cmd.strip("\n\r") + " :: Disk is initialized")
                if ID.returncode != 0:
                    logging.error(cmd)
                    logging.error(err + "DG init failed")
                    sys.exit('Exiting disk init, check logs')
        logging.into("Disks intitialized and added to DG")
        print ('Disks are initialized and Added to DGs')

    def creating_diskgroups(file):



    def creating_volumes(file):



    def formatting_filesystems(file):



    def fstab():


    def directories(file):



    def mount_fs():


    def health_check():
