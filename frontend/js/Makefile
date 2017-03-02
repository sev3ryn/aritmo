run:
	@go test ./... && \
	go test ../core/... && \
	$$GOPATH/bin/gopherjs build -m -o static/aritmo.js && \
	open index.html && \
	$$GOPATH/bin/gopherjs build -wm -o static/aritmo.js
