ngx.say('test test test aaa')
--[[local redis = require "redis_iresty"
local red = redis:new()
local ok, err = red:set("dog", "an animal")
if not ok then
	ngx.say("failed to set dog: ", err)
	return
end]]--
local _M = {}
local redis = require "resty.redis"
function _M.run()
ngx.say("run")
local red = redis:new()

red:set_timeout(1000)

local ok, err = red:connect("127.0.0.1", 6379)
if not ok then
	ngx.say("failed to connect ", err)
	return
end

ok, err = red:set("dog", "an animal")
if not ok then
	ngx.say("failed to set dog: ", err)
	return
end

ngx.say("set result: ", ok)

local ok, err = red:set_keepalive(10000, 20)
if not ok then
	ngx.say("failed to set keepalive: ", err)
	return
end
end
return _M
