# Bitly URL shortner API in Golang, callable from Python

Created a library function in Golang for generating short URL links using the Bitly API. This function is exported for Python usage, as the intent for this was to integrate it with StormBreaker and have it swith the ngrok links with more common/trusted(?) Bitly URLs. 

>Requirements:

* github.com/retgits/bitly/client

>Interface:

getLink(url *C.char)
