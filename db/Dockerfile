FROM postgres:latest

COPY migrations migrations

RUN mkdir -p /var/lib/postgresql/data

ADD create_tables.sql /docker-entrypoint-initdb.d/

CMD ["postgres", "-D", "/var/lib/postgresql/data"]