#!/usr/bin/env ruby
# encoding: utf-8
$:.push('gen-rb')

require 'thrift'
require 'rpc_service'

begin
  port = ARGV[0] || 19090

  transport = Thrift::FramedTransport.new(Thrift::Socket.new('127.0.0.1', port))
  protocol = Thrift::BinaryProtocol.new(transport)
  client = RpcService::Client.new(protocol)

  transport.open()

  100.times do
    puts time = Time.now
    puts client.funCall(Time.now.to_i, "login", {"name" => "wzc"})
    puts "time is #{Time.now - time}"
  end

  transport.close()

rescue Thrift::Exception => tx
  print 'Thrift::Exception: ', tx.message, "\n"
end
