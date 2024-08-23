# ./swagger generate server -f ./swagger.json

# ./swagger generate client -f ./swagger.json

# openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 -subj "/C=US/ST=State/L=Locality/O=Organization/CN=localhost"

ex00_server:
	go run ./ex00/cmd/candy-server-server/main.go --tls-certificate=./ex00/cert.pem --tls-key=./ex00/key.pem --port=3333

ex00_client:
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/v1/buy_candy


ex01_server:
	 go run ./ex01/server/cmd/candy-server-server/main.go --tls-certificate=./ex01/ca/minica.pem --tls-key=./ex01/ca/minica-key.pem --tls-port=3333 --port=3332

ex01_client:
	go build -o ex01/candy-client ex01/client/main.go
	cd ex01 && ./candy-client -k AA -c 2 -m 50

ex02_server:
	 go run ./ex02/server/cmd/candy-server-server/main.go --tls-certificate=./ex02/ca/minica.pem --tls-key=./ex02/ca/minica-key.pem --tls-port=3333 --port=3332

ex02_client:
	go build -o ex02/candy-cow ex02/client/main.go
	cd ex02 && ./candy-cow -k AA -c 2 -m 50

clean:
	rm -rf ex01/candy-client ex02/candy-cow