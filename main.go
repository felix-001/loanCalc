package main

import (
	"fmt"
	"math"
)

func main() {
	// 输入贷款本金、年利率、还款期限
	principal := 1140000.0      // 贷款本金
	interestRate := 0.0475 / 12 // 年利率，月利率 = 年利率 / 12
	paymentPeriod := 20 * 12    // 还款期限，以月为单位
	paymentPerMonth := 0.0      // 每月还款额
	totalInterest := 0.0        // 总利息
	totalPayment := 0.0         // 还款总额

	// 每月还款额:贷款本金×[月利率×(1+月利率)^还款月数]÷[(1+月利率)^还款月数-1]
	paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate,
		float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)

	// 计算总利息和还款总额
	totalInterest = paymentPerMonth*float64(paymentPeriod) - principal
	totalPayment = paymentPerMonth * float64(paymentPeriod)

	// 输出计算结果
	fmt.Println("每月还款额,总利息,还款总额")
	fmt.Printf("%.2f, %.2f, %.2f\n", paymentPerMonth, totalInterest, totalPayment)

	fmt.Println("期,月供,本金,利息")
	// 打印每个月的月供和利息
	for i := 1; i <= paymentPeriod; i++ {
		// 每月应还本金=贷款本金×月利率×(1+月利率)^(还款月序号-1)÷〔(1+月利率)^还款月数-1〕
		principalPerMon := (principal * interestRate * math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		//每月应还利息：贷款本金×月利率×〔(1+月利率)^还款月数-(1+月利率)^(还款月序号-1)〕÷〔
		// 			(1+月利率)^还款月数-1〕
		monthInterest := principal * interestRate * (math.Pow(1+interestRate,
			float64(paymentPeriod)) - math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		fmt.Printf("%03d,%.2f元,%.2f元,%.2f元\n", i, paymentPerMonth, principalPerMon, monthInterest)
	}
}
