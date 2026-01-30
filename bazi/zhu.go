package bazi

// 這里是柱， 四柱之一

import "fmt"

// TZhu 柱
type TZhu struct {
	pGanZhi  *TGanZhi  // 幹支
	pGan     *TGan     // 天幹
	pZhi     *TZhi     // 地支
	pCangGan *TCangGan // 藏幹
	pShiShen *TShiShen // 十神
	nDayGan  int       // 日幹值
}

// NewZhu 新建柱子
func NewZhu() *TZhu {
	return &TZhu{}
}

// String 打印
func (m *TZhu) String() string {
	return fmt.Sprintf("%v", m.pGanZhi)
}

// 設置日幹值
func (m *TZhu) setDayGan(nDayGan int) *TZhu {
	m.nDayGan = nDayGan
	return m
}

// 生成藏幹
func (m *TZhu) genCangGan() {
	// 生成藏幹數據
	if m.pZhi != nil {
		m.pCangGan = NewCangGan(m.nDayGan, m.pZhi)
	}
}

// 生成十神
func (m *TZhu) genShiShen() {
	m.pShiShen = NewShiShenFromGan(m.nDayGan, m.pGan)
}

//
func (m *TZhu) genBaseGanZhi(nGanZhi int) *TZhu {
	// 直接設置成品幹支
	m.pGanZhi = NewGanZhi(nGanZhi)
	// 拆分幹支
	// 獲得八字年的幹0-9 對應 甲到癸
	// 獲得八字年的支0-11 對應 子到亥
	m.pGan, m.pZhi = m.pGanZhi.ExtractGanZhi()

	return m
}

// genYearGanZhi 生成年幹支
func (m *TZhu) genYearGanZhi(nYear int) *TZhu {
	// 通過年獲取幹支
	// 獲得八字年的幹支，0-59 對應 甲子到癸亥
	m.pGanZhi = NewGanZhiFromYear(nYear)
	// 拆分幹支
	// 獲得八字年的幹0-9 對應 甲到癸
	// 獲得八字年的支0-11 對應 子到亥
	m.pGan, m.pZhi = m.pGanZhi.ExtractGanZhi()

	// 在這里計算藏幹
	m.genCangGan()
	m.genShiShen()
	return m
}

// genMonthGanZhi 生成月幹支
func (m *TZhu) genMonthGanZhi(nMonth int, nYearGan int) *TZhu {
	// 根據口訣從本年幹數計算本年首月的幹數
	switch nYearGan {
	case 0, 5:
		// 甲己 丙佐首
		nYearGan = 2
	case 1, 6:
		// 乙庚 戊爲頭
		nYearGan = 4
	case 2, 7:
		// 丙辛 尋庚起
		nYearGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nYearGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nYearGan = 0
	}

	// 計算本月幹數
	nYearGan += ((nMonth - 1) % 10)

	// 拆幹
	m.pGan = NewGan(nYearGan % 10)
	m.pZhi = NewZhi((nMonth - 1 + 2) % 12)

	// 組合幹支
	m.pGanZhi = CombineGanZhi(m.pGan, m.pZhi)
	// 在這里計算藏幹
	m.genCangGan()
	m.genShiShen()
	return m
}

// genDayGanZhi 生成日幹支
func (m *TZhu) genDayGanZhi(nAllDays int) *TZhu {

	// 通過總天數來獲取
	// 獲得八字年的幹支，0-59 對應 甲子到癸亥
	m.pGanZhi = NewGanZhiFromDay(nAllDays)
	// 拆分幹支
	// 獲得八字年的幹0-9 對應 甲到癸
	// 獲得八字年的支0-11 對應 子到亥
	m.pGan, m.pZhi = m.pGanZhi.ExtractGanZhi()

	// 直接保存日幹
	m.setDayGan(m.pGan.Value())

	// 在這里計算藏幹
	m.genCangGan()
	m.genShiShen()
	return m
}

// genHourGanZhi 生成時幹支
func (m *TZhu) genHourGanZhi(nHour int) *TZhu {
	// 取出日幹
	nGan := m.nDayGan

	// 24小時校驗
	nHour %= 24
	if nHour < 0 {
		nHour += 24
	}

	nZhi := 0
	if nHour == 23 {
		// 次日子時
		nGan = (nGan + 1) % 10
	} else {
		nZhi = (nHour + 1) / 2
	}

	// Gan 此時是本日幹數，根據規則換算成本日首時辰幹數
	if nGan >= 5 {
		nGan -= 5
	}

	// 計算此時辰幹數
	nGan = (2*nGan + nZhi) % 10

	m.pGan = NewGan(nGan)
	m.pZhi = NewZhi(nZhi)

	// 組合幹支
	m.pGanZhi = CombineGanZhi(m.pGan, m.pZhi)

	// 在這里計算藏幹
	m.genCangGan()
	m.genShiShen()
	return m
}

// Gan 獲取幹
func (m *TZhu) Gan() *TGan {
	return m.pGan
}

// Zhi 獲取支
func (m *TZhu) Zhi() *TZhi {
	return m.pZhi
}

// GanZhi 獲取幹支
func (m *TZhu) GanZhi() *TGanZhi {
	return m.pGanZhi
}

// ToYinYang 從柱里獲取陰陽 (陰 == 0,  陽 == 1)
func (m *TZhu) ToYinYang() *TYinYang {
	return NewYinYangFromZhu(m)
}

// CangGan 獲取藏幹
func (m *TZhu) CangGan() *TCangGan {
	return m.pCangGan
}

// ShiShen 獲取十神
func (m *TZhu) ShiShen() *TShiShen {
	return m.pShiShen
}
