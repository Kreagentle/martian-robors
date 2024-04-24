package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "First case 1 1 E",
			input:          "5 3\n1 1 E\nRFRFRFRF\nno\n",
			expectedOutput: "1 1 E",
			expectedError:  nil,
		},
		{
			name:           "Second case 3 2 N",
			input:          "5 3\n3 2 N\nFRRFLLFFRRFLL\nno\n",
			expectedOutput: "3 3 N LOST",
			expectedError:  nil,
		},
		{
			name:           "Third case 3 2 N right",
			input:          "5 3\n3 2 N\nR\nno\n",
			expectedOutput: "3 2 E",
			expectedError:  nil,
		},
		{
			name:           "Fourth case 3 2 N left",
			input:          "5 3\n3 2 N\nL\nno\n",
			expectedOutput: "3 2 W",
			expectedError:  nil,
		},
		{
			name:           "Fifth case 3 2 S right",
			input:          "5 3\n3 2 S\nR\nno\n",
			expectedOutput: "3 2 W",
			expectedError:  nil,
		},
		{
			name:           "Sixth case 3 2 S left",
			input:          "5 3\n3 2 S\nL\nno\n",
			expectedOutput: "3 2 E",
			expectedError:  nil,
		},
		{
			name:           "Seventh case 3 2 N forward",
			input:          "10 10\n3 2 N\nF\nno\n",
			expectedOutput: "3 3 N",
			expectedError:  nil,
		},
		{
			name:           "Eighth case 3 2 N forward forward",
			input:          "10 10\n3 2 N\nFF\nno\n",
			expectedOutput: "3 4 N",
			expectedError:  nil,
		},
		{
			name:           "Ninth case 0 0 N",
			input:          "0 0\n0 0 N\nF\nno\n",
			expectedOutput: "0 0 N LOST",
			expectedError:  nil,
		},
		{
			name:           "Tenth case 3 2 N nothing",
			input:          "10 10\n3 2 N\n\nno\n",
			expectedOutput: "3 2 N",
			expectedError:  nil,
		},
		{
			name:           "Eleventh case 1 1 N",
			input:          "1 1\n0 0 N\nFF\nno\n",
			expectedOutput: "0 1 N LOST",
			expectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("", "example")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err = tmpfile.Write([]byte(test.input)); err != nil {
				t.Fatal(err)
			}
			if err = tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin, err = os.Open(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}

			tmpOutputFile, err := ioutil.TempFile("", "output")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpOutputFile.Name())
			oldStdout := os.Stdout
			defer func() { os.Stdout = oldStdout }()
			os.Stdout = tmpOutputFile

			err = Execute()
			if err != test.expectedError {
				t.Errorf("Unexpected error: got %v, want %v", err, test.expectedError)
			}

			file, err := os.Open(tmpOutputFile.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			var outputLines []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if index := strings.Index(line, "The result coordinates are:"); index != -1 {
					outputLines = append(outputLines, line[index+len("The result coordinates are: "):])
				}
			}

			if err = scanner.Err(); err != nil {
				t.Fatal(err)
			}

			if len(outputLines) == 0 {
				t.Fatal("Expected output not found")
			}

			gotOutput := strings.Join(outputLines, "\n")
			if gotOutput != test.expectedOutput {
				t.Errorf("Unexpected output: got %s, want %s", gotOutput, test.expectedOutput)
			}
		})
	}
}
