
generate:
	docker run --rm \
	  -v ${PWD}/config:/config \
	  -v ${PWD}/out:/out \
	  openapitools/openapi-generator-cli generate \
	  -i /config/app-store-connect-openapi-spec.json \
	  -g go \
	  -c /config/config.yml \
	  -o /out/
