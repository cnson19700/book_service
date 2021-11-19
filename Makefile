herokulog:
	heroku logs --tail --app book-sservice
	heroku logs --tail --app authent-sservice
	heroku logs --tail --app comment-sservice
	heroku logs --tail --app users-sservice

herokurestart:
	heroku dyno:restart --app authent-service
	heroku dyno:restart --app book-sservice
	heroku dyno:restart --app users-sservice
	heroku dyno:restart --app comment-sservice

run:
	@GO111MODULE=off go get -u github.com/husol/fresh
	fresh -c fresh.conf

lint:
	@hash golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1
	@GO111MODULE=on CGO_ENABLED=0 golangci-lint run

# gci:
# 	@GO111MODULE=off go get github.com/daixiang0/gci
# 	gci -w -local trungpham/gowebbasic .

# migratedown:
# 	migrate -path migration--database "mysql://root:Teotu_19@tcp(127.0.0.1:3306)/the_movie_db?charset=utf8mb4&parseTime=True" --verbose down

# migrateup:
# 	migrate -path migration --database "mysql://root:Teotu_19@tcp(127.0.0.1:3306)/the_movie_db?charset=utf8mb4&parseTime=True" --verbose up
