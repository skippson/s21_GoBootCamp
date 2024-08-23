
all: clean mkdir ex

mkdir:
	@mkdir build

ex:
	@go build -o build/ex00 ex00/main.go
	@go build -o build/ex01 ex01/main.go
	@go build -o build/ex02 ex02/main.go
	@go build -o build/ex03 ex03/main.go
	@echo "\n\033[32mex00-ex03 done \033[0m\n"


clean:
	@rm -rf build

fix:
	curl -XPUT -H "Content-Type: application/json" "http://localhost:9200/places/_settings" -d '{	  "index" : {"max_result_window" : 20000}	}'