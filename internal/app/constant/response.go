package constant

type MetadataResponse struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	PageCount  int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
	Links      any   `json:"links"`
}

type StandardResponse struct {
	ResponseCode int              `json:"response_code"`
	Message      string           `json:"message"`
	Data         any              `json:"data"`
	Metadata     MetadataResponse `json:"metadata"`
}
