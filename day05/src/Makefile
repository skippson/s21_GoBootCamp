all: mkbuild ex00 ex01 ex02 ex03

mkbuild:
	@rm -rf build
	@mkdir build

ex00:
	@echo "\n\033[32mex00 \033[0m"
	@go build -o build/ex00 ex00/main.go
	@./build/ex00

ex01:
	@echo "\n\033[32mex01 \033[0m"
	@go build -o build/ex01 ex01/main.go
	@./build/ex01

ex02:
	@echo "\n\033[32mex02 \033[0m"
	@go build -o build/ex02 ex02/main.go
	@./build/ex02

ex03:
	@echo "\n\033[32mex03 \033[0m"
	@go build -o build/ex03 ex03/main.go
	@./build/ex03

clean:
	@rm -rf build/

.PHONY: mkbuild ex00 ex01 ex02 ex03