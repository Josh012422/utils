install:
	@echo "**********Downloading and Installing Dependencies**********"
	@go get -u github.com/spf13/cobra/cobra
	@go get -u github.com/mgutz/ansi
	@go get -u github.com/Josh012422/gocharm/cmd
	@go get -u github.com/Josh012422/gocharm/misc
	@echo "**********Installing CLI**********"
	@go install
	@echo "**********Creating config file**********"
	@go run main.go config
	@echo "***Succesfully installed CLI***"
	@cat .bin.txt
	@echo "Done. Now you can run the cli by typing gocharm [command] or just utils for a help menu"
