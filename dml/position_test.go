package dml

import (
	"github.com/samuel-jimenez/xml"
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

func TestMarshalPoistionH(t *testing.T) {
	tests := []struct {
		positionH   *PositionH
		expectedXML string
	}{
		{
			positionH: &PositionH{
				RelativeFrom: dmlst.RelFromHMargin,
				PosOffset:    100,
			},
			expectedXML: `<wp:positionH relativeFrom="margin"><wp:posOffset>100</wp:posOffset></wp:positionH>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.positionH)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalPoistionH(t *testing.T) {
	tests := []struct {
		inputXML    string
		expectedPos PositionH
	}{
		{
			inputXML: `<wp:positionH relativeFrom="margin"><wp:posOffset>100</wp:posOffset></wp:positionH>`,
			expectedPos: PositionH{
				RelativeFrom: dmlst.RelFromHMargin,
				PosOffset:    100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var pos PositionH

			err := xml.Unmarshal([]byte(tt.inputXML), &pos)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if pos.RelativeFrom != tt.expectedPos.RelativeFrom {
				t.Errorf("Expected RelativeFrom %s, but got %s", tt.expectedPos.RelativeFrom, pos.RelativeFrom)
			}
			if pos.PosOffset != tt.expectedPos.PosOffset {
				t.Errorf("Expected PosOffset %d, but got %d", tt.expectedPos.PosOffset, pos.PosOffset)
			}
		})
	}
}

func TestMarshalPoistionV(t *testing.T) {
	tests := []struct {
		positionV   *PositionV
		expectedXML string
	}{
		{
			positionV: &PositionV{
				RelativeFrom: dmlst.RelFromVParagraph,
				PosOffset:    200,
			},
			expectedXML: `<wp:positionV relativeFrom="paragraph"><wp:posOffset>200</wp:posOffset></wp:positionV>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.positionV)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalPoistionV(t *testing.T) {
	tests := []struct {
		inputXML    string
		expectedPos PositionV
	}{
		{
			inputXML: `<wp:positionV relativeFrom="paragraph"><wp:posOffset>200</wp:posOffset></wp:positionV>`,
			expectedPos: PositionV{
				RelativeFrom: dmlst.RelFromVParagraph,
				PosOffset:    200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var pos PositionV

			err := xml.Unmarshal([]byte(tt.inputXML), &pos)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if pos.RelativeFrom != tt.expectedPos.RelativeFrom {
				t.Errorf("Expected RelativeFrom %s, but got %s", tt.expectedPos.RelativeFrom, pos.RelativeFrom)
			}
			if pos.PosOffset != tt.expectedPos.PosOffset {
				t.Errorf("Expected PosOffset %d, but got %d", tt.expectedPos.PosOffset, pos.PosOffset)
			}
		})
	}
}
