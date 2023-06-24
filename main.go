package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var principal float64 // 贷款本金
	var years int
	var months int
	var rate float64
	var start int
	var prepayment float64
	var duration int

	flag.Float64Var(&principal, "p", 1140000.0, "贷款本金")
	flag.IntVar(&years, "y", 20, "贷款年限")
	flag.Float64Var(&rate, "r", 0.0475, "贷款年利率")
	flag.IntVar(&months, "m", 0, "贷款月数")
	flag.IntVar(&start, "s", 1, "已还月数")
	flag.Float64Var(&prepayment, "e", 300000.0, "提前还款金额")
	flag.IntVar(&duration, "d", 12, "每隔多少个月提前还款")
	flag.Parse()

	interestRate := rate / 12   // 年利率，月利率 = 年利率 / 12
	paymentPeriod := years * 12 // 还款期限，以月为单位
	if months != 0 {
		paymentPeriod = months
	}

	paymentPerMonth := 0.0 // 每月还款额
	totalInterest := 0.0   // 总利息
	totalPayment := 0.0    // 还款总额

	// 每月还款额:贷款本金×[月利率×(1+月利率)^还款月数]÷[(1+月利率)^还款月数-1]
	paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate,
		float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)

	// 计算总利息和还款总额
	totalInterest = paymentPerMonth*float64(paymentPeriod) - principal
	totalPayment = paymentPerMonth * float64(paymentPeriod)

	// 输出计算结果
	fmt.Println("每月还款额(元), 总利息(元), 还款总额(元)")
	fmt.Printf("%.2f, %.2f, %.2f\n", paymentPerMonth, totalInterest, totalPayment)

	//principalPerMon := (principal * interestRate * math.Pow(1+interestRate, float64(11-1))) /
	//	(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
	//fmt.Printf("第11个月的本金: %.2f\n", principalPerMon)

	cnt := 1
	left := principal
	fmt.Println("期, 月供(元), 本金(元), 利息(元), 剩余贷款(元)")
	// 打印每个月的月供和利息
	for i := start; i <= paymentPeriod; i++ {
		extra := 0
		// 每月应还本金=贷款本金×月利率×(1+月利率)^(还款月序号-1)÷〔(1+月利率)^还款月数-1〕
		principalPerMon := (principal * interestRate * math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		//每月应还利息：贷款本金×月利率×〔(1+月利率)^还款月数-(1+月利率)^(还款月序号-1)〕÷〔
		// 			(1+月利率)^还款月数-1〕
		monthInterest := principal * interestRate * (math.Pow(1+interestRate,
			float64(paymentPeriod)) - math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		// 剩余贷款
		left -= principalPerMon
		if i == duration {
			extra = int(prepayment)
			paymentPeriod -= i
			left -= prepayment
			principal = left
			paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate,
				float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
			i = 0
		}

		fmt.Printf("%03d, %.2f, %.2f, %.2f, %.2f\n", cnt, paymentPerMonth+float64(extra), principalPerMon, monthInterest, left)
		cnt++
	}
}
