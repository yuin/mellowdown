
local renderer = {}

function renderer:acceptable() 
  return node_fenced_code, "sample"
end

function renderer:render(writer, node, context)
  writer:write("write from lua:" .. tostring(node:text()) .. "\n")
end

function renderer:header(writer, context)
  writer:write("<!-- header from lua -->")
end

function renderer:footer(writer, context)
  writer:write("<!-- footer from lua -->")
end

return renderer
