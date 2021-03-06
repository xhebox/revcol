function utf8check(s)
  local r = {}

	local e = string.len(s)
	local c = 1
	local p = 1

	while p <= e do
		local ch1 = string.byte(s, p)
		local ch2, ch3, ch4

		if ch1 < 0x80 then
			-- print(1)
			r[c] = ch1
			p = p + 1
		elseif ch1 < 0xC0 then
			-- print(1.5)
			r[c] = 0x49
			p = p + 1
		elseif ch1 < 0xE0 then
			ch2 = string.byte(s, p+1)
			p = p + 2

			-- print(2, (ch1-0xC0)*0x40 + (ch2-0x80), 0x80, 0x7FF)
			if ch2 == nil or ch2 >= 0xC0 or ch2 < 0x80 or
			((ch1-0xC0) * 0x4 + (ch2-0x80) / 0xF) < 0x8 then
				r[c] = 0x49
			end

			if r[c] ~= 0x49 then
				r[c] = ch1
				r[c+1] = ch2
				c = c + 1
			end
		elseif ch1 < 0xF0 then
			ch2 = string.byte(s, p+1)
			ch3 = string.byte(s, p+2)
			p = p + 3

			-- print(3, (ch1-0xE0)*0x1000 + (ch2-0x80)*0x40 + (ch3-0x80), 0x800, 0xFFFF)
			if ch2 == nil or ch2 >= 0xC0 or ch2 < 0x80 or
			ch3 == nil or ch3 >= 0xC0 or ch3 < 0x80 or
			((ch1-0xE0) * 0xF + (ch2-0x80) / 0x4) < 0x8 then
				r[c] = 0x49
			end

			if r[c] ~= 0x49 then
				r[c] = ch1
				r[c+1] = ch2
				r[c+2] = ch3
				c = c + 2
			end
		elseif ch1 < 0xF8 then
			ch2 = string.byte(s, p+1)
			ch3 = string.byte(s, p+2)
			ch4 = string.byte(s, p+3)
			p = p + 4

			-- print(4, (ch1-0xF0)*0x40000 + (ch2-0x80)*0x1000 + (ch3-0x80)*0x40 + (ch4-0x80), 0x10000, 0x10FFFF)
			if ch2 == nil or ch2 >= 0xC0 or ch2 < 0x80 or
			ch3 == nil or ch3 >= 0xC0 or ch3 < 0x80 or
			ch4 == nil or ch4 >= 0xC0 or ch4 < 0x80 or
			((ch2-0x80) / 0xF) < 0x1 then
				r[c] = 0x49
			end

			if r[c] ~= 0x49 then
				r[c] = ch1
				r[c+1] = ch2
				r[c+2] = ch3
				r[c+3] = ch4
				c = c + 3
			end
		else
			-- print(5)
			r[c] = 0x49
			p = p + 1
		end

		c = c + 1
	end

	return string.char(unpack(r))
end

print(utf8check("ะแธโฒค๐๐ค๐๊ ๊งศ๐๐ฅ๊ก๐๐ตวฌ๐ฟ๐ล๐๐ฏ๐ด๐๐๊ซลธ๐กแบฃ๐ขฦ๐ผแธแบฟแตฎโ๐แฅ๐ะบฮนแนีคโฑบ๐๐ฒ๐ฃ๐ลง๐ขแนฝแบ๐แงลพ1234567890!@#$%^&*()-_=+[{]};:',<.>/?~ัฆ๐ฑฦแฮฃโฑิาคูก๐ะ๐๐ฦศ๐ธ๐แน๐ขแนฎแนบฦฒแ๊ซ๐๐ญ๐ถแรงแซ๐๐ฟ๐แธง๐๐ฃาษญแธฟ๐๐จ๐๐ขแน๐ผัรบ๐ณแบโคฌ๐ฒ๐1234567890!@#$%^&*()-_=+[{]};:',<.>/?~๐ ฮ๐๐๐ด๐ฤขศแป๐ต๊ส๐ผโเงฆ๐ธ๐คี๊ขแนฐวโค๐โฒฌ๐๐๐ข๐ค "))
