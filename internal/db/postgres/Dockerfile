FROM postgres:12.16-alpine
# Copiar el archivo de creación de la base de datos
COPY up.sql /docker-entrypoint-initdb.d/

# Copiar el archivo de creación de la base de datos
# COPY up.sql /docker-entrypoint-initdb.d/
# Variables de entorno para la configuración
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=leal-challenge-database

# Puerto expuesto por el contenedor
EXPOSE 5432

# Comando de inicio de la base de datos
CMD ["postgres"]