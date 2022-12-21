package cmc_rater

import (
	"net/url"

	"github.com/aitsvet/coinconv/internal/app/model"
)

type HttpClient interface {
	Get(path string) (body string, err error)
}

type JSONExtractor interface {
	Extract(from string, path []string) (result string, err error)
}

type RateReader interface {
	Read(from string) (to model.Rate, err error)
}

type Rater struct {
	httpClient HttpClient
	extractor  JSONExtractor
	reader     RateReader
}

func New(httpClient HttpClient, extractor JSONExtractor, reader RateReader) *Rater {
	return &Rater{httpClient: httpClient, reader: reader, extractor: extractor}
}

func (r *Rater) GetRate(from, to model.Currency) (model.Rate, error) {
	q := url.Values{}
	q.Add("symbol", from.Symbol())
	q.Add("convert", to.Symbol())
	u := url.URL{Path: "v1/cryptocurrency/quotes/latest"}
	u.RawQuery = q.Encode()
	body, err := r.httpClient.Get(u.RequestURI())
	if err != nil {
		return nil, err
	}
	result, err := r.extractor.Extract(body, []string{"data", from.Symbol(), "quote", to.Symbol(), "price"})
	return r.reader.Read(result)
}
