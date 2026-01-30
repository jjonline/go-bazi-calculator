package bazi

import "fmt"

// NewBazi 新建八字
func NewBazi(pSolarDate *TSolarDate, nSex int) *TBazi {
	//
	pBazi := &TBazi{
		pSolarDate: pSolarDate,
		nSex:       nSex,
	}
	return pBazi.init()
}

// NewBaziFromLunarDate 新建八字 從農歷
func NewBaziFromLunarDate(pLunarDate *TLunarDate, nSex int) *TBazi {
	pBazi := &TBazi{
		pLunarDate: pLunarDate,
		nSex:       nSex,
	}

	return pBazi.init()
}

// GetBazi 舊版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	// 先解決時間問題. 然後開始處理八字問題
	pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	if pSolarDate == nil {
		return nil
	}

	return NewBazi(pSolarDate, nSex)
}

// TBazi 八字大類
type TBazi struct {
	pSolarDate *TSolarDate // 新歷的日期
	pLunarDate *TLunarDate // 農歷日期
	pBaziDate  *TBaziDate  // 八字歷
	pSiZhu     *TSiZhu     // 四柱嗯
	nSex       int         // 性別1男其他女
	pDaYun     *TDaYun     // 大運
	pQiYunDate *TSolarDate // 起運時間XX年XX月開始起運
}

// 八字初始化
func (m *TBazi) init() *TBazi {
	// 1. 新農互轉
	if m.pSolarDate == nil {
		if m.pLunarDate == nil {
			return nil
		}
		// 農轉新
		m.pSolarDate = m.pLunarDate.ToSolarDate()
	} else {
		// 新轉農
		m.pLunarDate = m.pSolarDate.ToLunarDate()
	}

	// 1. 拿到新歷的情況下, 需要計算八字歷
	m.pBaziDate = m.pSolarDate.ToBaziDate()

	// 2. 根據八字歷, 準備計算四柱了
	m.pSiZhu = NewSiZhu(m.pSolarDate, m.pBaziDate)

	// 3. 計算大運
	m.pDaYun = NewDaYun(m.pSiZhu, m.nSex)

	// 4. 計算起運時間
	m.pQiYunDate = NewQiYun(m.pDaYun.ShunNi(), m.pBaziDate.PreviousJie().ToSolarDate(), m.pBaziDate.NextJie().ToSolarDate(), m.pSolarDate)

	// 5. 起運時間融入到大運中
	nAge := m.QiYunDate().Year() - m.Date().Year()
	for i := 0; i < 12; i++ {
		m.pDaYun.nAge[i] = nAge + 10*i
	}

	return m
}

// String 打印用
func (m *TBazi) String() string {
	return fmt.Sprintf("%v\n %v\n %v\n%v\n%v \n起運時間%v", m.pSolarDate, m.pLunarDate, m.pBaziDate, m.pSiZhu, m.pDaYun, m.pQiYunDate)
}

// SiZhu 四柱
func (m *TBazi) SiZhu() *TSiZhu {
	return m.pSiZhu
}

// Date 獲取日期， 默認就是新歷
func (m *TBazi) Date() *TSolarDate {
	return m.pSolarDate
}

// SolarData 獲取新歷日期
func (m *TBazi) SolarData() *TSolarDate {
	return m.Date()
}

// LunarDate 獲取農歷日期
func (m *TBazi) LunarDate() *TLunarDate {
	return m.pLunarDate
}

// DaYun 獲取大運
func (m *TBazi) DaYun() *TDaYun {
	return m.pDaYun
}

// QiYunDate 起運時間
func (m *TBazi) QiYunDate() *TSolarDate {
	return m.pQiYunDate
}
