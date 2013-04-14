require 'base32'
require 'json'
require 'openssl'
require 'sinatra'
require 'rubygems'
#require 'pg'
require 'data_mapper'
#require 'dm-postgres-adapter'
require 'dm-core'  
require 'dm-timestamps'
require 'dm-validations'
require 'dm-migrations'
require 'pp'

#DataMapper.setup(:default, 'postgres://qbestvanfyprfa:2mg-TW896fJalG_1_UR8bV0v9x@ec2-107-21-107-194.compute-1.amazonaws.com:5432/d9v8ggrp2ui4g4')
#DataMapper.setup(:default, 'postgres://julienbordellier@localhost/requireris')
DataMapper.setup(:default, "sqlite://#{Dir.pwd}/database.db")
DataMapper.finalize

class Key
	include DataMapper::Resource
	property :id,			Serial
	property :value, 		String, :unique => true
	property :name,			String
	property :account_id,	Integer
	validates_presence_of :value, :name, :account_id
end

class Account
	include DataMapper::Resource
	property :id,			Serial
	property :mail,			String, :unique => true
	property :password,		String
	validates_presence_of :mail, :password
end

DataMapper.auto_upgrade!

enable :sessions
set :session_secret, 'this_is_the_fucking_requireris'

get '/' do
	if session[:user] != nil then
		@account = Account.all
		@key = Key.all(:account_id => session[:account_id])
		erb :index
	else
		redirect to('/login')
	end
end

get '/login' do
	erb :login
end

post '/login' do
	user = Account.all(:mail => params[:mail], :password => OpenSSL::Digest.digest("MD5", params[:password]))
	if user != [] then
		session[:user] = user.first.mail
		session[:account_id] = user.first.id
		redirect to('/')
	else
		redirect to('/login')
	end
end

get '/logout' do
	session.clear
	redirect to('/login')
end

post '/register' do
	DataMapper.finalize
	@account = Account.create(
		:mail 			=>	params[:mail],
		:password		=>	OpenSSL::Digest.digest("MD5", params[:password])
	)
	session[:user] = params[:mail]
	session[:account_id] = @account.id
	redirect to('/')
end

post '/key' do
	test = Key.new
	test.value = params[:key]
	test.name = params[:name]
	test.account_id = session[:account_id]
	test.save
	@key = Key.all(:account_id => session[:account_id])
	redirect to('/')
end

get '/otp' do
	otps = []
	keys = Key.all(:account_id => session[:account_id])
	for k in keys do
		otps << totp_gen(k.value)
	end
	otps.to_json
end

get '/delete/:id' do
	Key.first(:id => params[:id]).destroy
	redirect to('/')
end

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
