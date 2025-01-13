# 🌟 GRPC Hello World with Go

## 📘 About The Project
This is a simple Go application that demonstrates how to implement a gRPC client-server architecture. It includes instructions for generating gRPC code and running both the server and client.

## Ensure you have the following installed:

Ensure you have the following installed:

Go (at least version 1.23.2):
```bash
go version
```
Protocol Buffers Compiler (protoc):
Download and install from here.

## 📥 Installation
1.- Clone the repository
   ```sh
   git clone https://github.com/Karina1014/Grpc_golang_Hello_World.git
  ```
2.- Initialize the module:
   ```sh
   go mod init grpcGo/architectural-styles/grpc
  ```
3.- Install necessary dependencies:
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
- Add gRPC to your project
 ```sh
   go get google.golang.org/grpc
  ```
4.- Tidy up the dependencies:
```sh
   go mod tidy
  ```
## Running
Run the server: Start the server by running:
```sh
   go run server.go
  ```
  **The server will start and listen for gRPC requests.**

2.-Run the client:
Open a new terminal and run the client:
```sh
   go run client.go
  ```
**The client will send requests to the server, and you'll see the responses in the terminal.**

##Result:

![image](https://github.com/user-attachments/assets/65492a8f-2356-4777-a158-5c8b7e264b9b)

🎉 That's it! Your gRPC application is now up and running.


