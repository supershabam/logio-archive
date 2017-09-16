package logio

import (
	"net"
	"net/url"
	"os"
)

type Logio interface {
	Info(string)
	WithIP(string, net.IP) Logio
	WithString(string, string) Logio
	WithFloat64(string, string) Logio
	WithURL(string, url.URL) Logio
}

/*
types of context

collector axiom
- linux version
- origin/hostname
- timestamp normalization

service axiom - can't be overriden once set
- version
- service name
*/

type Emitter interface {
}

var version string

func NewEmitter(service []string, opts ...EmitterOptionFunc) {

}

func main() {
	eio, err := emitio.NewEmitter(
		[]string{"vulcan", "indexer"},
		emitio.WithFields(map[string]emitio.Valuer{
			"environment": os.Getenv("ENV"),
			"version":     emitio.MustSemver(version),
		}),
		emitio.WithOutput(os.Stderr),
	)
	if err != nil {
		panic(err)
	}

	e := eio.Event()
	e2 := e.Induce().Event("took our jobs")
	i := e.Induce().SerializeHTTP(w)
}

// 	eio, err := NewLogio(
// 		WithAxioms(

// 		)
// 		WithContext(
// 			WithString
// 		)
// 	)

// 	e, err := eio.Hydrate(httprequest)
// 	if err != nil {
// 		eio.WithError(err).Error("unable to hydrate http request into event")
// 	}

// 	func do(ctx context.Context, r *http.Request) {
// 		event, err := eio.Extract(r)
// 		if err != nil {
// 			eio.WithError(err).Error()
// 		}
// 		ec := event.Child("fetching from database",
// 			eio.WithFields(
// 				eio.String("lulz", "lulz")
// 			),
// 		)
// 		defer ec.End()
// 		r, err := db.Do("thing")
// 		if err != nil {
// 			ec.WithError(err)
// 		}
// 	}

// }
