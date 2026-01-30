package bazi

// 五行

// GetWuXingFromNumber  從數字獲得五行名, 0-4
func GetWuXingFromNumber(nValue int) string {
	// {* 五行字符串，以通常的金木水火土爲順序 }
	// 這里沒用五行相生或者相克來排列
	switch nValue {
	case 0:
		return "金"
	case 1:
		return "木"
	case 2:
		return "水"
	case 3:
		return "火"
	case 4:
		return "土"
	}
	return ""
}

// GetWuXingColorFromNumber 獲取五行的顏色
func GetWuXingColorFromNumber(nValue int) string {
	// {* 五行字符串，以通常的金木水火土爲順序 }
	// 這里沒用五行相生或者相克來排列
	switch nValue {
	case 0:
		return "gold"
	case 1:
		return "green"
	case 2:
		return "black"
	case 3:
		return "red"
	case 4:
		return "brown"
	}
	return ""
}

// GetWuXingFromGan 獲得某幹的五行，0-4 對應 金木水火土
// 甲乙爲木，丙丁爲火，戊己爲土，庚辛爲金，壬癸爲水，
func GetWuXingFromGan(pGan *TGan) *TWuXing {
	return pGan.ToWuXing()
}

// NewWuXing 創建五行
func NewWuXing(nValue int) *TWuXing {
	nValue %= 5
	wuxing := TWuXing(nValue)
	return &wuxing
}

// TWuXing 五行
type TWuXing int

// ToString 轉換成可閱讀的字符串
func (m *TWuXing) ToString() string {
	return m.String()
}

// ToInt 轉換成int
func (m *TWuXing) ToInt() int {
	return m.Value()
}

// Value 轉換成int
func (m *TWuXing) Value() int {
	return (int)(*m)
}

// String 轉換成可閱讀的字符串
func (m *TWuXing) String() string {
	return GetWuXingFromNumber(m.Value())
}

// Color 五行顏色
func (m *TWuXing) Color() string {
	return GetWuXingColorFromNumber(m.Value())
}
