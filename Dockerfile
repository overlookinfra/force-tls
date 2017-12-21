FROM scratch
COPY build/force-tls /force-tls
ENTRYPOINT ["/force-tls"]
