#!/bin/sh
### BEGIN INIT INFO
# Provides:          server_remote
# Required-Start:    $network $syslog
# Required-Stop:     $network $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start go server at boot time
# Description:       your cool web app
### END INIT INFO

case "$1" in
  start)
    exec /home/go/src/remote-port/server/server &
    ;;
  stop)
    pkill -x server
    ;;
  *)
    echo $"Usage: $0 {start|stop}"
    exit 1
esac
exit 0