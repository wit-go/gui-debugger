# git remote add github git@github.com:wit-go/gui-debugger.git

all:
	go mod tidy

github:
	git push -u github master
	git push -u github devel
	git push github --tags
