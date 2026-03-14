FROM postgres:18.3-bookworm

COPY init.sql /docker-entrypoint-initdb.d/init.sql

USER postgres

CMD ["postgres"]
