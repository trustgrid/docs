
build:
	hugo --minify

publish:
	aws s3 sync public/ s3://tg-dev-docs/ --delete 

run:
	hugo server

clean:
	rm -rf public