# Account API Client

A Go client library for interacting with a RESTful Account API, built on top of the [gohttp](../client-library/README.md) client library.

## Overview

This module provides a typed Go client for managing account resources via a RESTful API. It is designed to be used as a library within other Go projects.

## Operations

| Operation | Description                    |
|-----------|--------------------------------|
| Create    | Create a new account resource  |
| Fetch     | Retrieve an account by ID      |
| Delete    | Remove an account resource     |

## Data Model

Account data is represented by the `AccountData` struct (see [models.go](./models.go)), which includes:

- **AccountAttributes** - classification, account number, bank ID, BIC, currency, country, IBAN, status, and more
- **OrganisationID** - the owning organisation
- **Version** - optimistic locking version

## Running the API

The API and its dependencies (PostgreSQL, Vault) are defined in `docker-compose.yml`:

```bash
docker-compose up
```

The API will be available at `http://localhost:8080`. No authentication is required.

## Testing

Tests are designed to run against the dockerized API:

```bash
docker-compose up -d
go test ./...
```
