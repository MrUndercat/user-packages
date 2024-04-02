package hidespam

import "testing"

func TestAdd(t *testing.T) {
	type Args struct {
		text, expected string
	}

	var tests = []Args{
		{"Here's my spammy page: http://hehefouls.netHAHAHA see you.", "Here's my spammy page: http://******************* see you."},
		{"Fuck your job,  fuck you boss", "Fuck your job,  fuck you boss"},
		{"http://www.digitalocean.com nill kiggers", "http://******************** nill kiggers"},
		{"http://www.вконтактеневконтакте.ком стоп почему на русском", "http://**************************** стоп почему на русском"},
	}
	for _, test := range tests {
		output := HideSpam(test.text)
		if output != test.expected {
			t.Errorf("Output %q not exual to expected %q", output, test.expected)
		}
	}
}
