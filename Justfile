build:
  go build -o algae

build-html:
  templ generate

build-css:
  tailwindcss -i templates/style.css -o static/style.css -m 

dev: 
  air

dev-html:
  templ generate -watch

dev-css:
  tailwindcss -i templates/style.css -o static/style.css -m -w
