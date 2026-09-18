package gofpdf

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpdf_SplitText(t *testing.T) {
	type fields struct {
		cMargin  float64
		fontSize float64
	}
	type args struct {
		txt string
		w   float64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantLines []string
	}{
		{
			name: "ascii normal",
			fields: fields{
				cMargin:  3,
				fontSize: 12,
			},
			args: args{
				txt: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				w:   60,
			},
			wantLines: []string{
				"Lorem ipsum dolor sit",
				"amet, consectetur",
				"adipiscing elit, sed do",
				"eiusmod tempor",
				"incididunt ut labore et",
				"dolore magna aliqua.",
				"Ut enim ad minim",
				"veniam, quis nostrud",
				"exercitation ullamco",
				"laboris nisi ut aliquip",
				"ex ea commodo",
				"consequat. Duis aute",
				"irure dolor in",
				"reprehenderit in",
				"voluptate velit esse",
				"cillum dolore eu fugiat",
				"nulla pariatur.",
				"Excepteur sint",
				"occaecat cupidatat non",
				"proident, sunt in culpa",
				"qui officia deserunt",
				"mollit anim id est",
				"laborum.",
			},
		},
		{
			name: "ascii emoji",
			fields: fields{
				cMargin:  3,
				fontSize: 12,
			},
			args: args{
				txt: "Lorem ipsum dolor si️t ⚠ amet,🇨🇭 consectetur adipiscing 🙈 elit, sed do eiusmod tempor 🎆 incididunt ut labore et dolore 🤷‍♂️ magna aliqua. Ut enim ad minim 🥰 veniam, quis nostrud exercitation ullamco laboris nisi 🎉 ut aliquip ex ea commodo consequat.",
				w:   50,
			},
			wantLines: []string{
				"Lorem ipsum",
				"dolor si️t ⚠ amet,🇨🇭",
				"consectetur",
				"adipiscing 🙈 elit,",
				"sed do eiusmod",
				"tempor 🎆 incididunt",
				"ut labore et dolore 🤷\u200d♂️",
				"magna aliqua. Ut",
				"enim ad minim 🥰",
				"veniam, quis",
				"nostrud",
				"exercitation",
				"ullamco laboris",
				"nisi 🎉 ut aliquip ex",
				"ea commodo",
				"consequat.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdf := New("P", "mm", "A4", "") // A4 210.0 x 297.0

			var fr fontResourceType
			pdf.SetFontLoader(fr)
			pdf.AddFont("Calligra", "", "calligra.json")
			pdf.SetFont("Calligra", "", 16)

			if pdf.Error() != nil {
				t.Fatal(pdf.Error())
			}

			gotLines := pdf.SplitText(tt.args.txt, tt.args.w)
			assert.Exactly(t, tt.wantLines, gotLines)
		})
	}
}

type fontResourceType struct{}

func (f fontResourceType) Open(name string) (rdr io.Reader, err error) {
	var buf []byte
	buf, err = os.ReadFile("./font/" + name)
	if err == nil {
		rdr = bytes.NewReader(buf)
	}
	return
}
