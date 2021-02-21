FROM node:15.9-alpine

WORKDIR /app
COPY ./web /app

RUN apk --no-cache update \
    && apk add curl

CMD sh ./entry-point.sh