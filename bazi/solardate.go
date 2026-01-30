package bazi

// 新歷日期時間， 采用標準的日歷

import (
	"fmt"
)

// NewSolarDate 創建一個新歷時間
func NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond int) *TSolarDate {
	// 把具體時間實例化出來
	pDate := &TSolarDate{
		nYear:   nYear,   // 年
		nMonth:  nMonth,  // 月
		nDay:    nDay,    // 日
		nHour:   nHour,   // 時
		nMinute: nMinute, // 分
		nSecond: nSecond, // 秒
	}

	if !pDate.GetDateIsValid(nYear, nMonth, nDay) {
		fmt.Println("無效的日期", nYear, nMonth, nDay)
		return nil
	}

	// 檢查時間是否合法, 傳入一個大值HOUR導致崩潰的BUG // fix chadwi https://github.com/warrially/BaziGo/issues/3
	if !pDate.GetTimeIsValid(nHour, nMinute, nSecond) {
		fmt.Println("無效的時間", nHour, nMinute, nSecond)
		return nil
	}
	// 計算64位時間戳值
	// pDate.Get64TimeStamp()

	return pDate
}

// NewSolarDateFrom64TimeStamp 從64位時間戳反推日期
func NewSolarDateFrom64TimeStamp(nTimeStamp int64) *TSolarDate {
	pDate := &TSolarDate{}
	// 計算出年份
	pDate.GetYearFrom64TimeStamp(nTimeStamp)
	// 計算月份
	pDate.GetMonthFrom64TimeStamp(nTimeStamp)
	// 計算其他參數
	pDate.GetDayTimeFrom64TimeStamp(nTimeStamp)

	return pDate
}

// TSolarDate 日期
type TSolarDate struct {
	nYear   int // 年
	nMonth  int // 月
	nDay    int // 日
	nHour   int // 時
	nMinute int // 分
	nSecond int // 秒
}

// GetDiffSeconds 獲取兩個日期之間相差的秒數
func (m *TSolarDate) GetDiffSeconds(other *TSolarDate) int64 {
	return other.Get64TimeStamp() - m.Get64TimeStamp()
}

// Get64TimeStamp 生成64位時間戳
func (m *TSolarDate) Get64TimeStamp() int64 {
	nAllDays := m.GetAllDays() // 先獲取公元原點的日數
	nResult := int64(nAllDays)
	nResult *= 24 * 60 * 60 // 天數換成秒

	//再計算出秒數
	nResult += int64(m.nHour) * 60 * 60
	nResult += int64(m.nMinute) * 60
	nResult += int64(m.nSecond)

	return nResult
}

// GetYearFrom64TimeStamp 從64位時間戳反推年
func (m *TSolarDate) GetYearFrom64TimeStamp(nTimeStamp int64) *TSolarDate {
	// 準備進行二分法
	nLow := 0
	nHigh := 3001

	for {
		nMid := (nLow + nHigh) / 2

		// 拿到中間年的數據
		v := NewSolarDate(nMid, 1, 1, 0, 0, 0).Get64TimeStamp()

		if v <= nTimeStamp {
			nLow = nMid
		} else {
			nHigh = nMid
		}

		if nHigh == nLow+1 {
			break
		}
	}
	m.nYear = nLow
	return m
}

// GetMonthFrom64TimeStamp 從64位時間戳反推月,
func (m *TSolarDate) GetMonthFrom64TimeStamp(nTimeStamp int64) {
	// 這里開始特殊處理
	for i := 1; i <= 11; i++ {
		if nTimeStamp < NewSolarDate(m.nYear, i+1, 1, 0, 0, 0).Get64TimeStamp() {
			m.nMonth = i
			return
		}
	}
	m.nMonth = 12
}

// GetDayTimeFrom64TimeStamp 從64位時間戳反推其他參數
func (m *TSolarDate) GetDayTimeFrom64TimeStamp(nTimeStamp int64) {
	nTimeStamp -= NewSolarDate(m.nYear, m.nMonth, 1, 0, 0, 0).Get64TimeStamp()

	// 計算日
	m.nDay = int(nTimeStamp / (24 * 60 * 60))
	// 扣掉日
	nTimeStamp -= int64(m.nDay) * 24 * 60 * 60

	m.nDay++ // 因爲每個月的天數是從1開始的, 所以這里需要補1天
	if m.nYear == 1582 && m.nMonth == 10 && m.nDay >= 5 {
		m.nDay += 10 // 1582 年需要補10天
	}
	m.nHour = int(nTimeStamp / (60 * 60))
	nTimeStamp -= int64(m.nHour) * 60 * 60
	m.nMinute = int(nTimeStamp / 60)
	nTimeStamp -= int64(m.nMinute) * 60
	m.nSecond = int(nTimeStamp)
}

// GetMonthDays 取本月天數，不考慮 1582 年 10 月的特殊情況
func (m *TSolarDate) GetMonthDays(nYear, nMonth int) int {
	switch nMonth {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2: // 閏年
		if m.GetIsLeapYear(nYear) {
			return 29
		}
		return 28
	}
	return 0
}

// GetIsLeapYear 返回某公歷是否閏年
func (m *TSolarDate) GetIsLeapYear(nYear int) bool {
	if m.GetCalendarType(nYear, 1, 1) == ctGregorian {
		return (nYear%4 == 0) && ((nYear%100 != 0) || (nYear%400 == 0))
	} else if nYear >= 0 {
		return nYear%4 == 0
	} else { // 需要獨立判斷公元前的原因是沒有公元 0 年
		return (nYear-3)%4 == 0
	}
}

