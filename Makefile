PHONEY: release test 

release: test readme
	goreleaser release --clean --config=.goreleaser.yml

snapshot:
	goreleaser build --config=.goreleaser.yml --snapshot

build: 
	go build .

readme: build
	echo "# unifi-adopt\n" > README.md
	echo "## Usage" >> README.md
	echo "" >> README.md
	echo "\`\`\`bash" >> README.md
	./unifi-adopt -h >> README.md
	echo "\`\`\`" >> README.md
	echo "" >> README.md
	echo "### version" >> README.md
	echo "" >> README.md
	echo "\`\`\`bash" >> README.md
	./dist/unifi-adopt_darwin_arm64_v8.0/unifi-adopt version >> README.md 2>&1
	echo "\`\`\`" >> README.md

test:
	go test -cover -v . ./cmd/
