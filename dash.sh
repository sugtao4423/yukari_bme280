#!/bin/bash

DASH_IP='192.168.1.210'

while true
do
    ping -c 1 -w 100 -t 1 $DASH_IP > /dev/null 2>&1
    if [ $? -eq 0 ]; then
        php /home/tao/yukari_bme280/yukari_bme280.php
        sleep 5
    fi
done
