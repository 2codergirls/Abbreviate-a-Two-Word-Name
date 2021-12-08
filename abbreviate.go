package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	firstName := names[0]
	lastName := names[1]

	firstNameInitial := strings.ToUpper(string(firstName[0:1]))
	lastNameInitial := strings.ToUpper(string(lastName[0:1]))

	return fmt.Sprintf("%s.%s", firstNameInitial, lastNameInitial)
}
