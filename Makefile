install:
	@echo "**********Downloading and Installing Dependencies**********"
	@go get -u github.com/spf13/cobra/cobra
	@go get -u github.com/mgutz/ansi
	@go get -u github.com/Josh012422/utils/cmd
	@go get -u github.com/Josh012422/utils/misc
	@echo "**********Installing CLI**********"
	@go install
	@echo "**********Creating config file**********"
	@touch .config.yml
	@echo 'default: ""' >> .config.yml
	@echo "*****File succesfully created in cmd/.config.yml*****"
	@echo "***Succesfully installed CLI***"
	@cat .bin.txt
	@echo "Done. Now you can run the cli by typing utils [command] or just utils for a help menu"
