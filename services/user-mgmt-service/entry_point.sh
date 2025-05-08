#!/bin/bash
set -e

# Run DB migrations
python manage.py makemigrations --noinput
python manage.py migrate --noinput

# Start the Django dev server (change this to gunicorn for production)
python manage.py runserver 0.0.0.0:8000
