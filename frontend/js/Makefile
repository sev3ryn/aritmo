run:
	@go test ./... && \
	go test ../core/... && \
	$$GOPATH/bin/gopherjs build -m -o aritmo.js && \
	open index.html && \
	$$GOPATH/bin/gopherjs build -wm -o aritmo.js
