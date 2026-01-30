package bazi

/*
大運的作用
每一步大運統管十年禍福，在這十年之中，又有十個不同的幹支，這十個不同的幹支對大運來說，有相生、相合、有扶有泄的不同之說。
這十個幹支同樣對四柱中幹支也會出現生扶抑克的影響。所以只推大運還不行，還要加上流年。所謂“流年”，又叫行年太歲。
就是從命造出生的那年開始，不論陰年和陽年，也不管男命女命，一律往下排，如某人出生於甲子年，就從甲子爲一歲流年，乙醜爲二歲流年，丙寅爲三歲流年，依次一直排到壽終。
推排流年有兩種作用：第一種是給求測從生至死的每一年的事情；第二種是答復求測人所問某一年的吉凶。
比如一個人要問45那年的吉凶如何？就從出生年起推出45那年是什麼幹支，再結合命局與大運進行分析，定其吉凶禍福。
*/

// NewDaYun 新大運
func NewDaYun(pSiZhu *TSiZhu, nSex int) *TDaYun {
	p := &TDaYun{}
	p.init(pSiZhu, nSex)
	return p
}

// TDaYun 大運
type TDaYun struct {
	zhuList  [12]*TZhu // 12個大運柱列表
	nAge     [12]int   // 12個大運對應年齡
	isShunNi bool      // 順轉還是逆轉(true 順,  false 逆)
}

func (m *TDaYun) init(pSiZhu *TSiZhu, nSex int) *TDaYun {
	for i := 0; i < 12; i++ {
		m.zhuList[i] = NewZhu() // 新建12個柱
	}

	// 第一判斷年柱的陰陽
	yinyang := pSiZhu.YearZhu().ToYinYang()
	// ! 第二判斷性別的男女

	// 月柱的幹支
	nMonthGanZhi := pSiZhu.MonthZhu().GanZhi().Value()

	for i := 0; i < 12; i++ {
		if yinyang.Value() == nSex {
			m.isShunNi = true
			m.zhuList[i].genBaseGanZhi((nMonthGanZhi + 61 + i) % 60)
		} else {
			m.isShunNi = false
			m.zhuList[i].genBaseGanZhi((nMonthGanZhi + 59 - i) % 60)

		}
	}

	return m
}

// String
func (m *TDaYun) String() string {
	strResult := "大運:\n"

	for i := 0; i < 12; i++ {
		strResult += m.zhuList[i].GanZhi().String() + " "
	}

	return strResult
}

// ShunNi 順逆
func (m *TDaYun) ShunNi() bool {
	return m.isShunNi
}

// Zhu 獲取柱
func (m *TDaYun) Zhu(nIndex int) *TZhu {
	nIndex %= 12
	return m.zhuList[nIndex]
}

// Size 容量就是12
func (m *TDaYun) Size() int {
	return 12
}

// Age 獲取年齡
func (m *TDaYun) Age(nIndex int) int {
	nIndex %= 12
	return m.nAge[nIndex]
}
