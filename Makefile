build:
	hugo --minify

publish:
	aws s3 sync public/ s3://tg-dev-docs/ --delete

run:
	hugo server

clean:
	rm -rf public

apidocs:
	npx @redocly/cli build-docs static/swagger.yml --output=static/apidocs.html