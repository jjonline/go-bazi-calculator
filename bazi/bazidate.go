package bazi

import (
	"fmt"
)

// NewBaziDate 從新歷轉成八字歷
func NewBaziDate(pSolarDate *TSolarDate) *TBaziDate {
	p := &TBaziDate{}
	p.init(pSolarDate)
	return p
}

// TBaziDate 八字歷法
// 八字歷法的年  和 新歷的 和 農歷的都不一樣. 八字歷法是按照立春爲1年. 然後每個節氣爲月
type TBaziDate struct {
	nYear  int // 年. 立春
	nMonth int // 月.
	nDay   int // 天
	nHour  int // xiaohsi

	pJieQi       *TJieQi     // 節氣名稱
	pPreviousJie *TJieQiDate // 上一個節(氣)
	pNextJie     *TJieQiDate // 下一個節(氣)
}

func (m *TBaziDate) init(pSolarDate *TSolarDate) *TBaziDate {
	m.nYear = GetLiChunYear(pSolarDate)                   // 拿到八字年, 根據立春來的
	m.pPreviousJie, m.pNextJie = GetJieQiDate(pSolarDate) // 拿到前後兩個的日期
	// 節氣
	nJieQi := m.pPreviousJie.JieQi
	m.pJieQi = &nJieQi
	// 月
	m.nMonth = m.pJieQi.Month()
	return m
}

func (m *TBaziDate) String() string {
	return fmt.Sprintf("八字歷: %4d 年 %02d 月 \n上一個:%v\n下一個:%v",
		m.nYear, m.nMonth, m.pPreviousJie, m.pNextJie)
}

// Year  年. 立春
func (m *TBaziDate) Year() int {
	return m.nYear
}

// Month  月.
func (m *TBaziDate) Month() int {
	return m.nMonth
}

// Day  天
func (m *TBaziDate) Day() int {
	return m.nDay
}

// Hour 小時
func (m *TBaziDate) Hour() int {
	return m.nHour
}

// PreviousJie 上一個節氣
func (m *TBaziDate) PreviousJie() *TJieQiDate {
	return m.pPreviousJie
}

// NextJie 下一個節氣
func (m *TBaziDate) NextJie() *TJieQiDate {
	return m.pNextJie
}
