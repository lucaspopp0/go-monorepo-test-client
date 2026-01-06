package gomonorepotestclient

import (
	"github.com/lucaspopp0/go-monorepo-test/a"
	"github.com/lucaspopp0/go-monorepo-test/a/c"
	"github.com/lucaspopp0/go-monorepo-test/a/c/d"
	"github.com/lucaspopp0/go-monorepo-test/b"
)

func main() {
	_ = a.A{}
	_ = b.B{}
	_ = c.C{}
	_ = d.D{}
}
