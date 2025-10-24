package responses

type Response struct {
	MetaData MetaDataResponse `json:"meta_data"`
	Data     any              `json:"data"`
}

type MetaDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Total       int32 `json:"total"`
	CurrentPage int32 `json:"current_page"`
	Limit       int32 `json:"limit"`
	Pages       int32 `json:"page"`
}
