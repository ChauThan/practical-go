# Practical Go

A hands-on workspace dedicated to mastering the Go programming language through practice, experimentation, and building real-world applications.

## Getting Go on Ubuntu

```bash
sudo rm -rf /usr/local/go
wget https://go.dev/dl/go1.25.7.linux-amd64.tar.gz -O go-latest.tar.gz
sudo tar -C /usr/local -xzf go-latest.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

Check [go.dev/dl](https://go.dev/dl/) for the latest version.

## Projects

**[basic](./basic)** - A collection of fundamental Go programming examples.

**[task-cli](./task-cli)** - A CLI task manager demonstrating Go fundamentals: structs, slices, error handling, and JSON persistence without traditional OOP concepts.

**[url-checker](./url-checker)** - A concurrent website checker and downloader demonstrating Go concurrency patterns: goroutines, channels, WaitGroups, mutex, and the worker pool pattern for bounded parallelism.

**[bookstore-api](./bookstore-api)** - A RESTful API for bookstore management demonstrating layered architecture (handlers → services → repositories), JWT authentication, GORM with PostgreSQL, and Go 1.22+ routing patterns.

**[crypto-monitor](./crypto-monitor)** - A real-time crypto price monitor demonstrating WebSocket streaming, Redis caching, alert pipelines, and context-driven shutdown.
