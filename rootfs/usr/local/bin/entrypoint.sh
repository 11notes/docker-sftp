#!/bin/ash
  adduser -s /bin/ash -H -D foo && echo "foo:r8N0P3u7" | chpasswd
  /usr/bin/ssh-keygen -q -N "" -t ed25519 -f /run/ssh/ssh_host_ed25519_key
  set -- /usr/sbin/sshd -D -e

exec $@