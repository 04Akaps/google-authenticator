package types

type AddReq struct {
	URL           string   `json:"url" binding:"required"`
	CardSelector  string   `json:"cardSelector" binding:"required"`
	InnerSelector string   `json:"innerSelector" binding:"required"`
	Tag           []string `json:"tag" binding:"required"`
}

type ViewReq struct {
	URL string `json:"url" binding:"required"`
}

type UpdateReq struct {
	URL           string   `json:"url" binding:"required"`
	CardSelector  string   `json:"cardSelector" binding:"required"`
	InnerSelector string   `json:"innerSelector" binding:"required"`
	Tag           []string `json:"tag" binding:"required"`
}

type DeleteReq struct {
	URL string `json:"url" binding:"required"`
}
