dev: 
  air

dev-html:
  templ generate -watch

dev-css:
  tailwindcss -i templates/style.css -o static/style.css -m -w