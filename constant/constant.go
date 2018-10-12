package constant

import (
	
	"github.com/gorilla/sessions"

)

var Store = sessions.NewCookieStore([]byte("xes-something-very-secret"))
