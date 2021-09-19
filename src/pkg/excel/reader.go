package excel

import (
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"encoding/csv"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strings"
)

// ReadSignals from provided file pass, it just support xlsx and csv files.
func ReadSignals(filePath, sheet string) ([]entities.Signal, error) {
	if cond := strings.HasSuffix(strings.ToLower(filePath), ".xlsx"); cond {
		return xlsxHandler(filePath, sheet)
	}
	if cond := strings.HasSuffix(strings.ToLower(filePath), ".csv"); cond {
		return csvHandler(filePath)
	}
	return nil, errors.New("unsupported file type")
}

func xlsxHandler(filePath, sheet string) ([]entities.Signal, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	if sheet == "" {
		sheet = f.GetSheetName(f.GetActiveSheetIndex())
	}
	rows := f.GetRows(sheet)
	if len(rows) == 0 {
		return []entities.Signal{}, errors.New("the file is empty or is ruined")
	}
	return makeOutput("xlsx", rows), nil
}

func csvHandler(filePath string) ([]entities.Signal, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return nil, err
	}
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return []entities.Signal{}, errors.New("the file is empty or is ruined")
	}
	return makeOutput("csv", rows), nil
}

func makeOutput(kind string, rows [][]string) []entities.Signal {
	var outPut []entities.Signal
	for i, row := range rows {
		if row[2] != "" && (i != 0 || kind == "csv") { /// rows that are fake OR first row
			outPut = append(outPut, entities.Signal{
				Pair:        row[2],
				Date:        row[3],
				ImageURL:    row[4],
				TimeFrame:   row[5],
				EntryPrice:  row[7],
				TargetPrice: row[8],
				StopLoss:    row[9],
				Risk2Reward: row[12],
				TradeVolume: row[21],
			})
		}
	}
	return outPut
}
