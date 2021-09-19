package requests

type BrowseFilesRequest struct {
	PerPage int    `query:"per_page"`
	Page    int    `query:"page"`
	Search  string `query:"search"`
	Order   string `query:"order"`
	Sort    string `query:"sort"`
}
