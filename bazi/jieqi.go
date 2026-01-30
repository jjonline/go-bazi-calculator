package bazi

// TJieQi 節氣類
type TJieQi int

// GetJieQiFromNumber 從數字獲得節氣名, 0-23
func GetJieQiFromNumber(nValue int) string {
	switch nValue {
	case 0: // 節氣  Beginning of Spring   0
		return "立春"
	case 1: // 中氣  Rain Water            1
		return "雨水"
	case 2: // 節氣  Waking of Insects     2
		return "驚蟄"
	case 3: // 中氣  March Equinox         3
		return "春分"
	case 4: // 節氣  Pure Brightness       4
		return "清明"
	case 5: // 中氣  Grain Rain            5
		return "谷雨"
	case 6: // 節氣  Beginning of Summer   6
		return "立夏"
	case 7: // 中氣  Grain Full            7
		return "小滿"
	case 8: // 節氣  Grain in Ear          8
		return "芒種"
	case 9: // 中氣  Summer Solstice       9
		return "夏至"
	case 10: // 節氣  Slight Heat           10
		return "小暑"
	case 11: // 中氣  Great Heat            11
		return "大暑"
	case 12: // 節氣  Beginning of Autumn   12
		return "立秋"
	case 13: // 中氣  Limit of Heat         13
		return "處暑"
	case 14: // 節氣  White Dew             14
		return "白露"
	case 15: // 中氣  September Equinox     15
		return "秋分"
	case 16: // 節氣  Cold Dew              16
		return "寒露"
	case 17: // 中氣  Descent of Frost      17
		return "霜降"
	case 18: // 節氣  Beginning of Winter   18
		return "立冬"
	case 19: // 中氣  Slight Snow           19
		return "小雪"
	case 20: // 節氣  Great Snow            20
		return "大雪"
	case 21: // 中氣  Winter Solstice       21
		return "冬至"
	case 22: // 節氣  Slight Cold           22   	，這是一公歷年中的第一個節氣
		return "小寒"
	case 23: // 中氣  Great Cold            23
		return "大寒"
	}
	return ""
}

// IsJie 節氣是否是節,   節氣分成節和氣,
func (m *TJieQi) IsJie() bool {
	n := m.Value()
	return n%2 == 0
}

// ToString 轉換成可閱讀的字符串
func (m *TJieQi) ToString() string {
	return m.String()
}

// ToInt 轉換成int
func (m *TJieQi) ToInt() int {
	return m.Value()
}

// Month 節氣月份
func (m *TJieQi) Month() int {
	return m.ToMonth()
}

// ToMonth 轉成節氣月
func (m *TJieQi) ToMonth() int {
	// 節氣0 是立春 是1月
	return m.Value()/2 + 1
}

// Value 轉換成int
func (m *TJieQi) Value() int {
	return (int)(*m)
}

// String 轉換成可閱讀的字符串
func (m *TJieQi) String() string {
	return GetJieQiFromNumber(m.Value())
}
