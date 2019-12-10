APPS_FOLDER=apps
APP_TEMPLATES_FOLDER=setup-cobra

# help / usage
print-%: ; @echo $*=$($*)
guard-%:
	@test ${${*}} || (echo "FAILED! Environment variable $* not set " && exit 1)
	@echo "-> use env var $* = ${${*}}";

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

cobra.apps:
	mkdir -p $(APPS_FOLDER)

## cobra.apps.list:   : list cobra apps in APPS_FOLDER ( e.g.: make cobra.apps.list )
cobra.apps.list: cobra.apps
	@echo "Your cobra apps in: $(APPS_FOLDER) ..."
	ls -la $(APPS_FOLDER)

## cobra.app.build:   : compile cobra app ( e.g.: make cobra.app.build COBRA_APP=example-001 )
cobra.app.build: guard-COBRA_APP cobra.apps
	cd $(APPS_FOLDER)/$(COBRA_APP) && go build && ./app --help && ls -la
	@echo "Your app is build here: ./$(APPS_FOLDER)/$(COBRA_APP)/app"

## cobra.app.create:   : create a new cobra app ( e.g.: make cobra.app.create COBRA_APP=example-001 )
cobra.app.create: guard-COBRA_APP cobra.apps
	cobra init --pkg-name app $(APPS_FOLDER)/$(COBRA_APP)
	cp $(APP_TEMPLATES_FOLDER)/go.mod $(APPS_FOLDER)/$(COBRA_APP)
	cp $(APP_TEMPLATES_FOLDER)/README.md $(APPS_FOLDER)/$(COBRA_APP)
	cd $(APPS_FOLDER)/$(COBRA_APP) && ls -la && go build
	make cobra.app.build COBRA_APP=$(COBRA_APP)
	cd $(APPS_FOLDER)/$(COBRA_APP) && ./app --help

## cobra.app.delete:   : delete a cobra app ( e.g.: make cobra.app.delete COBRA_APP=example-001 )
cobra.app.delete: guard-COBRA_APP cobra.apps.list
	@echo "Your app will be deleted from: $(APPS_FOLDER)/$(COBRA_APP) ..."
	ls -la $(APPS_FOLDER)/$(COBRA_APP)
	rm -rf $(APPS_FOLDER)/$(COBRA_APP)
	@echo "Your app has been deleted from: $(APPS_FOLDER)/$(COBRA_APP)"
	make cobra.apps.list

## cobra.command.create:   : add a new command to cobra app ( e.g.: cobra.command.create COBRA_APP=example-001 )
cobra.command.create: guard-COBRA_APP guard-COBRA_COMMAND
	cd $(APPS_FOLDER)/$(COBRA_APP) && cobra add $(COBRA_COMMAND)
	make cobra.app.build COBRA_APP=$(COBRA_APP)
	@echo "app: $(COBRA_APP) command: $(COBRA_COMMAND)"
	cd $(APPS_FOLDER)/$(COBRA_APP) && ls -la
	cd $(APPS_FOLDER)/$(COBRA_APP) && ./app --help
	cd $(APPS_FOLDER)/$(COBRA_APP) && ./app $(COBRA_COMMAND) --help
