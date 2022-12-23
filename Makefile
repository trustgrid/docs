
build:
	hugo --minify

publish:
	aws s3 sync public/ s3://tg-dev-docs/ --delete 
	aws cloudfront create-invalidation --distribution-id E1NX79FZ8CBQEU --paths "/*"

run:
	hugo server