# HTTP API Deployment Exercise

[![Build, Test, Push, and Deploy Docker Image](https://github.com/agavrushenko/containerized-api-cicd/actions/workflows/docker-cicd.yml/badge.svg?branch=main)](https://github.com/agavrushenko/containerized-api-cicd/actions/workflows/docker-cicd.yml)

## Overview

This repository contains a simple web application built with the Fiber framework using Golang. The application exposes an HTTP API endpoint that returns a JSON object with a message and a dynamically generated timestamp. Additionally, it includes Docker containerization and GitHub Workflow for CI/CD.

## API Documentation   
   
#### `GET /`
Returns a message and a timestamp.

#### Example Response   
```JSON
{"message":"My name is Alexey Gavrushenko","timestamp":1791410525578}
```

## Usage

### Golang Application

1. Install Golang: [https://go.dev/dl/](https://go.dev/dl/)
2. Clone this repository:
   ```sh
   git clone https://github.com/agavrushenko/containerized-api-cicd.git
   cd containerized-api-cicd
   ```
 3. Install Fiber:
   ```sh
   go get github.com/gofiber/fiber/v3
   ```
 4. Run the application:
   ```sh
   go run server.go
   ```

### Run Docker from Docker Image repository

 1. Pull the Docker image:  
   ```sh
   docker pull ghcr.io/agavrushenko/containerized-api-cicd:latest
   ```
 3. Run the Docker container:  
   ```sh
   docker run -p 80:8080 --name api-endpoint ghcr.io/agavrushenko/containerized-api-cicd:latest
   ```
