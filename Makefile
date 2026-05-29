.PHONY: build publish run clean index-docs

build:
	hugo --minify

publish:
	aws s3 sync public/ s3://tg-dev-docs/ --delete

run:
	hugo server

clean:
	rm -rf public

index-docs:
	git diff --name-status main > changed_files.txt
	@echo "=== Changed files ==="
	@cat changed_files.txt
	@echo "====================="
	python3 .github/scripts/index_changed_docs.py changed_files.txt
