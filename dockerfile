# Use the official Postgres image
FROM postgres:latest

# Expose PostgreSQL default port
EXPOSE 5432

# Set default environment variables (replace values as needed)
ENV POSTGRES_PASSWORD=Test!Pass123
ENV POSTGRES_DB=exampledb
ENV POSTGRES_USER=exampleuser