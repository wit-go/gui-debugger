# git remote add github git@github.com:wit-go/gui-debugger.git

all:
	@echo a 'gui' debugger using the gui

redomod:
	rm -f go.*
	go mod init
	go mod tidy

github:
	git push -u github master
	git push -u github devel
	git push github --tags
