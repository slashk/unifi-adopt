PHONEY: release test 

release: test
	goreleaser release --clean --config=.goreleaser.yml

snapshot:
	goreleaser build --config=.goreleaser.yml --snapshot --clean

build: 
	go build .

readme: build snapshot
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
	go test -coverprofile=cov.xml -v . ./cmd/

dep:
	# go get -u github.com/golang/lint/golint
	# go get -u github.com/opennota/check/cmd/aligncheck
	# go get -u github.com/opennota/check/cmd/structcheck
	# go get -u github.com/opennota/check/cmd/varcheck
	go get -u
	go mod tidy