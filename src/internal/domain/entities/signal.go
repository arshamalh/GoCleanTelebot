package entities

import (
	"strconv"
	"strings"
)

// Signal entity
type Signal struct {
	ID          string
	Pair        string
	Date        string // unix timestamp stored as string
	ImageURL    string // It's just a url have to be downloaded
	TimeFrame   string
	EntryPrice  string // my Own recognition
	TargetPrice string // maybe it can be multi target
	StopLoss    string
	Risk2Reward string
	TradeVolume string // there is not a trade volume on excel file
}

// IsValid check signal validity, for example, len of characters and return suitable result
func (s Signal) IsValid() (bool, string) {
	if !(len(s.Pair) > 3 || len(s.Pair) < 20) { // Pair len has to be between 3 & 20
		return false, "Pair is not correct"
	} else if val, err := strconv.ParseFloat(s.StopLoss, 64); err != nil || !(0 < len(s.StopLoss) && len(s.StopLoss) < 20) || val == 0 {
		return false, "StopLoss is not correct"
	} else if val, err := strconv.ParseFloat(s.EntryPrice, 64); err != nil || !(0 < len(s.EntryPrice) && len(s.EntryPrice) < 20) || val == 0 {
		return false, "EntryPrice is not correct"
	} else if val, err := strconv.ParseFloat(s.TargetPrice, 64); err != nil || !(0 < len(s.TargetPrice) && len(s.TargetPrice) < 20) || val == 0 {
		return false, "TargetPrice is not correct"
	}
	return true, ""
}

// ToString make an abstract string of signal
func (s Signal) ToString() string {
	return "Pair: " + s.Pair + "\nDate: " + s.Date + "\nEntryPrice: " + s.EntryPrice + "\nTargetPrice: " + s.TargetPrice + "\nStopLoss: " + s.StopLoss + "\n"
}

// SignalDirection You have to buy or sell depends on this signal?
func (s Signal) SignalDirection() string {
	entry, _ := strconv.ParseFloat(s.EntryPrice, 64)
	target, _ := strconv.ParseFloat(s.TargetPrice, 64)
	if entry < target {
		return "Buy \U0001F7E2"
	}

	return "Sell 🔴"

}

// DateRef make date like this 1401/5/3 => 01/05/03
func (s Signal) DateRef(sep string) string {
	// Do will refactor date to appropriate format
	b := strings.Split(s.Date, sep)
	for i, v := range b {
		if len(v) == 1 {
			b[i] = "0" + b[i]
		} else if len(v) == 4 {
			b[i] = b[i][2:]
		}
	}
	return strings.Join(b, sep)
}
