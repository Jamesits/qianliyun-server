package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("=\x1b\x8d\xc1_\xcf8\xeb+\xde\xc4\xe1\x16\xd9\xb2\xd0G\x88&\xd8Uj\x8e\xc46p[\xb4\xdb\xf8\xaaLc\x19\xc1Cp\xd7\x07w\x80\x0b\xf8g9\xea\x07G\xe9J#\xf7\x89\xb1\xa2\xb0\xedO\xc3\xe5\xc4\x0cs\x1b"), []byte("L\x02`\x12\xba\xbf\x00\x99\xa2\xbaQwa\x03\x7ft\xc6b\x8e\xa4<\xe2\xbc\xd5GP\x82Tl\x92\xad5\x94j\x88\xbc2|\x7f&|*\x0f\x0b\x1c@0u\x1f(\xb8\xf9Ep\xfa\x14\xa3\xe0\x95\xde\xe9\xf5 \x93"))

func main() {
	http.HandleFunc("/api/login.php", loginHandler)
	http.HandleFunc("/api/query_user_info.php", queryUserInfoHandler)
	http.HandleFunc("/api/query_reseller_info.php", nil)
	http.HandleFunc("/api/new_live_session.php", newLiveSessionHandler)
	http.HandleFunc("/api/update_live_session.php", updateLiveSessionHandler)
	http.HandleFunc("/api/query_live_session.php", queryLiveSessionHandler)
	http.HandleFunc("/api/update_customer_info.php", updateCustomerInfoHandler)
	http.HandleFunc("/api/query_customer_info.php", queryCustomerInfoHandler)
	log.Fatalln(http.ListenAndServe("[::1]:6452", nil))
}
