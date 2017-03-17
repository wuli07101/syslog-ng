local socketHttp = require("socket.http")
local cjson = require "cjson"
local ltn12 = require("ltn12")

function lua_init()
	socketHttp.TIMEOUT = 1
end

function lua_queue(msg)
	local sampleJson = msg['MESSAGE']
    local data = cjson.decode(sampleJson)
	
	print(data["url"])

	local request_body = [[login=user&password=123]]	
	local response_body = {}
	local result, statuscode, content = socketHttp.request{
		url = "http://www.baidu.com",
        method = "POST",
		headers =
        {
            ["Content-Type"] = "application/x-www-form-urlencoded",
            ["Content-Length"] = #request_body,
        },
		source = ltn12.source.string(request_body),
		sink = ltn12.sink.table(response_body),
	}


--    local result, respcode, respheaders, respstatus = socketHttp.request {
--        url = data["url"],
--        method = "POST",
    --    source = ltn12.source.string(data["data"]),
--    }

	print(statuscode,result)
	print(table.concat(response_body))
end
