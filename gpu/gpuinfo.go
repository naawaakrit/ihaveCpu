// Copyright (c) 2026 Nawakarit
// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package gpu

import "github.com/jaypipes/ghw"

func gpu() *ghw.GPUInfo {
	gpuInfo, err := ghw.GPU()
	if err != nil {
		return nil
	}
	return gpuInfo

}
