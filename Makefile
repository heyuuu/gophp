gophp-ui:
	go build -o ./temp/ ./internal/cmd/gp-ui
	./temp/gp-ui
gophp-dev:
	go build -o ./temp/ ./cmd/gophp
	./temp/gophp -r "var_dump(1+1);"