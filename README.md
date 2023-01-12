# the-better-backend
A template for a GoLang backend using Fiber, MongoDB, a Makefile, and more

## Getting Started

### VIDEOS

- [GoLang Download & Setup](https://www.youtube.com/watch?v=Q7uh85_i0-M)
- [Video Breakdown](https://youtu.be/6C-2R92L01Q)

### Prerequisites

- [GoLang](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

### Installing

0. Install extra packages: 
    ```go install github.com/cosmtrek/air@latest```
    ```go install github.com/swaggo/swag/cmd/swag@latest```
1. Clone the repo
2. Create your own .env file
3. ```make dev```
4. view docs at http://localhost:8080/swagger

### Scripts

- ```make dev``` - runs the server in development mode
- ```make swagger``` - generates the swagger docs