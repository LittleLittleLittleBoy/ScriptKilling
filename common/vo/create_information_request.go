package vo

type CreateInformationRequest struct {
	Type  int    `json:"type"`
	Title string `json:"title"`
	Stage int    `json:"stage"`
	URL   string `json:"url"`
	Rid   int    `json:"rid"`
	Sid   int    `json:"sid"`
}
