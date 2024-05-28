module github.com/lepingbeta/go-common-v2-dh-utils

replace github.com/lepingbeta/go-common-v2-dh-log => ../go-common-v2-dh-log

go 1.22.1

require (
	github.com/lepingbeta/go-common-v2-dh-json v0.0.0-20240518060951-a591de6cf4ec
	github.com/lepingbeta/go-common-v2-dh-log v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
	go.mongodb.org/mongo-driver v1.15.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
