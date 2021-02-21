FROM golang:1.15-alpine
ENV LANG='en_US.utf8'
ENV GO111MODULE='off'

ARG SQL_DB
ARG SQL_USER
ARG SQL_PASSWORD
ARG SQL_PORT
ENV SQL_DB $SQL_DB
ENV SQL_USER $SQL_USER
ENV SQL_PASSWORD $SQL_PASSWORD
ENV SQL_PORT $SQL_PORT

WORKDIR /go/app
COPY ./api /go/app

RUN apk --no-cache update \
    && apk add curl

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air