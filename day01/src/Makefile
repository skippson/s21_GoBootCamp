all: ex00 ex01 ex02

ex00: clean
	@echo "\n\033[32mex00 \033[0m"
	@go build -o readDB ex00/main.go 
	@./readDB -f recipes/original.xml

ex01: clean
	@echo "\n\n\033[32mex01 \033[0m"
	@go build -o compareDB ex01/main.go
	@./compareDB --old recipes/original.xml --new recipes/snitch.json

ex02: clean
	@echo "\n\033[32mex02 \033[0m"
	@go build -o compareFS ex02/main.go
	@./compareFS --old snapshots/snapshot1.txt --new snapshots/snapshot2.txt

clean:
	@rm -rf readDB compareDB compareFS

.PHONY: ex00