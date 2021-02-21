FROM mysql:8.0
ENV LANG en_US.utf8

ARG SQL_DB
ARG SQL_USER
ARG SQL_PASSWORD
ENV MYSQL_DB $SQL_DB
ENV MYSQL_USER $SQL_USER
ENV MYSQL_ROOT_PASSWORD $SQL_PASSWORD

WORKDIR /docker-entrypoint-initdb.d

RUN apt-get update