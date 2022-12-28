package pkg

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"regexp"
	"strconv"
	"strings"
)

const (
	defaultPerPage = 100
	maxPerPage     = 500
)

type Paginator struct {
	Page    uint
	PerPage uint
	OrderBy string
	RawURL  string
	Totals  int64
}

func NewPaginator(data dto.BaseSearchRequestDto) *Paginator {
	perPage := uint(defaultPerPage)
	if data.PerPage != nil {
		perPage = *data.PerPage
	}
	if perPage > 500 {
		perPage = maxPerPage
	}
	var page uint
	if data.Page != nil {
		page = *data.Page
	}
	if page == 0 {
		page++
	}
	return &Paginator{Page: page, PerPage: perPage, RawURL: data.RawURL}
}

func (p *Paginator) ToRepresentation() dto.PaginationInfo {
	return dto.PaginationInfo{
		Page:    p.Page,
		PerPage: p.PerPage,
		Total:   uint(p.Totals),
		Pages:   p.GetTotalPages(),
		Next:    p.MakeNextLink(),
		Prev:    p.MakePrevLink(),
	}
}

func (p *Paginator) ToLinkHeader() string {
	var link []string

	if p.MakeNextLink() != "" {
		link = append(link, "<"+p.MakeNextLink()+">; rel=\"next\"")
	}
	if p.MakePrevLink() != "" {
		link = append(link, "<"+p.MakePrevLink()+">; rel=\"prev\"")
	}
	if p.MakeFirstLink() != "" {
		link = append(link, "<"+p.MakeFirstLink()+">; rel=\"first\"")
	}
	if p.MakeLastLink() != "" {
		link = append(link, "<"+p.MakeLastLink()+">; rel=\"last\"")
	}

	return strings.Join(link[:], ",\n")
}

func (p *Paginator) GetOffset() uint {
	return p.PerPage * (p.Page - 1)
}

func (p *Paginator) GetLimit() uint {
	return p.PerPage
}

func (p *Paginator) GetTotalPages() uint {
	pages, mod := Divmod(int(p.Totals), int(p.PerPage))
	if mod > 0 {
		pages++
	}
	return uint(pages)
}

func (p *Paginator) MakeNextLink() string {
	totalPages := int(p.GetTotalPages())
	r := regexp.MustCompile("&page=([0-9]+)")
	if r.MatchString(p.RawURL) {
		currentPage, err := strconv.Atoi(r.FindStringSubmatch(p.RawURL)[1])
		if err != nil {
			return ""
		}
		if currentPage == 0 {
			currentPage = 1
		}
		nextItem := currentPage + 1
		if nextItem > totalPages || totalPages == 1 {
			return ""
		}
		return r.ReplaceAllString(p.RawURL, "&page="+strconv.Itoa(nextItem))
	} else if totalPages > 1 {
		return p.RawURL + "&page=2"
	}
	return ""
}

func (p *Paginator) MakePrevLink() string {
	r := regexp.MustCompile("&page=([0-9]+)")
	if r.MatchString(p.RawURL) {
		currentPage, err := strconv.Atoi(r.FindStringSubmatch(p.RawURL)[1])
		if err != nil {
			return ""
		}
		if currentPage == 0 {
			currentPage = 1
		}
		prevItem := currentPage - 1
		if prevItem <= 0 {
			return ""
		}
		return r.ReplaceAllString(p.RawURL, "&page="+strconv.Itoa(prevItem))
	}
	return ""
}

func (p *Paginator) MakeFirstLink() string {
	return p.RawURL + "&page=1"
}

func (p *Paginator) MakeLastLink() string {
	return p.RawURL + "&page=" + strconv.Itoa(int(p.GetTotalPages()))
}
