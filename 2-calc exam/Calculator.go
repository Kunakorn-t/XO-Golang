package main

import (
	"fmt"
	"math"
	"strconv"
)

var results, resultInput1, resultInput2, resultInput3 int
var operator, minus string
var exit int = 1

var arrResultInput1, arrResultInput2 []int

func main() {
	var input1, input2, input3 []string
	var inp1, inp2, inp3 string
	var number1Input1, number2Input1, number1Input2, number2Input2, number1Input3 string
	var intNumber1Input1, intNumber2Input1, intNumber1Input2, intNumber2Input2, intNumber1Input3 int

	// while
	// if inp1 == "/exit" {
	// 	fmt.Println("Bye!")
	// 	fmt.Println()
	// 	fmt.Println("Process finished with exit code 0")

	// 	fmt.Scan(&exit)

	// }
	// if exit == 0 {
	// 	break
	// }

	for {
		fmt.Scanf("%s %s %s", &inp1, &inp2, &inp3)

		// input1
		num := len(inp1)
		for i := 0; i < num; i++ {
			input1 = append(input1, string([]rune(inp1)[i]))
		}

		for i, s := range input1 {
			switch {
			case s == "*":

				n := 0
				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					resultInput1 = intNumber1Input1

				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 >= num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							break
						}

					}
					intNumber2Input1, _ = strconv.Atoi(number2Input1)
					resultInput1 = intNumber1Input1 * intNumber2Input1
					arrResultInput1 = append(arrResultInput1, resultInput1)
				}
				break
			case s == "/":
				// fmt.Println(input1)

				n := 0

				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					// fmt.Println(intNumber1Input1)
					resultInput1 = intNumber1Input1
				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 == num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							break
						}
					}
					intNumber2Input1, _ = strconv.Atoi(number2Input1)
					// fmt.Println(intNumber2Input1)

					resultInput1 = intNumber1Input1 / intNumber2Input1
					arrResultInput1 = append(arrResultInput1, resultInput1)
				}
				break
			case s == "%":
				// fmt.Println(input1)

				n := 0

				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					// fmt.Println(intNumber1Input1)
					resultInput1 = intNumber1Input1
				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 == num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							break
						}
					}
					intNumber2Input1, _ = strconv.Atoi(number2Input1)
					// fmt.Println(intNumber2Input1)

					resultInput1 = intNumber1Input1 % intNumber2Input1
					arrResultInput1 = append(arrResultInput1, resultInput1)
				}
				break
			case s == "+":
				// fmt.Println(input1)

				n := 0

				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					// fmt.Println(intNumber1Input1)
					resultInput1 = intNumber1Input1
				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 == num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							break
						}
					}
					intNumber2Input1, _ = strconv.Atoi(number2Input1)
					// fmt.Println(intNumber2Input1)

					resultInput1 = intNumber1Input1 + intNumber2Input1
					arrResultInput1 = append(arrResultInput1, resultInput1)
				}
				break
			case s == "-":
				n := 0

				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					// fmt.Println(intNumber1Input1)
					resultInput1 = intNumber1Input1
				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 == num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							break
						}
					}
					intNumber2Input1, _ = strconv.Atoi(number2Input1)
					// fmt.Println(intNumber2Input1)

					resultInput1 = intNumber1Input1 - intNumber2Input1
					arrResultInput1 = append(arrResultInput1, resultInput1)
					minus = "minus"
				}
				break
			case s == "^":
				n := 0
				if input1[i-1] != "" {
					for {
						number1Input1 += input1[n]
						n++
						if n == i {
							break
						}
					}
					intNumber1Input1, _ = strconv.Atoi(number1Input1)
					resultInput1 = intNumber1Input1
				}
				if i+1 <= num {
					n2 := i + 1

					for {
						if i+1 == num {
							break
						}
						number2Input1 += input1[n2]
						n2++
						if n2 == len(input1) {
							intNumber2Input1, _ = strconv.Atoi(number2Input1)
							float1 := float64(intNumber1Input1)
							float2 := float64(intNumber2Input1)
							logarithm := math.Pow(float1, float2)
							resultInput1 = int(logarithm)

							arrResultInput1 = append(arrResultInput1, resultInput1)
							break
						}
					}
				}
				break
			default:
				intNumber1Input1, _ = strconv.Atoi(inp1)
				resultInput1 = intNumber1Input1
				break
			}
		}

		iarr := 0
		for {

			if arrResultInput1[iarr] != 0 || minus == "minus" {
				resultInput1 = arrResultInput1[iarr]
				break
			}
			iarr++
		}

		// input2
		// ต้องเช็คว่า เป็น operator or number
		// และเช็คต่อด้วยว่า มี operator or number Input1ตามหลังอีกมั้ย
		num2 := len(inp2)
		for i := 0; i < num2; i++ {
			input2 = append(input2, string([]rune(inp2)[i]))
		}

		if inp2 != "" {
			switch {
			case input2[0] == "*":
				operator = "*"
				break
			case input2[0] == "/":
				operator = "/"
				break
			case input2[0] == "%":
				operator = "%"
				break
			case input2[0] == "+":
				operator = "+"
				break
			case input2[0] == "-":
				operator = "-"
				break
			case input2[0] == "^":
				operator = "^"
				break
			case input2[0] == "(":
				if input2[num2-1] == ")" {
					operator = "*"
					for i, s := range input2 {

						switch {
						case s == "*":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								resultInput2 = intNumber1Input2 * intNumber2Input2
							}
							break

						case s == "/":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								resultInput2 = intNumber1Input2 / intNumber2Input2
							}
							break
						case s == "%":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								resultInput2 = intNumber1Input2 % intNumber2Input2
							}
							break
						case s == "+":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								resultInput2 = intNumber1Input2 + intNumber2Input2
							}
							break
						case s == "-":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								resultInput2 = intNumber1Input2 - intNumber2Input2
							}
							break
						case s == "^":
							if input2[1] != "" {
								n2 := 1
								for {
									number1Input2 += input2[n2]
									n2++
									if n2 == i {
										break
									}
								}
								intNumber1Input2, _ = strconv.Atoi(number1Input2)
								resultInput2 = intNumber1Input2
							}
							if input2[i+1] != "" {
								n2 := i + 1

								for {
									if i+1 == num2 {
										break
									}
									number2Input2 += input2[n2]
									n2++
									if n2 >= num2-1 {
										break
									}
								}
								intNumber2Input2, _ = strconv.Atoi(number2Input2)
								float1 := float64(intNumber1Input2)
								float2 := float64(intNumber2Input2)
								logarithm := math.Pow(float1, float2)
								resultInput2 = int(logarithm)
							}
							break

						}
					}
				} else {
					fmt.Println("Invalid expression1")
				}
				break

			default:
				fmt.Println("Invalid expression2")
				break
			}
		}

		num3 := len(inp3)
		for i := 0; i < num3; i++ {
			input3 = append(input3, string([]rune(inp3)[i]))
		}

		n3 := 0
		if inp3 != "" && resultInput3 == 0 {
			for {
				number1Input3 += input3[n3]
				n3++
				if n3 >= num3 {
					break
				}
			}
			intNumber1Input3, _ = strconv.Atoi(number1Input3)
			resultInput3 = intNumber1Input3
		}

		fmt.Println(resultInput1, resultInput2, resultInput3, operator)

		switch {
		case operator == "*":
			if resultInput2 == 0 {
				results = resultInput1 * resultInput3
			}
			if resultInput3 == 0 {
				results = resultInput1 * resultInput2
			}
			break
		case operator == "/":
			if resultInput2 == 0 {
				results = resultInput1 / resultInput3
			}
			if resultInput3 == 0 {
				results = resultInput1 / resultInput2
			}
			break
		case operator == "%":
			if resultInput2 == 0 {
				results = resultInput1 % resultInput3
			}
			if resultInput3 == 0 {
				results = resultInput1 % resultInput2
			}
			break
		case operator == "+":
			if resultInput2 == 0 {
				results = resultInput1 + resultInput3
			}
			if resultInput3 == 0 {
				results = resultInput1 + resultInput2
			}
			break
		case operator == "-":
			if resultInput2 == 0 {
				results = resultInput1 - resultInput3
			}
			if resultInput3 == 0 {
				results = resultInput1 - resultInput2
			}
			break
		default:
			results = resultInput1
			break
		}
		fmt.Println(results)

		results, resultInput1, resultInput2, resultInput3 = 0, 0, 0, 0

		input1, input2, input3, arrResultInput1 = nil, nil, nil, nil

		inp1, inp2, inp3, operator, number1Input1, number2Input1, number1Input2, number2Input2, number1Input3, minus = "", "", "", "", "", "", "", "", "", ""
		intNumber1Input1, intNumber2Input1, intNumber1Input2, intNumber2Input2, intNumber1Input3 = 0, 0, 0, 0, 0
		if inp1 == "/exit" {
			break
		}

	}

}

// สุดท้ายคือเอา 3 input มา implement กัน
// ถ้ามากกว่า 3 input ให้แสดงข้อความบอกใช้ได้แค่ 3
