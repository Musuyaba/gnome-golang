FROM alpine:latest
WORKDIR /certs/

RUN apk add --no-cache openssl
COPY generate.sh .
COPY domains.ext .

CMD [ "/certs/generate.sh" ]