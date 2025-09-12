package fancyindex

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/code-innovator-zyx/gvm/internal/registry/base"
	"github.com/code-innovator-zyx/gvm/internal/registry/internal"
	"github.com/code-innovator-zyx/gvm/internal/version"
	"strings"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/11 下午6:23
* @Package:
 */

type Registry struct {
	base *base.Registry
}

func NewRegistry(mirrorUrl string) (*Registry, error) {
	baseRegistry, err := base.NewBaseRegistry(mirrorUrl)
	if err != nil {
		return nil, err
	}
	return &Registry{base: baseRegistry}, nil
}

func (r Registry) StableVersions() (versions []*version.Version, err error) {
	return r.AllVersions()
}

func (r Registry) UnstableVersions() (versions []*version.Version, err error) {
	return r.AllVersions()
}

func (r Registry) ArchivedVersions() (versions []*version.Version, err error) {
	return r.AllVersions()
}

func (r Registry) AllVersions() (versions []*version.Version, err error) {
	trees := r.base.Doc.Find("table").First().Find("tbody").Find("tr")
	items := make([]*internal.GoFileItem, 0, trees.Length())
	trees.Each(func(j int, tr *goquery.Selection) {
		tds := tr.Find("td")
		anchor := tds.Filter(".link").Find("a")
		href := anchor.AttrOr("href", "")
		if !strings.HasPrefix(href, "go") || strings.HasSuffix(href, "/") {
			return
		}

		items = append(items, &internal.GoFileItem{
			FileName: anchor.Text(),
			URL:      r.base.Url.JoinPath(href).String(),
			Size:     strings.TrimSpace(tds.Filter(".size").Text()),
		})
	})
	if len(items) == 0 {
		return nil, nil
	}
	return internal.Convert2Versions(items)
}
