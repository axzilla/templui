# Dev Tools
templ:
	templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

server:
	air \
	--build.cmd "go build -o tmp/bin/main ./cmd/server" \
	--build.bin "tmp/bin/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

dev:
	make -j3 templ server tailwind

debug:
	make -j2 templ tailwind

# Package Tools
generate-icons:
	go run cmd/icongen/main.go

generate-lib:
	templ generate ./pkg/...

generate-css:
	npx tailwindcss -i ./assets/css/input.css -o ./pkg/styles/goilerplate.css

update-package: generate-css generate-lib # generate-icons <- if needed
	@echo "Package updated successfully"
