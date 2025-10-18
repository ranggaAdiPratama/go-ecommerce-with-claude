package responses

type Response struct {
	MetaData MetaDataResponse `json:"meta_data"`
	Data     any              `json:"data"`
}

type MetaDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
