package aojchecker

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Record struct {
	Rank      string
	UserName  string
	Date      string
	Language  string
	Version   string
	CPUMemory string
}

type Records []*Record

func (r *Record) String() string {
	return fmt.Sprintf("Rank:%v, UserName:%v, Date:%v, Language:%v, Version:%v, CPUMemory:%v", r.Rank, r.UserName, r.Date, r.Language, r.Version, r.CPUMemory)
}

func (r Records) HasUser(userName string) (exists bool) {
	for _, record := range r {
		if userName == record.UserName {
			return true
		}
	}
	return false
}

func GetDoc(id string) (doc *goquery.Document, err error) {
	doc, err = goquery.NewDocument(fmt.Sprintf("http://judge.u-aizu.ac.jp/onlinejudge/problem_statistics.jsp?id=%v#1", id))
	return doc, err
}

func ParseRecords(doc *goquery.Document) (r Records) {
	filterUser := func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		a := strings.Split(s.Text(), "\n")
		u := &Record{}

		u.Rank = strings.TrimSpace(a[1])
		u.UserName = strings.TrimSpace(a[2])
		u.Date = strings.TrimSpace(a[3])
		u.Language = strings.TrimSpace(a[5])
		u.Version = strings.TrimSpace(a[6])
		u.CPUMemory = strings.TrimSpace(a[7])

		r = append(r, u)
	}

	s := doc.Find(".status_list")
	s.Find("tr").Each(filterUser)

	return r
}
