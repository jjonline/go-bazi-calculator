package bazi

// 地支藏幹表
var cangganlist = [12][3]int{
	{9, -1, -1}, // 子水 藏幹 癸水。
	{5, 9, 7},   // 醜土 藏幹 己土、癸水、辛金。
	{0, 2, 4},   // 寅木 藏幹 甲木、丙火、戊土。
	{1, -1, -1}, // 卯木 藏幹 乙木。
	{4, 1, 9},   // 辰土 藏幹 戊土、乙木、癸水。
	{2, 4, 6},   // 巳火 藏幹 丙火、戊土、庚金。
	{3, 5, -1},  // 午火 藏幹 丁火、己土。
	{5, 1, 3},   // 未土 藏幹 己土、乙木、丁火。
	{6, 4, 8},   // 申金 藏幹 庚金、戊土、壬水。
	{7, -1, -1}, // 酉金 藏幹 辛金。
	{4, 7, 3},   // 戌土 藏幹 戊土、辛金、丁火。
	{8, 0, -1}}  // 亥水 藏幹 壬水、甲木。

// NewCangGan 新建藏幹
func NewCangGan(nDayGan int, pZhi *TZhi) *TCangGan {
	pCangGan := &TCangGan{
		nDayGan: nDayGan,
	}

	pCangGan.init(nDayGan, pZhi)

	return pCangGan
}

// TCangGan 藏幹
type TCangGan struct {
	cangGanList []*TGan
	shishenList []*TShiShen
	nDayGan     int // 記錄用日幹
}

func (m *TCangGan) init(nDayGan int, pZhi *TZhi) {
	nZhi := pZhi.Value()
	for i := 0; i < 3; i++ {
		// 判斷藏幹有效性
		if cangganlist[nZhi][i] >= 0 {
			// 添加藏幹
			pGan := NewGan(cangganlist[nZhi][i])
			pShiShen := NewShiShenFromGan(nDayGan, pGan)
			m.cangGanList = append(m.cangGanList, pGan)
			m.shishenList = append(m.shishenList, pShiShen)
			// 添加十神
		} else {
			break
		}
	}
}

// Size 內容
func (m *TCangGan) Size() int {
	return len(m.cangGanList)
}

// Gan 獲取具體某個索引
func (m *TCangGan) Gan(nIdx int) *TGan {
	if nIdx < 0 {
		return nil
	}
	if nIdx >= m.Size() {
		return nil
	}
	return m.cangGanList[nIdx]
}

// ShiShen 十神
func (m *TCangGan) ShiShen(nIdx int) *TShiShen {
	if nIdx < 0 {
		return nil
	}
	if nIdx >= m.Size() {
		return nil
	}
	return m.shishenList[nIdx]
}

func (m *TCangGan) String() string {
	strResult := ""

	for i := 0; i < m.Size(); i++ {
		strResult += m.Gan(i).String() + "[" + m.ShiShen(i).String() + "]"
	}

	return strResult
}
