package test

type Assistant struct {
	Model        string           `json:"model"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Instructions string           `json:"instructions"`
	Tools        []AssistantTools `json:"tools"`
	Metadata     struct {
	} `json:"metadata"`
}

type AssistantTools struct {
	Type string `json:"type"`
}
