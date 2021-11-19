package env_bug_test

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvBug(t *testing.T) {
	fmt.Println(os.Getenv("DB_URL"))
}
