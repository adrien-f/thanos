package ui

import (
	"github.com/improbable-eng/thanos/pkg/objstore/client"
	"html/template"
	"net/http"
	"path"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/route"
)

type BucketUI struct {
	*BaseUI
	flagsMap map[string]string
	buckets  []client.BucketConfig
}

func NewBucketUI(logger log.Logger, flagsMap map[string]string, buckets []client.BucketConfig) *BucketUI {
	return &BucketUI{
		BaseUI:   NewBaseUI(logger, "bucket_menu.html", bucketTmplFuncs()),
		flagsMap: flagsMap,
		buckets:  buckets,
	}
}

func bucketTmplFuncs() template.FuncMap {
	return template.FuncMap{}
}
func (b *BucketUI) bucket(w http.ResponseWriter, r *http.Request) {
	prefix := GetWebPrefix(b.logger, b.flagsMap, r)
	b.executeTemplate(w, "buckets.html", prefix, b)
}

// root redirects / requests to /bucket, taking into account the path prefix value
func (b *BucketUI) root(w http.ResponseWriter, r *http.Request) {
	prefix := GetWebPrefix(b.logger, b.flagsMap, r)

	http.Redirect(w, r, path.Join(prefix, "/bucket"), http.StatusFound)
}

func (b *BucketUI) Register(r *route.Router) {
	instrf := prometheus.InstrumentHandlerFunc

	r.Get("/", instrf("root", b.root))
	r.Get("/bucket", instrf("bucket", b.bucket))

	r.Get("/static/*filepath", instrf("static", b.serveStaticAsset))
}
