
local renderer = {}

function renderer:accept(node) 
  return node:fencedCodeBlock():info() == "sample"
end

function renderer:render(writer, node)
  writer:write("write from lua:" .. tostring(node:text()) .. "\n")
end

function renderer:header(writer)
  writer:write("<!-- header from lua -->")
end

function renderer:footer(writer)
  writer:write("<!-- footer from lua -->")
end

return renderer
