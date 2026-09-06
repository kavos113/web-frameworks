package openapi

//go:generate go tool oapi-codegen -config model-cfg.yaml ../../../../../specs/openapi.yaml
//go:generate go tool oapi-codegen -config server-cfg.yaml ../../../../../specs/openapi.yaml
//go:generate go tool oapi-codegen -config spec-cfg.yaml ../../../../../specs/openapi.yaml
