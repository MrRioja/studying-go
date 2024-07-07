package main

import "fmt"

func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func dayOfWeek2(day int) string {
	switch {
	case day == 1:
		return "Sunday"
	case day == 2:
		return "Monday"
	case day == 3:
		return "Tuesday"
	case day == 4:
		return "Wednesday"
	case day == 5:
		return "Thursday"
	case day == 6:
		return "Friday"
	case day == 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func dayOfWeekWithFallthrough(day int) string {
	var dayOfWeek string

	switch day {
	case 1:
		dayOfWeek = "Sunday"
		fallthrough
	case 2:
		dayOfWeek = "Monday"
	case 3:
		dayOfWeek = "Tuesday"
	case 4:
		dayOfWeek = "Wednesday"
	case 5:
		dayOfWeek = "Thursday"
	case 6:
		dayOfWeek = "Friday"
	case 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "Invalid day"
	}

	return dayOfWeek
}

func main() {
	fmt.Println(dayOfWeek(1))
	fmt.Println(dayOfWeek(2))
	fmt.Println(dayOfWeek(3))
	fmt.Println(dayOfWeek(4))
	fmt.Println(dayOfWeek2(5))
	fmt.Println(dayOfWeek2(6))
	fmt.Println(dayOfWeek2(7))
	fmt.Println(dayOfWeek2(8))
	fmt.Println()
	fmt.Println(dayOfWeekWithFallthrough(1))
}
