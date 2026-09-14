package hardware

import "testing"

func TestParseOutput(t *testing.T) {
	output := " system " + separator + " bios " + separator + " cpu " + separator + " cache " + separator + " ram " + separator + " board " + separator + " pcie " + separator + " power " + separator + " battery "

	data, err := ParseOutput(output)
	if err != nil {
		t.Fatalf("ParseOutput() error = %v", err)
	}

	expected := Data{
		System:    "system",
		BIOS:      "bios",
		CPU:       "cpu",
		Cache:     "cache",
		RAM:       "ram",
		Mainboard: "board",
		PCIe:      "pcie",
		Power:     "power",
		Battery:   "battery",
	}
	if data != expected {
		t.Fatalf("ParseOutput() = %+v, want %+v", data, expected)
	}
}

func TestParseOutputRejectsIncompleteData(t *testing.T) {
	_, err := ParseOutput("system" + separator + "bios")
	if err == nil {
		t.Fatal("ParseOutput() error = nil, want incomplete output error")
	}
}
