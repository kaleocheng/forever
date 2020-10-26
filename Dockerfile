FROM alpine as builder
RUN apk --no-cache add gcc libc-dev
COPY c/main.c  /main.c
RUN gcc -o /forever /main.c && chmod +x /forever

FROM alpine
COPY --from=builder /forever /usr/bin/forever
CMD /usr/bin/forever
