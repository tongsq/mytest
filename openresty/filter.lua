--ngx.say('test test test aaa')
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
--	ngx.exit(403)
	local redis = require "redis_iresty"
	local red = redis:new()
	local key = "f" .. ngx.var.remote_addr 
	local ok, err = red:incr(key)
--	ngx.say("key",key,"incr",ok)
	if ok > 100 then
--		ngx.say("forbiden")
		ngx.exit(ngx.HTTP_FORBIDDEN)
	end
	if ok == 1 then
		red:expire(key, 5)
	end
end
return _M
