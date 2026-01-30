package bazi

// 地支

// GetDiZhiFromNumber 從數字獲得地支名, 0-9
func GetDiZhiFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "子"
	case 1:
		return "醜"
	case 2:
		return "寅"
	case 3:
		return "卯"
	case 4:
		return "辰"
	case 5:
		return "巳"
	case 6:
		return "午"
	case 7:
		return "未"
	case 8:
		return "申"
	case 9:
		return "酉"
	case 10:
		return "戌"
	case 11:
		return "亥"
	}

	return ""
}

// NewZhi 創建地支
func NewZhi(nValue int) *TZhi {
	nValue %= 12
	Zhi := TZhi(nValue)
	return &Zhi
}

// TZhi 地支
type TZhi int

// ToString 轉換成可閱讀的字符串
func (m *TZhi) ToString() string {
	return m.String()
}

// ToInt 轉換成int
func (m *TZhi) ToInt() int {
	return m.Value()
}

// ToWuXing 地支轉化成五行
func (m *TZhi) ToWuXing() *TWuXing {
	switch m.Value() {
	case 8, 9:
		return NewWuXing(0)
	case 2, 3:
		return NewWuXing(1)
	case 0, 11:
		return NewWuXing(2)
	case 5, 6:
		return NewWuXing(3)
	case 1, 4, 7, 10:
		return NewWuXing(4) // 土
	}
	return nil
}

// Value 轉換成int
func (m *TZhi) Value() int {
	return (int)(*m)
}

// String 轉換成可閱讀的字符串
func (m *TZhi) String() string {
	return GetDiZhiFromNumber(m.Value())
}
