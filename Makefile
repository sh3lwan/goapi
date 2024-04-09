build:
	@templ generate view
	@go build -o bin/webgo main.go

run:
	@./bin/webgo


tailwind:
	@npx tailwindcss -i ./static/css/style.css -o ./public/style.css --watch

