def hotp_gen(password, interval_no)
	b = Base32.decode(password.upcase)
	dig = OpenSSL::HMAC.digest(OpenSSL::Digest::Digest.new('sha1'), b, interval_no.pack('Q>'))
	offset = dig[19].ord & 15
	dt = dig[offset..(offset+4)]
	ui = dt.to_s.unpack('I>')
	otp = (ui[0] & 0x7fffffff) % 1000000
	return otp.to_s
end

def totp_gen(password)
	return hotp_gen(password, [Time.now.to_f / 30])
end