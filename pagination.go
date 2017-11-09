// Copyright 2017 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gdax

import (
	"net/url"
	"strconv"
)

type PaginationParams struct {
	Limit  int
	Before string
	After  string
	Extra  map[string]string
}

func (p *PaginationParams) Encode(direction string) string {
	values := url.Values{}

	if p.Limit > 0 {
		values.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Before != "" && direction == "prev" {
		values.Add("before", p.Before)
	}
	if p.After != "" && direction == "next" {
		values.Add("after", p.After)
	}

	for k, v := range p.Extra {
		values.Add(k, v)
	}

	return values.Encode()
}

func (p *PaginationParams) Done() bool {
	if p.Before == "" && p.After == "" {
		return true
	}

	return false
}

func (p *PaginationParams) AddExtraParam(key, value string) {
	if p.Extra == nil {
		p.Extra = make(map[string]string)
	}
	p.Extra[key] = value
}
