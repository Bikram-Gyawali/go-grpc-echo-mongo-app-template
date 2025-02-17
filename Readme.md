# 🚀 GO GRPC ECHO API TEMPLATE

## Overview
This project provides a **robust and scalable API development setup** using the **Echo framework**, one of the fastest Go frameworks. It leverages **gRPC and Protocol Buffers** for efficient inter-service communication while maintaining HTTP compatibility. Additionally, it includes a working **MongoDB integration**, ensuring seamless data management.

## 🔥 Features
- **Modular Service Implementations** – Clean and structured code organization.
- **gRPC & Protocol Buffers** – High-performance, contract-based communication.
- **MongoDB Integration** – Utilizes MongoDB official Go driver.
- **Dockerized Architecture** – Easily deployable using Docker Compose.
- **Auto-generated Swagger Documentation** – All services generate Swagger JSON from `.proto` files.
- **Swagger UI Integration** – Access and explore API documentation through a web-based interface.

## 📦 Project Structure
```
📂 go-grpc-echo-mongo-app-template
 .
├── Dockerfile
├── Dockerfile.gateway
├── Makefile
├── Readme.md
├── api
│   ├── pb
│   │   ├── book.pb.go
│   │   ├── book.pb.gw.go
│   │   ├── book_grpc.pb.go
│   │   ├── user.pb.go
│   │   ├── user.pb.gw.go
│   │   └── user_grpc.pb.go
│   └── proto
│       ├── book.proto
│       └── user.proto
├── cmd
│   └── gateway
│       └── main.go
├── config
│   └── config.go
├── db
│   └── db.go
├── docker-compose.yml
├── docs
│   ├── api
│   │   ├── book.swagger.json
│   │   └── user.swagger.json
│   ├── merged_swagger.go
│   └── swagger-ui
│       └── index.html
├── go.mod
├── go.sum
├── internal
│   ├── repository
│   │   ├── book_repository.go
│   │   ├── init.go
│   │   └── user_repository.go
│   ├── rpcs
│   │   ├── init.go
│   │   └── user_grpc_service.go
│   └── services
│       ├── book_service.go
│       ├── init.go
│       └── user_service.go
├── middleware
│   ├── auth.go
│   └── logger.go
└── tests
    └── service_test.go

```

## 🚀 Getting Started
### 1️⃣ Clone the Repository
```sh
git clone git@github.com:Bikram-Gyawali/go-grpc-echo-mongo-app-template.git
cd go-grpc-echo-mongo-app-template
```

### 2️⃣ Run the Project
Using Docker Compose:
```sh
docker-compose up -d --build
```

This command will:
- Build and start the API service.
- Set up the MongoDB instance.
- Expose the necessary ports for API access.

### 3️⃣ Access the API Documentation
Once the project is running, access Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

## 🛠️ Environment Variables
Configure the `.env` file with necessary values:
```ini
MONGO_URI=mongodb://localhost:27017
GRPC_PORT=50051
HTTP_PORT=8080
```

## 🏗️ Extending the Project
### Adding a New gRPC Service
1. Define a new `.proto` file inside the `proto` folder.
2. Run the following command to generate Go code:
   ```
   make generate PROTO_FILE=example.proto
   ```
3. Implement the repository, services and rpcs
4. Register all service,repository and services in their own corresponding init.go file

## 🤝 Contributing
Feel free to submit pull requests and raise issues to enhance this project.

## 📜 License
This project is **open-source** and licensed under the MIT License.

---

💡 **Tip:** Follow best practices when implementing new services and ensure that gRPC contracts remain backward compatible!

## 🤝 TODO: 
 [ ] - Add zap logger or similar
 [ ] - Add tests cases for services
 [ ] - Error handling and definitions
 [ ] - Proper structured success and error response format
 [ ] - 