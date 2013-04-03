import sys
import base64
import hashlib
import hmac
import struct
import time

password = "ymybvnckruprgkgr"

def hotp_gen(password, interval):
	key = base64.b32decode(password, casefold=True)
	big = struct.pack('>Q', interval) # set unsigned long long bigendian
	dig = hmac.new(key, big, hashlib.sha1).digest()
	offset = ord(dig[19])&15
	dt = dig[offset:offset + 4]
	print base64.b32encode(dt)
	uint = struct.unpack('>I', dt) # reset to unsigned int bigendian
	print uint
	otp = (uint[0] & 0x7fffffff) % 1000000
	return otp

if len(sys.argv) > 1:
	print hotp_gen(sys.argv[1], time.time()/30)