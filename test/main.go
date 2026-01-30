package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jjonline/go-bazi-calculator/bazi"
)

func main() {
	var nYear int
	var nMonth int
	var nDay int
	var nHour int
	var nMinute int
	var nSecond int
	var nSex int

	flag.IntVar(&nYear, "y", 1989, "-y=1995 ")
	flag.IntVar(&nMonth, "m", 7, "-m=6 ")
	flag.IntVar(&nDay, "d", 1, "-d=16 ")
	flag.IntVar(&nHour, "h", 1, "-h=19 ")
	flag.IntVar(&nMinute, "n", 0, "-n=7 ")
	flag.IntVar(&nSecond, "s", 0, "-s=0 ")
	flag.IntVar(&nSex, "x", 0, "-x=0  1是男0是女 ")

	flag.Parse() //解析輸入的參數

	inputTime := time.Date(
		nYear,
		time.Month(nMonth),
		nDay,
		nHour,
		nMinute,
		nSecond,
		0,
		time.FixedZone("Local", int(8*3600)),
	)

	sunTime := bazi.ApparentSolarTime(inputTime, 114.12)

	bz := bazi.GetBazi(
		sunTime.Year(),
		int(sunTime.Month()),
		sunTime.Day(),
		sunTime.Hour(),
		sunTime.Minute(),
		sunTime.Second(),
		nSex,
	)
	fmt.Println(sunTime)
	fmt.Println(bz.DaYun())
}
