# Go Hexagonal API

This project is a study case to configure an hexagonal (clean architecture) with go.

To exercise this project we are going to implement a classic TODO app.


## TODO API app

We are creating a TODO API app that can:

- INSERT a new TODO item
- LIST my TODO items
- DELETE one TODO item


## Execute

```bash
docker build -t go-hex-api . && docker run -it -p 8080:8080 go-hex-api
```


## References

-  https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a
