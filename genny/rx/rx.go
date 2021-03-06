package rx

import (
	"github.com/gobuffalo/genny/v2"
	packr "github.com/gobuffalo/packr/v2"
)

var templates = packr.New("github.com/gobuffalo/clara/genny/rx/templates", "../rx/templates")

// func init() {
// 	plush.Helpers.Add("partialFeeder", templates.FindString)
// }

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	g.Merge(goCheck(opts))
	if opts.SkipBuffalo {
		return g, nil
	}
	if !opts.SkipNode {
		g.Merge(nodeChecks(opts))
		g.Merge(npmChecks(opts))
		g.Merge(yarnChecks(opts))
	}

	if !opts.SkipDB {
		g.Merge(postgresChecks(opts))
		g.Merge(mysqlChecks(opts))
		g.Merge(sqliteChecks(opts))
		g.Merge(cockroachChecks(opts))
	}

	g.Merge(buffaloChecks(opts))

	return g, nil
}
