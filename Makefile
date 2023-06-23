commit: format test
	git add -A .
	git commit -m "$(info)"
	git push origin master
list-tag:
	git tag -l
tag:
	git tag $(tag)
	git push origin $(tag)
format:
	@find ./ -type f -name "*.go" -exec gofmt -w {} \;
test:
	go test ./...
init:
	go mod init go.upyun.com/zhao.zhang/$(project)
	echo '.idea' >> .gitignore
	echo "# $(project)" >> README.md
	git init
	git remote add origin git@github.com:kougazhang/$(project).git
