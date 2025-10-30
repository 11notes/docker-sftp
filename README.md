![banner](https://github.com/11notes/defaults/blob/main/static/img/banner.png?raw=true)

# SFTP
![size](https://img.shields.io/docker/image-size/11notes/sftp/10.2?color=0eb305)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)![pulls](https://img.shields.io/docker/pulls/11notes/sftp?color=2b75d6)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)[<img src="https://img.shields.io/github/issues/11notes/docker-sftp?color=7842f5">](https://github.com/11notes/docker-sftp/issues)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)![swiss_made](https://img.shields.io/badge/Swiss_Made-FFFFFF?labelColor=FF0000&logo=data:image/svg%2bxml;base64,PHN2ZyB2ZXJzaW9uPSIxIiB3aWR0aD0iNTEyIiBoZWlnaHQ9IjUxMiIgdmlld0JveD0iMCAwIDMyIDMyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxyZWN0IHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgZmlsbD0idHJhbnNwYXJlbnQiLz4KICA8cGF0aCBkPSJtMTMgNmg2djdoN3Y2aC03djdoLTZ2LTdoLTd2LTZoN3oiIGZpbGw9IiNmZmYiLz4KPC9zdmc+)

Run sftp rootless.

# INTRODUCTION 📢

[OpenSSH](https://www.openssh.org/) (created by [OpenBSD](https://www.openbsd.org/)) is the premier connectivity tool for remote login with the SSH protocol. It encrypts all traffic to eliminate eavesdropping, connection hijacking, and other attacks. In addition, OpenSSH provides a large suite of secure tunneling capabilities, several authentication methods, and sophisticated configuration options.

# SYNOPSIS 📖
**What can I do with this?** This image will run a [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) sftp server based on openssh. Unlike other images, this image does not provide chroot jails, but it's intended to run as a single user with all the shares needed mounted in ```/home/%u```. Use an SSH proxy in front of this image when you need to expose multiple endpoints via a single entry point. The image’s entrypoint does quickly elevate to root to change the username and drops instantly back down to 1000 for the ssh server. If no **ssh_host_key** secret is set, the image will create a new one on each start of the container.

# UNIQUE VALUE PROPOSITION 💶
**Why should I run this image and not the other image(s) that already exist?** Good question! Because ...

> [!IMPORTANT]
>* ... this image runs [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) as 1000:1000
>* ... this image is auto updated to the latest version via CI/CD
>* ... this image has a health check
>* ... this image runs read-only
>* ... this image is automatically scanned for CVEs before and after publishing
>* ... this image is created via a secure and pinned CI/CD process
>* ... this image is very small

If you value security, simplicity and optimizations to the extreme, then this image might be for you.

# COMPOSE ✂️
Checkout [compose.secrets.yml](https://github.com/11notes/docker-sftp/blob/master/compose.secrets.yml) if you want to use secrets instead of environment variables.
```yaml
name: "sftp"

x-lockdown: &lockdown
  # prevents write access to the image itself
  read_only: true

services:
  sftp:
    image: "11notes/sftp:10.2"
    <<: *lockdown
    environment:
      TZ: "Europe/Zurich"
      SSH_USER: "foo"
      SSH_PASSWORD: "${SSH_PASSWORD}"
    volumes:
      - "foo.var:/home/foo"
    tmpfs:
      # needed for read-only
      - "/run/ssh:uid=1000,gid=1000"
      - "/run/secrets:uid=1000,gid=1000"
    ports:
      - "8022:22/tcp"
    networks:
      frontend:
    sysctls:
      # allow rootless container to access ports < 1024
      net.ipv4.ip_unprivileged_port_start: 22
    restart: "always"

volumes:
  foo.var:

networks:
  frontend:
```
To find out how you can change the default UID/GID of this container image, consult the [RTFM](https://github.com/11notes/RTFM/blob/main/linux/container/image/11notes/how-to.changeUIDGID.md#change-uidgid-the-correct-way).

# DEFAULT SETTINGS 🗃️
| Parameter | Value | Description |
| --- | --- | --- |
| `user` | docker | user name |
| `uid` | 1000 | [user identifier](https://en.wikipedia.org/wiki/User_identifier) |
| `gid` | 1000 | [group identifier](https://en.wikipedia.org/wiki/Group_identifier) |
| `home` | /home | home directory of user docker |

# ENVIRONMENT 📝
| Parameter | Value | Default |
| --- | --- | --- |
| `TZ` | [Time Zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | |
| `DEBUG` | Will activate debug option for container image and app (if available) | |
| `SSH_USER` | username to access SSH server |  |
| `SSH_PASSWORD` | pasword to access SSH server | |
| `SSH_PASSWORD_FILE` *(optional)* | pasword file to access SSH server | |

# MAIN TAGS 🏷️
These are the main tags for the image. There is also a tag for each commit and its shorthand sha256 value.

* [10.2](https://hub.docker.com/r/11notes/sftp/tags?name=10.2)

### There is no latest tag, what am I supposed to do about updates?
It is of my opinion that the ```:latest``` tag is dangerous. Many times, I’ve introduced **breaking** changes to my images. This would have messed up everything for some people. If you don’t want to change the tag to the latest [semver](https://semver.org/), simply use the short versions of [semver](https://semver.org/). Instead of using ```:10.2``` you can use ```:10```. Since on each new version these tags are updated to the latest version of the software, using them is identical to using ```:latest``` but at least fixed to a major or minor version.

If you still insist on having the bleeding edge release of this app, simply use the ```:rolling``` tag, but be warned! You will get the latest version of the app instantly, regardless of breaking changes or security issues or what so ever. You do this at your own risk!

# REGISTRIES ☁️
```
docker pull 11notes/sftp:10.2
docker pull ghcr.io/11notes/sftp:10.2
docker pull quay.io/11notes/sftp:10.2
```

# SOURCE 💾
* [11notes/sftp](https://github.com/11notes/docker-sftp)

# PARENT IMAGE 🏛️
* [${{ json_readme_parent_image }}](${{ json_readme_parent_url }})

# BUILT WITH 🧰
* [sftp/sftp](https://github.com/sftp/sftp)
* [11notes/util](https://github.com/11notes/docker-util)

# GENERAL TIPS 📌
> [!TIP]
>* Use a reverse proxy like Traefik, Nginx, HAproxy to terminate TLS and to protect your endpoints
>* Use Let’s Encrypt DNS-01 challenge to obtain valid SSL certificates for your services

# ElevenNotes™️
This image is provided to you at your own risk. Always make backups before updating an image to a different version. Check the [releases](https://github.com/11notes/docker-sftp/releases) for breaking changes. If you have any problems with using this image simply raise an [issue](https://github.com/11notes/docker-sftp/issues), thanks. If you have a question or inputs please create a new [discussion](https://github.com/11notes/docker-sftp/discussions) instead of an issue. You can find all my other repositories on [github](https://github.com/11notes?tab=repositories).

*created 30.10.2025, 01:04:54 (CET)*