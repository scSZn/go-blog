package app

type Pager struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p *Pager) GetOffset() int {
	if p.Page <= 0 || p.Limit <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit
}

func (p *Pager) GetLimit() int {
	if p.Limit <= 0 {
		return 20 // todo: 配置化？
	}
	return p.Limit
}

func (p *Pager) GetPage() int {
	if p.Page <= 0 {
		return 0
	}
	return p.Page
}
