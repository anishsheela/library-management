FROM python:3.11-slim

# Install system dependencies for mysqlclient
RUN apt-get update && apt-get install -y \
    build-essential \
    libmariadb-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /code

COPY requirements.txt /code/
RUN pip install --upgrade pip && pip install --no-cache-dir -r requirements.txt

COPY . /code/

# Copy the JWKS file if needed in a specific location
COPY jwks.json /code/jwks.json

EXPOSE 8000

CMD ["python", "-m", "django", "runserver", "0.0.0.0:8000"]