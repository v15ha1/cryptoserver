#!/bin/sh

# Service start/stop script
# Usage : service.sh [start|stop]
#set -x
RED='\033[0;31m'
NC='\033[0m'
BIN_NAME=cryptoserver-app

function _log()
{
    msg=$1
    printf "${RED}${msg}${NC}\n"
}

cd "$(dirname "$0")"/..

if [ "$DEPLOY_ENV" != "dev" ] && [ "$DEPLOY_ENV" != "cert" ] && [ "$DEPLOY_ENV"                                                                                                              != "prod" ]; then
  _log "Error: DEPLOY_ENV must be set to one of [dev, cert, prod]"
  exit
fi

if [ -z "${CONFIG_FILE}" ]; then
   export CONFIG_FILE="cfg/${DEPLOY_ENV}/config.yaml"
fi

if [ ! -f "${CONFIG_FILE}" ]; then
   _log "Error: Cannot access configuration file: ${CONFIG_FILE}"
   exit
fi

function _start()
{
    _log "Starting service.."
    bin/${BIN_NAME} --conf $CONFIG_FILE
}

if [ "$1" == "start" ]; then
    _start
else
    _log "Usage: service.sh start"
    exit 1
fi

