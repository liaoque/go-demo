package weibo

type Article struct {
	Id int	`json:"id"`
	OUid string	`json:"o_uid"`
	Text string	`json:"text"`
	Title string `json:"title"`
	CreatedAt string `json:"created_at"`
	OtherId string `json:"other_id"`
}
