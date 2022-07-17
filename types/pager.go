package types

import "fmt"

type Pager struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

func (p *Pager) Validate() error {
	if p.Offset < 0 {
		return fmt.Errorf("Invalid offset")
	}

	if p.Limit <= 0 || p.Limit > 20 {
		return fmt.Errorf("Invalid limit")
	}

	return nil
}

type PagerResult struct {
	Next     int64       `json:"next"`
	Previous int64       `json:"previous"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

func (p *PagerResult) GetResult(pager *Pager, totalDocs int64, data interface{}) *PagerResult {
	var current int64

	if pager.Offset == 0 {
		current = 1 * pager.Limit
	} else {
		current = pager.Offset * pager.Limit
	}

	if current < totalDocs {
		p.Next = pager.Offset + 1
	} else {
		p.Next = 0
	}

	if current > pager.Limit {
		p.Previous = pager.Offset - 1
	} else {
		p.Previous = 0
	}

	p.Total = totalDocs
	p.Data = data

	return p
}
