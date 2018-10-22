package main

import (
	"fmt"
	"time"
)

const (
	Iac  = 255
	Will = 251
	Sb   = 250
	Se   = 240

	ComPortOpt = 44
	SetParity  = 3
	SetControl = 5

	PAR_NONE = 1
	PAR_EVEN = 3

	FLOW_OFF = 1
	DTR_ON   = 8
	DTR_OFF  = 9
	RTS_ON   = 11
	RTS_OFF  = 12
)

func telnetInit() {
	serialSend <- []byte{Iac, Will, ComPortOpt}
	telnetEscape(SetParity, PAR_NONE)
	telnetEscape(SetControl, FLOW_OFF)
	telnetEscape(SetControl, RTS_ON)
	telnetEscape(SetControl, DTR_OFF)
}

func telnetEscape(typ, val uint8) {
	if *verbose {
		fmt.Printf("{esc:%02X:%02X}", typ, val)
	}
	serialSend <- []byte{Iac, Sb, ComPortOpt, typ, val, Iac, Se}
}

func telnetReset(enterBoot bool) {
	telnetEscape(SetControl, DTR_ON)
	if enterBoot {
		telnetEscape(SetControl, RTS_OFF)
		telnetEscape(SetParity, PAR_EVEN)
	} else {
		telnetEscape(SetControl, RTS_ON)
		telnetEscape(SetParity, PAR_NONE)
	}
	time.Sleep(100 * time.Millisecond)
	telnetEscape(SetControl, DTR_OFF)
}

// telnetClean removes incoming telnet escape commands from the input buffer
func telnetClean(buf []byte, n int) int {
	j := 0
	for i := 0; i < n; i++ {
		b := buf[i]
		buf[j] = b
		switch tnState {
		case 0: // normal, copying
			if b == Iac {
				tnState = 1
			} else {
				j++
			}
		case 1: // seen Iac
			if b == Sb {
				tnState = 2
			} else if b >= Will {
				tnState = 3
			} else {
				j++
				tnState = 0
			}
		case 2: // inside command
			if b == Se {
				tnState = 0
			}
		case 3: // Will/Wont/Do/Dont seen
			tnState = 0
		}
	}
	return j
}
