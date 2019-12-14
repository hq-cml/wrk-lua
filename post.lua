idArr = {}
function init(args)
   falg = 1 
   for line in io.lines("body.txt") do
       print(line)
       idArr[falg] = line
       falg = falg + 1
   end
   falg = 0
end
 
request = function()
   --parms = idArr[falg%(table.getn(idArr) + 1)]
   --print(falg % table.getn(idArr))
   --print(falg % table.getn(idArr), parms)
   --falg = falg + 1
   score = math.random(1, 100)
   objId = math.random(1, 10)
   body = '{"objId":'..objId..', "score":'..score..', "bizIds":[1001,2001], "bizLevel":2, "uniqId":"xyz"}'
   return wrk.format("POST", "http://127.0.0.1:9527/post", nil, body) 
end
