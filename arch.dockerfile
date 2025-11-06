# ╔═════════════════════════════════════════════════════╗
# ║                       SETUP                         ║
# ╚═════════════════════════════════════════════════════╝
# GLOBAL
  ARG APP_UID=1000 \
      APP_GID=1000

# FOREIGN IMAGES
  FROM 11notes/distroless:openssl AS distroless-openssl


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
    COPY --from=build /distroless/ /
    COPY --from=distroless-openssl / /
    COPY ./rootfs /

  # :: app specific environment
    ENV HISTFILE=/.ash_history

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
    chown -R ${APP_UID}:${APP_GID} ${APP_ROOT};

# :: MONITOR
  HEALTHCHECK --interval=5s --timeout=2s --start-period=5s \
    CMD ["/usr/bin/nc", "-z", "127.0.0.1", "22"]

# :: EXECUTE
  USER ${APP_UID}:${APP_GID}
  ENTRYPOINT ["/usr/local/bin/entrypoint"]