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
   parms = idArr[falg%(table.getn(idArr) + 1)]
   --print(falg % table.getn(idArr))
   print(falg % table.getn(idArr), parms)
   falg = falg + 1
   return wrk.format("POST", "http://127.0.0.1:9527/post", nil, parms) 
end