const (
	ctInvalid   = iota //非法，
	ctJulian           //儒略，
	ctGregorian        //格利高里
)

// GetCalendarType 根據公歷日期判斷當時歷法
func (m *TSolarDate) GetCalendarType(nYear, nMonth, nDay int) int {
	if !m.GetDateIsValid(nYear, nMonth, nDay) {
		return ctInvalid
	}
	if nYear > 1582 {
		return ctGregorian
	} else if nYear < 1582 {
		return ctJulian
	} else if nMonth < 10 {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 4) {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 14) {
		return ctInvalid
	} else {
		return ctGregorian
	}
	// 在現在通行的歷法記載上，全世界居然有十天沒有任何人出生過，也沒有任何人死亡過，也沒有發生過大大小小值得紀念的人或事。這就是1582年10月5日至10月14日。格里奧，提出了公歷歷法。這個歷法被羅馬教皇格里高利十三世采納了。那麼誤差的十天怎麼辦？羅馬教皇格里高利十三世下令，把1582年10月4日的後一天改爲10月15日，這樣誤差的十天沒有了，歷史上也就無影無蹤地消失了十天，當然史書上也就沒有這十天的記載了。“格里高利公歷”一直沿用到今天。
}

// GetDateIsValid 返回公歷日期是否合法
func (m *TSolarDate) GetDateIsValid(nYear, nMonth, nDay int) bool {
	// 沒有公元0年
	if nYear == 0 {
		return false
	}

	// 1月開始, 12月結束
	if nMonth < 1 || nMonth > 12 {
		return false
	}

	// 1號開始, 獲取每個月有多少天結束
	if nDay < 1 || nDay > m.GetMonthDays(nYear, nMonth) {
		return false
	}

	// 1582 年的特殊情況
	if nYear != 1582 {
		return true
	}
	if nMonth != 10 {
		return true
	}
	//
	if nDay < 5 || nDay > 14 {
		return true
	}

	return false
}

// GetTimeIsValid 檢查時間是否合法
func (m *TSolarDate) GetTimeIsValid(nHour, nMinute, nSecond int) bool {
	if nHour < 0 || nHour > 23 {
		return false
	}

	if nMinute < 0 || nMinute > 59 {
		return false
	}

	if nSecond < 0 || nSecond > 59 {
		return false
	}
	return true
}

// GetAllDays 獲得距公元原點的日數 這里是公歷的年月日
func (m *TSolarDate) GetAllDays() int {
	nYear := m.Year()
	nMonth := m.Month()
	nDay := m.Day()
	if m.GetDateIsValid(nYear, nMonth, nDay) {
		return m.GetBasicDays(nYear, nMonth, nDay) + m.GetLeapDays(nYear, nMonth, nDay)
	}
	return 0
}

// GetBasicDays 獲取基本數據
func (m *TSolarDate) GetBasicDays(nYear, nMonth, nDay int) int {
	if !m.GetDateIsValid(nYear, nMonth, nDay) {
		return 0
	}

	var Result int

	// 去掉公元0年
	if nYear > 0 {
		Result = (nYear - 1) * 365
	} else {
		Result = nYear * 365
	}

	// 加上月天數
	for i := 1; i < nMonth; i++ {
		Result += m.GetMonthDays(nYear, i)
	}

	// 加上日天數
	Result += nDay
	// 返回基礎天數
	return Result
}

// GetLeapDays 獲取閏年天數
func (m *TSolarDate) GetLeapDays(nYear, nMonth, nDay int) int {
	if !m.GetDateIsValid(nYear, nMonth, nDay) {
		return 0
	}
	var Result int

	if nYear >= 0 {
		// 公元後
		if m.GetCalendarType(nYear, nMonth, nDay) < ctGregorian {
			Result = 0
		} else {
			// 1582.10.5/15 前的 Julian 歷只有四年一閏，歷法此日後調整爲 Gregorian 歷
			Result = 10 // 被 Gregory 刪去的 10 天

			// 修正算法簡化版，從 1701 年的 11 起
			if nYear > 1700 {
				// 每一世紀累加一
				Result += (1 + ((nYear - 1701) / 100))
				// 但 400 整除的世紀不加
				Result -= ((nYear - 1601) / 400)
			}
		}
		Result = ((nYear - 1) / 4) - Result // 4 年一閏數
	} else {
		// 公元前
		Result = -((-nYear + 3) / 4)
	}
	return Result
}

func (m *TSolarDate) String() string {
	return fmt.Sprintf("新歷: %d 年 %02d 月 %02d 日 %02d:%02d:%02d",
		m.nYear, m.nMonth, m.nDay, m.nHour, m.nMinute, m.nSecond)
}

// ToBaziDate 轉成八字日期
func (m *TSolarDate) ToBaziDate() *TBaziDate {
	return NewBaziDate(m)
}

// Year 年
func (m *TSolarDate) Year() int {
	return m.nYear
}

// Month 月
func (m *TSolarDate) Month() int {
	return m.nMonth
}

// Day 日
func (m *TSolarDate) Day() int {
	return m.nDay
}

// Hour 時
func (m *TSolarDate) Hour() int {
	return m.nHour
}

// Minute 分
func (m *TSolarDate) Minute() int {
	return m.nMinute
}

// Second 秒
func (m *TSolarDate) Second() int {
	return m.nSecond
}

// ToLunarDate 轉成農歷年
func (m *TSolarDate) ToLunarDate() *TLunarDate {
	return NewLunarDateFrom64TimeStamp(m.Get64TimeStamp())
}
