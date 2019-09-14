package command

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCommand(t *testing.T) {
	cmd , other := GetCommand("<@uuu> xx111")
	fmt.Println("cmd:", cmd, "other:", other)
}
