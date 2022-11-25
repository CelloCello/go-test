# go-test
my golang practice


## Docs

### REST API

Gen by gin-swagger

    $ swag init -o ./app/apiserver/docs -g ./app/apiserver/route.go


## Release

Use goreleaser. Check [here](https://goreleaser.com/quick-start/).

Test release building in local.

    $ goreleaser release --snapshot --rm-dist

release

    $ goreleaser release
