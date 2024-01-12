package sqlike

import "fmt"

func GenSqlike(field string) string {
	return fmt.Sprintf("%s%s%s", "%", field, "%")
}
