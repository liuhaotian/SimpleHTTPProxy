SimpleHTTPProxy
===============

Simple HTTP proxy in Go

Note
----
The 'smart' Go http DefaultServeMux uses cleanPath to alter the raw path. For example, http://127.0.0.1:8080/http://12.com will be redirected to http://127.0.0.1:8080/http:/12.com.

The 'smart' Go http CheckRedirect wants to follow the redirection. For instance, http://m.cnn.com uses redirection to set cookie and Go just follow the redirection without saving the cookie.