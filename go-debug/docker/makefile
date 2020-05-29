image:
	docker build -t micro-service-example .
stop:
	docker stop server8
run:
	docker run --rm --name server8 -v $PWD:/build -p8080:8080 -p2345:2345 delve-docker-vscode-example
