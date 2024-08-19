# Apartment Service API

![Build Status](https://github.com/PabloPerdolie/apartment-service/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/PabloPerdolie/apartment-service/graph/badge.svg?token=9FWV0I5PTC)](https://codecov.io/gh/PabloPerdolie/apartment-service)

## Overview

The **Apartment Service** is a Go-based application designed to manage apartments in various residential buildings. It provides an API for adding, updating, and retrieving apartment information within a specified house.

## Features

- **Apartment Management**: Create and retrieve apartment data.
- **House Management**: Manage houses and their associated apartments.
- **User Subscription**: Users can subscribe to receive notifications about new apartments.
- **Moderation**: Includes functionality to moderate apartment status and create house for moderators.

## Architecture

The project follows the principles of clean architecture, separating business logic, infrastructure and interfaces.
A service provider is used for dependency management and dependency injection.

## Technologies Used

- **Language**: Go 1.22 
- **Database**: PostgreSQL
- **Testing**: `Testify`, `sqlmock`
- **CI/CD**: GitHub Actions, Codecov
- **Deployment**: Docker & Docker compose

## Setup Instructions

Follow these steps to set up and run the Apartment Service API locally:

### Clone the repository

```bash
git clone https://github.com/PabloPerdolie/apartment-service.git
cd apartment-service
```
### Launching the application

```bash
make run
```

The service will start running on http://localhost:8080/.

#### To clean up:
```bash
make clean
```

_Thanks for checking out my service._
