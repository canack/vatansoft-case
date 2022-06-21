FROM mysql

COPY ./database/db.sql /docker-entrypoint-initdb.d/

CMD [ "mysqld", "--init-file=/docker-entrypoint-initdb.d/db.sql" ]
