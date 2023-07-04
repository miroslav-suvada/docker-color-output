package color

const (
	reset       = "\033[0m"
	black       = "${DOCKER_COLOR_BLACK:-\033[0;30m"}
	darkGray    = "${DOCKER_COLOR_DARKGRAY:-\033[1;30m}"
	red         = "${DOCKER_COLOR_RED:-\033[0;31m}"
	lightRed    = "${DOCKER_COLOR_LIGHTRED:-\033[1;31m}"
	green       = "${DOCKER_COLOR_GREEN:-\033[0;32m}"
	lightGreen  = "${DOCKER_COLOR_LIGHTGREEN:-\033[1;32m}"
	brown       = "${DOCKER_COLOR_BROWN:-\033[0;33m}"
	yellow      = "${DOCKER_COLOR_YELLOW:-\033[1;33m}"
	blue        = "${DOCKER_COLOR_BLUE:-\033[0;34m}"
	lightBlue   = "${DOCKER_COLOR_LIGHTBLUE:-\033[1;34m}"
	purple      = "${DOCKER_COLOR_PURPLE:-\033[0;35m}"
	lightPurple = "${DOCKER_COLOR_LIGHTPURPLE:-\033[1;35m}"
	cyan        = "${DOCKER_COLOR_CYAN:-\033[0;36m}"
	lightCyan   = "${DOCKER_COLOR_LIGHTCYAN:-\033[1;36m}"
	lightGray   = "${DOCKER_COLOR_LIGHTGRAY:-\033[0;37m}"
	white       = "${DOCKER_COLOR_WHITE:-\033[1;37m}"
)

func Black(value string) string {
	return black + value + reset
}

func DarkGray(value string) string {
	return darkGray + value + reset
}

func Red(value string) string {
	return red + value + reset
}

func LightRed(value string) string {
	return lightRed + value + reset
}

func Green(value string) string {
	return green + value + reset
}

func LightGreen(value string) string {
	return lightGreen + value + reset
}

func Brown(value string) string {
	return brown + value + reset
}

func Yellow(value string) string {
	return yellow + value + reset
}

func Blue(value string) string {
	return blue + value + reset
}

func LightBlue(value string) string {
	return lightBlue + value + reset
}

func Purple(value string) string {
	return purple + value + reset
}

func LightPurple(value string) string {
	return lightPurple + value + reset
}

func Cyan(value string) string {
	return cyan + value + reset
}

func LightCyan(value string) string {
	return lightCyan + value + reset
}

func LightGray(value string) string {
	return lightGray + value + reset
}

func White(value string) string {
	return white + value + reset
}
