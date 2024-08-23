all: ex00 ex01 ex02 ex03

ex00: clean
	@echo "\n\033[32mex00 \033[0m"
	@go build -o myFind ex00/main.go
	@./myFind test_dir

ex01:
	@echo "\n\033[32mex01 \033[0m"
	@go build -o myWc ex01/main.go
	@./myWc -m test_dir/for_testing.go test_dir/for_testing.h test_dir/orig.txt

ex02:
	@echo "\n\033[32mex02 \033[0m"
	@go build -o myXargs ex02/main.go
	@./myFind -f -ext 'go' test_dir | ./myXargs ./myWc -m

ex03:
	@go build -o myRotate ex03/main.go
	@./myRotate -a archive ex01/main.go test_dir/for_testing.h test_dir/orig.txt
	@echo "\n\033[32mex03 done \033[0m"

clean :
	@rm -rf myFind myWc myXargs myRotate archive

.PHONY : ex00 ex01 ex02 ex03