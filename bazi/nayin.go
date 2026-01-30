package bazi

//  {* 納音五行，與相鄰一對六十幹支對應}
// 甲子乙醜海中金 丙寅丁卯爐中火 戊辰己巳大林木
// 庚午辛未路旁土 壬申癸酉劍鋒金 甲戌乙亥山頭火
// 丙子丁醜澗下水 戊寅己卯城頭土 庚辰辛巳白蠟金
// 壬午癸未楊柳木 甲申乙酉井泉水 丙戌丁亥屋上土
// 戊子己醜霹靂火 庚寅辛卯松柏木 壬辰癸巳長流水
// 甲午乙未砂中金 丙申丁酉山下火 戊戌己亥平地木
// 庚子辛醜壁上土 壬寅癸卯金箔金 甲辰乙巳覆燈火
// 丙午丁未天河水 戊申己酉大驛土 庚戌辛亥釵釧金
// 壬子癸醜桑柘木 甲寅乙卯大溪水 丙辰丁巳砂中土
// 戊午己未天上火 庚申辛酉石榴木 壬戌癸亥大海水

// GetNaYinFromNumber 從數字獲得納音名, 0-29
func GetNaYinFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "海中金"
	case 1:
		return "爐中火"
	case 2:
		return "大林木"

	case 3:
		return "路旁土"
	case 4:
		return "劍鋒金"
	case 5:
		return "山頭火"

	case 6:
		return "澗下水"
	case 7:
		return "城牆土"
	case 8:
		return "白蠟金"

	case 9:
		return "楊柳木"
	case 10:
		return "泉中水"
	case 11:
		return "屋上土"

	case 12:
		return "霹雷火"
	case 13:
		return "松柏木"
	case 14:
		return "長流水"

	case 15:
		return "沙中金"
	case 16:
		return "山下火"
	case 17:
		return "平地木"

	case 18:
		return "壁上土"
	case 19:
		return "金箔金"
	case 20:
		return "佛燈火"

	case 21:
		return "天河水"
	case 22:
		return "大驛土"
	case 23:
		return "釵釧金"

	case 24:
		return "桑柘木"
	case 25:
		return "大溪水"
	case 26:
		return "沙中土"

	case 27:
		return "天上火"
	case 28:
		return "石榴木"
	case 29:
		return "大海水"
	}

	return ""
}

// NewNaYin 納音
func NewNaYin(nValue int) *TNaYin {
	nValue %= 30

	NaYin := TNaYin(nValue)
	return &NaYin
}

// TNaYin 納音
type TNaYin int

// Value 轉換成int
func (m *TNaYin) Value() int {
	return (int)(*m)
}

func (m *TNaYin) ToInt() int {
	return m.Value()
}

// String 轉換成可閱讀的字符串
func (m *TNaYin) String() string {
	return GetNaYinFromNumber(m.Value())
}

func (m *TNaYin) ToString() string {
	return m.String()
}
