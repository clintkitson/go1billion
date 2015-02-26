FROM scratch

ADD ./release/go1billion-Linux-static /bin/go1billion

ENV PORT 8080

EXPOSE 8080
ENTRYPOINT ["/bin/go1billion"]

CMD ["--help"]
