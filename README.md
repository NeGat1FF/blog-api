# Blog API

This is a simple Blog API built using Golang.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Get all posts](#get-all-posts)
  - [Add a post](#add-a-post)
  - [Get a post by ID](#get-a-post-by-id)
  - [Update a post by ID](#update-a-post-by-id)
  - [Delete a post by ID](#delete-a-post-by-id)
- [Models](#models)
- [License](#license)

## Overview

This API allows you to create, read, update, and delete blog posts. You can also filter posts based on a search term.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/NeGat1FF/blog-api.git
   ```
2. Navigate to the project directory:
   ```bash
   cd blog-api
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

The API will run on `http://localhost:8080`.

## Usage

You can interact with the API using any API testing tool (e.g., Postman, curl). The base URL for all endpoints is `http://localhost:8080`.

## API Endpoints

### Get all posts

- **Endpoint:** `GET /posts`
- **Description:** Retrieve all posts. Optionally, you can search for posts by a specific term.
- **Parameters:**
  - `term` (optional, query string): The search term to filter posts.
- **Response:**
  - `200 OK`: Returns an array of posts.
  - `500 Internal Server Error`: If something went wrong.

### Add a post

- **Endpoint:** `POST /posts`
- **Description:** Create a new blog post.
- **Request Body:**
  ```json
  {
    "title": "string",
    "content": "string",
    "category": "string",
    "tags": ["string"]
  }
  ```
- **Response:**
  - `200 OK`: Returns the ID of the created post.
  - `400 Bad Request`: If the request body is invalid.
  - `500 Internal Server Error`: If something went wrong.

### Get a post by ID

- **Endpoint:** `GET /posts/{id}`
- **Description:** Retrieve a specific post by its ID.
- **Parameters:**
  - `id` (required, path): The ID of the post to retrieve.
- **Response:**
  - `200 OK`: Returns the post data.
  - `400 Bad Request`: If the post ID is invalid.
  - `500 Internal Server Error`: If something went wrong.

### Update a post by ID

- **Endpoint:** `PUT /posts/{id}`
- **Description:** Update a specific post by its ID.
- **Parameters:**
  - `id` (required, path): The ID of the post to update.
- **Request Body:**
  ```json
  {
    "title": "string",
    "content": "string",
    "category": "string",
    "tags": ["string"]
  }
  ```
- **Response:**
  - `200 OK`: Returns a success message.
  - `400 Bad Request`: If the request body or post ID is invalid.
  - `500 Internal Server Error`: If something went wrong.

### Delete a post by ID

- **Endpoint:** `DELETE /posts/{id}`
- **Description:** Delete a specific post by its ID.
- **Parameters:**
  - `id` (required, path): The ID of the post to delete.
- **Response:**
  - `200 OK`: Returns a success message.
  - `400 Bad Request`: If the post ID is invalid.
  - `500 Internal Server Error`: If something went wrong.

## Models

### Post

```json
{
  "id": 1,
  "title": "string",
  "content": "string",
  "category": "string",
  "tags": ["string"],
  "created_at": "string",
  "updated_at": "string"
}
```

### AddPostRequest

```json
{
  "title": "string",
  "content": "string",
  "category": "string",
  "tags": ["string"]
}
```

### UpdatePostRequest

```json
{
  "title": "string",
  "content": "string",
  "category": "string",
  "tags": ["string"]
}
```
https://roadmap.sh/projects/blogging-platform-api

