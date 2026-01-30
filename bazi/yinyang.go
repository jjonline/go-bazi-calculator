package bazi

// 陰陽
// 從柱里獲取陰陽

// GetYinYangFromNumber (陰 == 0,  陽 == 1)
func GetYinYangFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "陰"
	case 1:
		return "陽"
	}
	return ""
}

// func GetYinYangFromZhu(pZhu *TZhu) int {

// 	return (pZhu.Gan.Value + 1) % 2
// }

// NewYinYang 創建陰陽
func NewYinYang(nValue int) *TYinYang {
	nValue %= 2
	yinyang := TYinYang(nValue)
	return &yinyang
}

// NewYinYangFromZhu 從柱里創建陰陽
func NewYinYangFromZhu(pZhu *TZhu) *TYinYang {
	return NewYinYangFromGan(pZhu.Gan())
}

// NewYinYangFromGan 從幹里創建陰陽
func NewYinYangFromGan(pGan *TGan) *TYinYang {
	nGan := pGan.Value()
	switch nGan {
	// 甲丙戊庚壬 0, 2, 4, 6, 8 陽 (1)
	case 0, 2, 4, 6, 8:
		return NewYinYang(1)
	// 乙丁己辛癸 1, 3, 5, 7, 9 陰 (0)
	case 1, 3, 5, 7, 9:
		return NewYinYang(0)
	}
	return nil
}

// TYinYang  陰陽
type TYinYang int

// ToString 轉換成可閱讀的字符串
func (m *TYinYang) ToString() string {
	return m.String()
}

// ToInt 轉換成int
func (m *TYinYang) ToInt() int {
	return m.Value()
}

// Value 轉換成int
func (m *TYinYang) Value() int {
	return (int)(*m)
}

// String 轉換成可閱讀的字符串
func (m *TYinYang) String() string {
	return GetYinYangFromNumber(m.Value())
}
