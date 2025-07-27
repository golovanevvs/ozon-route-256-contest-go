package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	openAnyFile := false
	for i := 1; i <= 24; i++ {
		file, err := os.Open(fmt.Sprintf("../tests/%d", i))
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %d\n", i)
			continue
		}
		defer file.Close()
		openAnyFile = true

		t.Run(fmt.Sprintf("Test: %d", i), func(t *testing.T) {
			in := bufio.NewReader(file)

			expecteds, err := os.ReadFile(fmt.Sprintf("../tests/%d.a", i))
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			Run(in, out)

			out.Flush()

			actuals, err := io.ReadAll(bufio.NewReader(&buffer))
			require.Nil(t, err)

			var sliceExpecteds []string
			if strings.HasSuffix(string(expecteds), "\r\n") {
				sliceExpecteds = strings.Split(string(expecteds), "\r\n")
			} else {
				sliceExpecteds = strings.Split(string(expecteds), "\n")
			}
			sliceExpecteds1 := sliceExpecteds[:len(sliceExpecteds)-1]
			sliceActuals := strings.Split(string(actuals), "\n")
			sliceActuals1 := sliceActuals[:len(sliceActuals)-1]
			for j, expected := range sliceExpecteds1 {
				actual := sliceActuals1[j]
				assert.Equal(t, expected, actual)
			}
		})
	}
	require.True(t, openAnyFile)
}
