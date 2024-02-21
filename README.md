# Web Crawler

This is a web crawler written in golang.

### 📋 Prerequirements

* [Golang-version-1.21.2] (https://go.dev/doc/install)
* [Code-editor] (https://code.visualstudio.com/download)
* [Docker] 

### 🔧 How to run the project?

After installed the tools on Prerequirements section and cloned the project to your local machine.

* First of all, in your terminal, certify you are in cmd directory and run the following command on your terminal:

```
go run main.go
```

Once everything is fine, in order to run MongoDB, certify you have docker installed in your machine, after it type on your terminal

```
docker compose up -d
```

You'll be uploading an instance of MongoDB Container on your local machine.

I've already let prepared a command in makefile to exec your container, in order to access mongoDB, so you can type in terminal:

```
make dockerexec
```

Once you're inside of your container, you can type "mongosh" to run mongo instance and see the database of your application

## ⚙️ Tests

In order to run tests locally, if you are using VSCode, type on your terminal

```
go test -v
```

## 🛠️ Tools used in this project

[Docker] 
[MongoDB] 
