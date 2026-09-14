// Copyright (c) 2026 Nawakarit
// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package hardware

import (
	"fmt"
	"os/exec"
	"strings"
)

const separator = "(-@_@-)"

// Data contains the hardware sections collected from dmidecode.
type Data struct {
	System    string
	BIOS      string
	CPU       string
	Cache     string
	RAM       string
	Mainboard string
	PCIe      string
	Power     string
	Battery   string
}

const dmidecodeScript = `dmidecode -t 1 && dmidecode -t 3 && dmidecode -t 12 && dmidecode -t 15 && dmidecode -t 23 && dmidecode -t 24 && dmidecode -t 32
echo '(-@_@-)' && dmidecode -t 0 && dmidecode -t 13 && dmidecode -t 40 && dmidecode -t 45
echo '(-@_@-)' && dmidecode -t 4
echo '(-@_@-)' && dmidecode -t 7
echo '(-@_@-)' && dmidecode -t 5 && dmidecode -t 6 && dmidecode -t 16 && dmidecode -t 17 && dmidecode -t 18 && dmidecode -t 19 && dmidecode -t 20 && dmidecode -t 33 && dmidecode -t 37
echo '(-@_@-)' && dmidecode -t 2 && dmidecode -t 10 && dmidecode -t 41
echo '(-@_@-)' && dmidecode -t 8 && dmidecode -t 9
echo '(-@_@-)' && dmidecode -t 25 && dmidecode -t 26 && dmidecode -t 27 && dmidecode -t 28 && dmidecode -t 29 && dmidecode -t 39
echo '(-@_@-)' && dmidecode -t 22`

// Load reads hardware data through pkexec and dmidecode.
func Load() (Data, error) {
	cmd := exec.Command("pkexec", "sh", "-c", dmidecodeScript)
	out, err := cmd.CombinedOutput()
	if err != nil {
		details := strings.TrimSpace(string(out))
		if details != "" {
			return Data{}, fmt.Errorf("ไม่สามารถอ่านข้อมูลฮาร์ดแวร์: %w\nรายละเอียด: %s", err, details)
		}
		return Data{}, fmt.Errorf("ไม่สามารถอ่านข้อมูลฮาร์ดแวร์: %w", err)
	}

	data, err := ParseOutput(string(out))
	if err != nil {
		return Data{}, err
	}
	return data, nil
}

// ParseOutput converts the dmidecode output into named hardware sections.
func ParseOutput(output string) (Data, error) {
	parts := strings.Split(output, separator)
	if len(parts) < 9 {
		return Data{}, fmt.Errorf("output ไม่ครบ: ได้ %d ส่วน", len(parts))
	}

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return Data{
		System:    parts[0],
		BIOS:      parts[1],
		CPU:       parts[2],
		Cache:     parts[3],
		RAM:       parts[4],
		Mainboard: parts[5],
		PCIe:      parts[6],
		Power:     parts[7],
		Battery:   parts[8],
	}, nil
}
