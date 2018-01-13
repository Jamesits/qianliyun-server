package main

import (
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(indexContent)))
	w.Header().Set("Server", "Apache/2.2.34'<!--")
	w.Header().Set("X-AspNet-Version", "2.0.50727'<!--")
	w.Header().Set("X-Powered-By", "PHP/5.5.38'<!--")
	_, _ = w.Write([]byte(indexContent))
}

func luciHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(luciContent)))
	w.Header().Set("Server", "Apache/2.2.34'<!--")
	w.Header().Set("X-AspNet-Version", "2.0.50727'<!--")
	w.Header().Set("X-Powered-By", "PHP/5.5.38'<!--")
	_, _ = w.Write([]byte(luciContent))
}

const indexContent string = "" +
	"<?xml version=\"1.0\" encoding=\"utf-8\"?>\n" +
	"<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">\n" +
	"<html xmlns=\"http://www.w3.org/1999/xhtml\">\n" +
	"<head>\n" +
	"<meta http-equiv=\"Cache-Control\" content=\"no-cache\" />\n" +
	"<meta http-equiv=\"refresh\" content=\"0; URL=/cgi-bin/luci\" />\n" +
	"</head>\n" +
	"<body style=\"background-color: white\">\n" +
	"<a style=\"color: black; font-family: arial, helvetica, sans-serif;\" href=\"/cgi-bin/luci\">LuCI - Lua Configuration Interface</a>\n" +
	"</body>\n" +
	"</html>\n"

const luciContent string = "" +
	"<!DOCTYPE html>\n" +
	"<html lang=\"en\">\n" +
	"	<head>\n" +
	"		<meta charset=\"utf-8\">\n" +
	"		<title>LEDE - LuCI</title>\n" +
	"		<!--[if lt IE 9]><script src=\"/luci-static/bootstrap/html5.js\"></script><![endif]-->\n" +
	"		<meta name=\"viewport\" content=\"initial-scale=1.0\">\n" +
	"		<link rel=\"stylesheet\" href=\"/luci-static/bootstrap/cascade.css\">\n" +
	"		<link rel=\"stylesheet\" media=\"only screen and (max-device-width: 854px)\" href=\"/luci-static/bootstrap/mobile.css\" type=\"text/css\" />\n" +
	"		<link rel=\"shortcut icon\" href=\"/luci-static/bootstrap/favicon.ico\">\n" +
	"		<script src=\"/luci-static/resources/xhr.js\"></script>\n" +
	"	</head>\n" +
	"\n" +
	"	<body class=\"lang_en\">\n" +
	"		<header>\n" +
	"			<div class=\"fill\">\n" +
	"				<div class=\"container\">\n" +
	"					<a class=\"brand\" href=\"#\">LEDE</a>\n" +
	"					\n" +
	"					<div class=\"pull-right\">\n" +
	"						\n" +
	"						<span id=\"xhr_poll_status\" style=\"display:none\" onclick=\"XHR.running() ? XHR.halt() : XHR.run()\">\n" +
	"							<span class=\"label success\" id=\"xhr_poll_status_on\">Auto Refresh on</span>\n" +
	"							<span class=\"label\" id=\"xhr_poll_status_off\" style=\"display:none\">Auto Refresh off</span>\n" +
	"						</span>\n" +
	"					</div>\n" +
	"				</div>\n" +
	"			</div>\n" +
	"		</header><div class=\"container\">\n" +
	"				<div class=\"alert-message warning\">\n" +
	"					<h4>No password set!</h4>\n" +
	"					There is no password set on this router. Please configure a root password to protect the web interface and enable SSH.<br>\n" +
	"					<a href=\"/cgi-bin/luci/admin/system/admin\">Go to password configuration...</a>\n" +
	"				</div>\n" +
	"			</div><div id=\"maincontent\" class=\"container\">\n" +
	"			\n" +
	"\n" +
	"\n" +
	"\n" +
	"<form method=\"post\" action=\"/cgi-bin/luci\"><div class=\"cbi-map\">\n" +
	"		<h2 name=\"content\">Authorization Required</h2>\n" +
	"		<div class=\"cbi-map-descr\">\n" +
	"			Please enter your username and password.\n" +
	"		</div>\n" +
	"		<fieldset class=\"cbi-section\"><fieldset class=\"cbi-section-node\">\n" +
	"			<div class=\"cbi-value\">\n" +
	"				<label class=\"cbi-value-title\">Username</label>\n" +
	"				<div class=\"cbi-value-field\">\n" +
	"					<input class=\"cbi-input-user\" type=\"text\" name=\"luci_username\" value=\"root\" />\n" +
	"				</div>\n" +
	"			</div>\n" +
	"			<div class=\"cbi-value cbi-value-last\">\n" +
	"				<label class=\"cbi-value-title\">Password</label>\n" +
	"				<div class=\"cbi-value-field\">\n" +
	"					<input class=\"cbi-input-password\" type=\"password\" name=\"luci_password\" />\n" +
	"				</div>\n" +
	"			</div>\n" +
	"		</fieldset></fieldset>\n" +
	"	</div>\n" +
	"\n" +
	"	<div>\n" +
	"		<input type=\"submit\" value=\"Login\" class=\"cbi-button cbi-button-apply\" />\n" +
	"		<input type=\"reset\" value=\"Reset\" class=\"cbi-button cbi-button-reset\" />\n" +
	"	</div>\n" +
	"</form>\n" +
	"<script type=\"text/javascript\">//<![CDATA[\n" +
	"	var input = document.getElementsByName('luci_password')[0];\n" +
	"	if (input)\n" +
	"		input.focus();\n" +
	"//]]></script>\n" +
	"\n" +
	"\n" +
	"\n" +
	"\n" +
	"   <footer>\n" +
	"    <a href=\"https://github.com/openwrt/luci\">Powered by LuCI lede-17.01 branch (git-17.290.79498-d3f0685)</a> / LEDE Reboot 17.01.4 r3560-79f57e422d\n" +
	"    \n" +
	"   </footer>\n" +
	"   </div>\n" +
	"  </div>\n" +
	" </body>\n" +
	"</html>\n" +
	"\n" +
	"\n"
