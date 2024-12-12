# Steradian Cars Order API

## Overview
Steradian Cars is project to test as Backend Engineer in Steradian. with Golang and MySQL.
This project is a RESTful API for managing orders. It allows you to list, view, add, update, and delete cars and orders.

## Requirements
- Git
- Docker

## Setup
1. Clone the repository
2. Go to directory of project
3. Do command `make run`

## Endpoints Car
- `GET http://localhost:8080/steradian/api/v1/cars` - List all cars
- `GET http://localhost:8080/steradian/api/v1/car/{id}` - Get details of a specific car
- `POST /http://localhost:8080/steradian/api/v1/cars` - Add a new car
- `PATCH http://localhost:8080/steradian/api/v1/car/{id}` - Update a specific car
- `DELETE http://localhost:8080/steradian/api/v1/car/{id}` - Delete a specific car

## Endpoints Order
- `GET http://localhost:8080/steradian/api/v1/orders` - List all orders
- `GET http://localhost:8080/steradian/api/v1/order/{id}` - Get details of a specific order
- `POST /http://localhost:8080/steradian/api/v1/orders` - Add a new order
- `PATCH http://localhost:8080/steradian/api/v1/order/{id}` - Update a specific order
- `DELETE http://localhost:8080/steradian/api/v1/order/{id}` - Delete a specific order

## Database Connection
The application will attempt to connect to the database with a maximum of 5 retries and a 5-second delay between each retry(by default). This is to ensure the database service is up and running before the application attempts to connect. You can adjust by environment in `docker-compose.yaml`.