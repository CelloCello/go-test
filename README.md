# go-test
my golang practice


## Run

### API Server

    $ go run ./app/apiserver

## Docs

### REST API

Gen by gin-swagger

    $ swag init -o ./app/apiserver/docs -g ./app/apiserver/main.go


## Release

Use goreleaser. Check [here](https://goreleaser.com/quick-start/).

Test release building in local.

    $ goreleaser release --snapshot --rm-dist

release

    $ goreleaser release


## Lint and Test

    $ make lint
    $ make test
