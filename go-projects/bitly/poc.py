#!/usr/local/bin/python3

import ctypes
from sys import argv

def main():
	try:
		assert len(argv) >= 2	
		so = ctypes.cdll.LoadLibrary("./_shorten.so")
		url = argv[1]
		so.getLink(url.encode())
	except AssertionError as e:
		print("Usage " + argv[0] + " <url>")
if __name__ =="__main__": main()
	
