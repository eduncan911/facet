package ctx

import (
	"fmt"
	"github.com/eduncan911/es/api"
	"github.com/eduncan911/facet/domain"
	"github.com/eduncan911/go/pkg/uuid"
	"github.com/gorilla/mux"
	"strings"
	"time"
)

// TODO add logging

// TODO make these configurable
const (
	SITECOOKIE    = "FACETS"
	AUTHSIGCOOKIE = "FACETSAUTH"
)

// Context represents the request's parsed elements
type Context struct {
	URLArticleID *domain.ArticleID
	URLUserID    *domain.UserID

	User        *User
	Fingerprint *Fingerprint
	IPAddress   *string
}

// User represents the current logged in or anonymous user
type User struct {
	UserID    *domain.UserID
	Username  *string
	Anonymous bool
}

// Fingerprint identifies the browser making the request
type Fingerprint struct {
	UserAgent string
	IPAddress string
	Language  string
}

func (f *Fingerprint) String() string {
	return fmt.Sprintf("blob \000%s;%s;%s", f.IPAddress, f.UserAgent, f.Language)
}

// Parse takes a valid api.Request and returns a parsed Context
func Parse(req *api.Request) *Context {
	c := &Context{
		URLArticleID: &domain.NoArticleID,
		URLUserID:    &domain.NoUserID,
		User: &User{
			UserID:    &domain.NoUserID,
			Anonymous: true,
		},
	}

	if req == nil || req.Raw == nil {
		return c
	}

	setupURLParams(req, c)
	setupIPAddress(req, c)
	setupFingerprint(req, c)
	setupUserContext(req, c)

	// TODO CORS
	// Access-Control-Allow-Origin
	// Access-Control-Allow-Methods
	// http://en.wikipedia.org/wiki/Cross-origin_resource_sharing
	//

	return c
}

func setupURLParams(req *api.Request, c *Context) {

	// TODO add loggin

	aid := mux.Vars(req.Raw)["articleId"]
	if aid != "" {
		id, err := uuid.ParseID(aid)
		if err == nil {
			c.URLArticleID = &domain.ArticleID{ID: id}
		}
	}

	uid := mux.Vars(req.Raw)["userId"]
	if uid != "" {
		id, err := uuid.ParseID(uid)
		if err == nil {
			c.URLUserID = &domain.UserID{ID: id}
		}
	}
}

func setupIPAddress(req *api.Request, c *Context) {

	if ip := req.Raw.Header.Get("X-Forwarded-For"); ip != "" {
		c.IPAddress = &strings.Split(ip, ",")[0]
	} else if ip := req.Raw.Header.Get("X-Real-IP"); ip != "" {
		c.IPAddress = &ip
	} else {
		c.IPAddress = &req.Raw.RemoteAddr
	}
}

func setupFingerprint(req *api.Request, c *Context) {

	c.Fingerprint = &Fingerprint{
		UserAgent: req.Raw.UserAgent(),
		IPAddress: *c.IPAddress,
		Language:  req.Raw.Header.Get("Accept-Language"),
	}
}

func setupUserContext(req *api.Request, c *Context) {

	v, valid := authCookie(req)
	if !valid {
		c.User.Username = c.IPAddress
		return
	}

	userID, err := uuid.ParseID(*v["UserID"])
	if err != nil {
		c.User.Username = c.IPAddress
		return
	}

	c.User.UserID = &domain.UserID{ID: userID}
	c.User.Username = v["Username"]
	c.User.Anonymous = false

	// TODO local lookup if user is still allowed on site
	//
}

func authCookie(req *api.Request) (map[string]*string, bool) {

	c, err := req.Raw.Cookie(AUTHSIGCOOKIE)
	if err != nil {
		return make(map[string]*string, 0), false
	}

	if time.Now().After(c.Expires) {
		return make(map[string]*string, 0), false
	}

	// decode
	//

	return make(map[string]*string, 0), false
}

func addCookie(res *api.Response) {

}
