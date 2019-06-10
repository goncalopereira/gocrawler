package http

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goncalocool/coolcoolcool/internal/data"
)

func TestBasicParserGetLinks(t *testing.T) {
	reader := strings.NewReader("<html><body><a href=\"link1\">Hello</a><a href=\"link2\">Hello2<a/></body></html>")

	p := BasicBodyParser{MaxLinks: 10}

	resp := p.Parse(reader, data.Request{})

	assert.Equal(t, 2, len(resp.Links))
}

func TestBasicParserPageTooManyLinks(t *testing.T) {
	//no deadlock

	reader := strings.NewReader("<html><body><a href=\"link1\">Hello</a><a href=\"link2\">Hello2<a/></body></html>")

	p := BasicBodyParser{MaxLinks: 1}

	resp := p.Parse(reader, data.Request{})

	assert.Equal(t, 1, len(resp.Links))
}
