package main


func createPlaceholder() [10][5]string {
	type placeholder [5]string

	zero := placeholder {
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder {
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder {
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder {
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder {
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder {
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder {
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder {
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder {
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder {
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [10][5]string{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	return digits
}

func createAlarmSign() [10][5]string {
	type placeholder [5]string

	emp := placeholder {
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	a := placeholder {
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	}

	l := placeholder {
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	}

	r := placeholder {
		"██ ",
		"█ █",
		"██ ",
		"█ █",
		"█ █ ",
	}

	m := placeholder {
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	}

	exl := placeholder {
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	}

	return [10][5]string {
		emp, emp, a, l, a, r, m, exl, emp, emp,
	}

	
}

var dot = [5]string {
	"   ",
	"   ",
	"   ",
	"   ",
	" █ ",
}