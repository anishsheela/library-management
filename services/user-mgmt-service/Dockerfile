FROM python:3.11-slim

# Install system dependencies required for building mysqlclient
RUN apt-get update && apt-get install -y \
    build-essential \
    libmariadb-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /code

COPY requirements.txt /code/
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . /code/

COPY . .

# Copy the JWKS file into the container
COPY jwks.json /app/jwks.json