package bazi

// 天幹
// 甲木、乙木、丙火、丁火、戊土、己土、庚金、辛金、壬水、癸水，其中甲 丙 戊 庚 壬爲陽性，乙丁己辛癸爲陰性
// 詩曰：
// 春季甲乙東方木，夏季丙丁南方火；
// 秋季庚辛西方金，冬季壬癸北方水；
// 戊己中央四季土。

// GetTianGanFromNumber 從數字獲得天幹名, 0-9
func GetTianGanFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "甲"
	case 1:
		return "乙"
	case 2:
		return "丙"
	case 3:
		return "丁"
	case 4:
		return "戊"
	case 5:
		return "己"
	case 6:
		return "庚"
	case 7:
		return "辛"
	case 8:
		return "壬"
	case 9:
		return "癸"
	}

	return ""
}

// NewGan 創建天幹
func NewGan(nValue int) *TGan {
	nValue %= 10
	pGan := TGan(nValue)
	return &pGan
}

// TGan 天幹
type TGan int

// ToString 轉換成可閱讀的字符串
func (m *TGan) ToString() string {
	return m.String()
}

// ToInt 轉換成int
func (m *TGan) ToInt() int {
	return m.Value()
}

// ToWuXing 天幹轉化成五行
func (m *TGan) ToWuXing() *TWuXing {
	// todo
	// 甲木、乙木、丙火、丁火、戊土、己土、庚金、辛金、壬水、癸水，其中甲 丙 戊 庚 壬爲陽性，乙丁己辛癸爲陰性
	switch m.Value() {
	case 0, 1:
		return NewWuXing(1) // 甲 陽木 乙 陰木
	case 2, 3:
		return NewWuXing(3) // 丙 陽火 丁 陰火
	case 4, 5:
		return NewWuXing(4) // 戊 陽土 己 陰土
	case 6, 7:
		return NewWuXing(0) // 庚 陽金 辛 陰金
	case 8, 9:
		return NewWuXing(2) // 壬 陽水 癸 陰水
	}
	return nil
}

// Value 轉換成int
func (m *TGan) Value() int {
	return (int)(*m)
}

// String 轉換成可閱讀的字符串
func (m *TGan) String() string {
	return GetTianGanFromNumber(m.Value())
}
