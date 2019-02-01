package navervod

import (
	"net/url"
)

var baseurl, _ = url.Parse("https://apis.naver.com/rmcnmv/rmcnmv/vod/play/v2.0/")

func Playurl(vid, key string, q url.Values) string {
	if q == nil {
		q = make(url.Values, 1)
	}
	_, ok := q["key"]
	if !ok {
		q.Set("key", key)
	}
	u := *baseurl
	u.Path += key
	u.RawQuery = q.Encode()
	return u.String()
}
