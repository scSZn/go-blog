package app

type Pager struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p *Pager) GetOffset() int {
	if p.Page <= 0 || p.Limit <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit
}
