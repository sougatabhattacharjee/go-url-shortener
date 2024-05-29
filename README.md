# URL Shortener Service

A simple URL shortener service built with Go and PostgreSQL. This service allows users to shorten URLs with a choice of predefined domains and provides analytics on the shortened URLs.

## Table of Contents

- [Features](#features)
- [Setup Instructions](#setup-instructions)
    - [Prerequisites](#prerequisites)
    - [Environment Variables](#environment-variables)
    - [Running the Application](#running-the-application)
- [Endpoints](#endpoints)
    - [Shorten URL](#shorten-url)
    - [Redirect to Long URL](#redirect-to-long-url)
    - [Get URL Details](#get-url-details)
    - [Get URL Analytics](#get-url-analytics)

## Features

- Shorten URLs with a choice of predefined domains
- Redirect to the original URL using the shortened URL
- Get details and analytics about the shortened URLs

## Setup Instructions

### Prerequisites

- Docker
- Docker Compose

### Environment Variables

Create a `.env` file in the root directory of the project with the following content:

```env
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=urlshortener
DATABASE_URL=postgres://user:password@db:5432/urlshortener?sslmode=disable
PORT=8080
SHORT_URL_DOMAINS=t.ly,ibit.ly
```

### Running the Application

1. **Rebuild and Start Services**:
   ```sh
   docker-compose up --build
   ```

2. **Check Logs**:
   Ensure the services are running correctly by checking the logs.
   ```sh
   docker-compose logs
   ```

## Endpoints

### Shorten URL

#### Request:
```sh
curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://example.com", "domain": "t.ly"}' http://localhost:8080/api/v1/shorten
```

#### Response:
```json
{
  "short_url": "t.ly/<shortened_code>"
}
```

### Redirect to Long URL

#### Request:
```sh
curl -L http://localhost:8080/<shortened_code>
```

#### Example:
```sh
curl -L http://localhost:8080/abc123
```

### Get URL Details

#### Request:
```sh
curl http://localhost:8080/api/v1/urls/<shortened_code>
```

#### Example:
```sh
curl http://localhost:8080/api/v1/urls/abc123
```

#### Response:
```json
{
  "short_url": "abc123",
  "long_url": "https://example.com",
  "created_at": "2024-05-28T12:34:56Z"
}
```

### Get URL Analytics

#### Request:
```sh
curl http://localhost:8080/api/v1/urls/<shortened_code>/analytics
```

#### Example:
```sh
curl http://localhost:8080/api/v1/urls/abc123/analytics
```

#### Response:
```json
{
  "short_url": "abc123",
  "long_url": "https://example.com",
  "clicks": 42,
  "created_at": "2024-05-28T12:34:56Z",
  "updated_at": "2024-05-28T14:34:56Z"
}
```