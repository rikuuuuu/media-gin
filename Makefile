APPDIR := app

PHONY: run
run:
	go run $(APPDIR)/server.go

PHONY: format
format:
	goimports -w $(APPDIR)/*.go
	goimports -w $(APPDIR)/domain/model/*.go
	goimports -w $(APPDIR)/infra/*.go
	goimports -w $(APPDIR)/interfaces/controllers/*.go
	goimports -w $(APPDIR)/interfaces/database/*.go
	goimports -w $(APPDIR)/usecase/*/*.go
	goimports -w $(APPDIR)/utils/*.go