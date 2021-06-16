FROM alpine
LABEL cryptoserver-clean="1_0"

RUN addgroup svc_g_ms && mkdir /appl && adduser svc_u_app --shell /bin/sh -G svc_g_ms -S -h /appl/cryptoserver-clean

COPY --chown=svc_u_app:svc_g_ms run/ /appl/cryptoserver-clean/

RUN find / -perm /6000 -type f -exec chmod a-s {} \; || true

USER svc_u_app

#HEALTHCHECK --interval=60s --timeout=30s --start-period=40s  CMD /appl/cryptoserver-clean/bin/docker-healthcheck.sh

CMD ["/appl/cryptoserver-clean/bin/service.sh", "start"]
