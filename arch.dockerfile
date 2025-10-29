# ╔═════════════════════════════════════════════════════╗
# ║                       SETUP                         ║
# ╚═════════════════════════════════════════════════════╝
# GLOBAL
  ARG APP_UID=1000 \
      APP_GID=1000


# ╔═════════════════════════════════════════════════════╗
# ║                       BUILD                         ║
# ╚═════════════════════════════════════════════════════╝
# :: ENTRYPOINT
  FROM 11notes/go:1.25 AS build
  COPY ./build /
  RUN set -ex; \
    cd /go/entrypoint; \
    eleven go build /entrypoint main.go; \
    eleven distroless /entrypoint;


# ╔═════════════════════════════════════════════════════╗
# ║                       IMAGE                         ║
# ╚═════════════════════════════════════════════════════╝
# :: HEADER
  FROM 11notes/alpine:stable

  # :: default arguments
    ARG TARGETPLATFORM \
        TARGETOS \
        TARGETARCH \
        TARGETVARIANT \
        APP_IMAGE \
        APP_NAME \
        APP_VERSION \
        APP_ROOT \
        APP_UID \
        APP_GID \
        APP_NO_CACHE

  # :: default environment
    ENV APP_IMAGE=${APP_IMAGE} \
        APP_NAME=${APP_NAME} \
        APP_VERSION=${APP_VERSION} \
        APP_ROOT=${APP_ROOT}

  # :: multi-stage
    COPY --chown=0:0 --from=build /distroless/ /
    COPY ./rootfs /

# :: INSTALL
  USER root
  RUN set -ex; \
    apk --update --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/edge/main add \
      rsync \
      openssh~=${APP_VERSION}; \
    rm -f /etc/motd; \
    mkdir -p /run/ssh; \
    touch /run/ssh/passwd; \
    rm -f /etc/passwd; \
    ln -s /run/ssh/passwd /etc/passwd; \
    chmod +x -R /usr/local/bin; \
    chmod 4755 /usr/local/bin/entrypoint;

# :: MONITOR
  HEALTHCHECK --interval=5s --timeout=2s --start-period=5s \
    CMD ["/usr/local/bin/entrypoint", "health"]

# :: EXECUTE
  USER ${APP_UID}:${APP_GID}
  ENTRYPOINT ["/usr/local/bin/entrypoint"]