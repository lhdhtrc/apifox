package main

type ImportApiOptions struct {
	TargetEndpointFolderID        int64  `json:"targetEndpointFolderId"`
	TargetSchemaFolderID          int64  `json:"targetSchemaFolderId"`
	EndpointOverwriteBehavior     string `json:"endpointOverwriteBehavior"`
	SchemaOverwriteBehavior       string `json:"schemaOverwriteBehavior"`
	UpdateFolderOfChangedEndpoint bool   `json:"updateFolderOfChangedEndpoint"`
	PrependBasePath               bool   `json:"prependBasePath"`
}

type ImportApiBody struct {
	Input   string           `json:"input"`
	Options ImportApiOptions `json:"options"`
}

type ApiFoxConfig struct {
	Body      ImportApiBody `json:"body"`
	Token     string        `json:"token"`
	ProjectId string        `json:"project_id"`
}
