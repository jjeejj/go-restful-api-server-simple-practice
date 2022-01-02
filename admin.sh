#!/bin/bash

SERVER="go-restful-api-server-simple-practice"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数
ARGS=""

function start()
{
    if [ "`pgrep -u $UID $SERVER`" != "" ]; then
        echo "$SERVER already running"
        exit 1
    fi
    nohup $BASE_DIR/$SERVER $ARGS server 1>/dev/null &
    echo "sleeping..." && sleep $INTERVAL

    # check status
    if [ "`pgrep -u $UID $SERVER`" == "" ]; then
        echo "$SERVER start fail"
        exit 1
    fi
    echo "start $SERVER server success"
}

function status()
{
    if [ "`pgrep -u $UID $SERVER`" != "" ]; then
        echo $SERVER is running
    else
        echo $SERVER is not running
    fi
}

function stop()
{
    if [ "`pgrep -u $UID $SERVER`" != "" ]; then
        kill -9 `pgrep -u $UID $SERVER`
    fi
    echo "sleeping..." && sleep $INTERVAL
    if [ "`pgrep -u $UID $SERVER`" != "" ]; then
        echo "$SERVER stop fail"
        exit 1
    fi
    echo "stop $SERVER server success"
}

case "$1" in 
    "start")
        start
        ;;
    "status")
        status
        ;;
    "stop")
        stop
        ;;
    "restart")
        stop && start
        ;;
     *)
        echo "usage: $0 {start|stop|status}"
        exit 1
esac