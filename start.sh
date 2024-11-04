#!/bin/sh

# Menjalankan migrasi
echo "Running migrations..."
/app/migrate -database "$DB_DRIVER://$DB_SOURCE" -path /app/migration -verbose up


# Memulai aplikasi
echo "Starting the application..."
exec "$@"