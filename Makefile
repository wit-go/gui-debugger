# git remote add github git@github.com:wit-go/gui-debugger.git

github:
	git push -u github master
	git push -u github devel
	git push github --tags
