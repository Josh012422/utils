install:
	@echo "**********Downloading and Installing Dependencies**********"
	@go get -u github.com/spf13/cobra/cobra
	@go get -u github.com/mgutz/ansi
	@echo "**********Installing Cli**********"
	@go install
	@echo "Done. Now you can run the cli by typing cli [command] or just cli for a help menu"
