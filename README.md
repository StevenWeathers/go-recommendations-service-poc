## Product Recommendations Service POC 

This is an example RESTful API in [Go](https://golang.org/)
using [Redis](https://github.com/antirez/redis)
and [Docker](https://www.docker.com/)
that returns Product Recommendations by productId

## Installation

This API is best served with [Docker](https://www.docker.com/)
```bash
docker build -t recommendations-services .
docker run -d -p 8080:8080 --name=recommendations-services recommendations-services
```
What about [Docker Compose](https://docs.docker.com/compose/)?
```bash
docker-compose up -d --build
```

What about example data
```bash
docker exec -it <dbContainerId> sh -c "cat /tmp/load/related-data.txt | redis-cli --pipe"
docker exec -it <dbContainerId> sh -c "cat /tmp/load/required-data.txt | redis-cli --pipe"
docker exec -it <dbContainerId> sh -c "cat /tmp/load/customeralsoviewed-data.txt | redis-cli --pipe"
```

## Open Source Projects
Project | License
--- | ---
[Docker](https://github.com/docker/docker) | [Apache-2.0](https://github.com/docker/docker/blob/master/LICENSE)
[Redis](https://github.com/antirez/redis) | [BSD-3-Clause](https://github.com/antirez/redis/blob/unstable/COPYING)
[Go](https://github.com/golang/go) | [BSD-3-Clause](https://github.com/golang/go/blob/master/LICENSE)


### ENDPOINT: /product/{productId}/related

#### Response
```json
{
	"related": [
		"3989748",
		"3898849",
		"3277489"
	]
}
```


### ENDPOINT: /product/{productId}/required

#### Response
```json
{
	"required": [
		"3989748",
		"3898849",
		"3277489"
	]
}
```


### ENDPOINT: /product/{productId}/customeralsoviewed

#### Response
```json
{
	"customeralsoviewed": [
		"3989748",
		"3898849",
		"3277489"
	]
}
```
