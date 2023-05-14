package stclient

//go:generate openapi-generator generate -i https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json -g go --additional-properties=packageName=stclient --additional-properties=enumClassPrefix=true
