.PHONY: dummy 

dummy:
	# get all dependencies
	@go get -d ./...

	@if ! [ -a $$GOPATH/bin/gopherjs ] ; \
	 then \
     echo "Installing gopherjs"; \
	 go install github.com/gopherjs/gopherjs; \
	 fi;

js: dummy
	$(MAKE) -C frontend/js $1

backend:
	$(MAKE) -C backend $1

clean:
	git clean -fd
	go clean