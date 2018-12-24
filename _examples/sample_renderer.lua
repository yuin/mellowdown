
local renderer = {}

function renderer:accept(node) 
  return node:fencedCodeBlock():info() == "sample"
end

function renderer:render(writer, node)
  writer:write("write from lua:" .. tostring(node:text()) .. "\n")
end

function renderer:render_header(writer)
  writer:write("<!-- header from lua -->")
end

function renderer:render_footer(writer)
  writer:write("<!-- footer from lua -->")
end

return renderer
