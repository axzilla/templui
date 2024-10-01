# Run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
templ:
	templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

# Run air to detect any go file changes to re-build and re-run the server.
server:
	air \
	--build.cmd "go build -o tmp/bin/main ./cmd/server" \
	--build.bin "tmp/bin/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# Run tailwindcss to generate the styles.css bundle in watch mode.
tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

# Start development server
dev:
	make -j3 templ server tailwind

# Start development server in debug mode - It's needed to start launch.json in VSCode
debug:
	make -j2 templ tailwind
