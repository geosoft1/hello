Test server for HTTP/HTTPS connections.
---

Generate self signed certificates:

	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt

Test HTTP connections:

	curl --verbose http://localhost:8080/hello

Test HTTPS connections:

	curl --verbose --insecure https://localhost:8090/hello
