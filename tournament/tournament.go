package tournament

import (
	"fmt"
	"io"
	"bytes"
	"strings"
	"sort"

)

type Record struct {
	MP, W, D, L, P int
	name string
}

type Records []*Record

func NewRecord(name string) (record *Record) {

	var new_record Record
	new_record.name = name
	return &new_record
}


func (record *Record) display() {

	fmt.Printf("\n name : %s points : %d", record.name, record.P)
}

func (record *Record) update_result(result string) {

	record.MP += 1
	switch result {

		case "win":
			record.W += 1
			record.P += 3
		case "loss":
			record.L += 1
		case "draw":
			record.D += 1
			record.P += 1
		default:
	}

}

func toggle_result(result string) string {

	if result == "win" {
		return "loss"
	} else if result == "loss" {
		return "win"
	} else {
		return "draw"
	}

}

func (records *Records) fetch_record(name string) (record *Record) {

	for _, record := range *records {

		if record.name == name {

			return record
		}
	}

	new_record := NewRecord(name)
	*records = append(*records, new_record)

	return new_record
}

func (records *Records) update(name1, name2, result string) {

	//fmt.Println("Inside records update")

	team1, team2 := records.fetch_record(name1), records.fetch_record(name2)

	//records = append(records, team1)

	//fmt.Println("team1 : ", team1)
	//fmt.Println("records : ", records)


	team1.update_result(result)
	team2.update_result(toggle_result(result))

}

func split_game(game string) (string, string, string) {

	match_day := strings.Split(game, ";")
	return match_day[0], match_day[1], match_day[2]
}

func (record *Record) String() string {

	// return fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d",
	//                    record.name, record.MP, record.W, record.D, record.L, record.P)
	format_string := "%-31s| %2v | %2v | %2v | %2v | %2v"
	return fmt.Sprintf(format_string,
					   record.name, record.MP, record.W, record.D, record.L, record.P)
}

func (records Records) Len() int {

	return len(records)
}

func (records Records) Less(i, j int) bool {

	return records[i].P > records[j].P
}

func (records Records) Swap(i, j int) {

	records[i], records[j] = records[j], records[i]
}

func parseInput(reader io.Reader) string {

	//fmt.Println("Inside parseInput reader : ", reader)

	var buffer = new(bytes.Buffer)
	var records Records
	
	buffer.ReadFrom(reader)
	input := buffer.String()
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "\n")
	games := strings.Split(input, "\n")
	//fmt.Println("Inside parseInput games : ", games)
	//fmt.Println("Inside parseInput games len: ", len(games))
	for _, game := range games {
		
		//fmt.Println("game :", game)
		team1, team2, result := split_game(game)
		records.update(team1, team2, result)
	}
	//fmt.Println("records :", records)

	sort.Sort(records)
	// sort.Slice(records, func(i, j int) bool {
	// 	return records[i].P <= records[i].P
	// })

	

	header := "Team                           | MP |  W |  D |  L |  P"

	// for _, record := range records {
	// 	record.display()
	// 	fmt.Printf("\n record : %s", record)
	// }

	op := header
	for _, record := range records {
		op += "\n" + fmt.Sprintf("%s", record)
	}
	//fmt.Println("op :", op)
	//fmt.Printf("\n nbytes : %d", nbytes)



	return op

}


func Tally(reader io.Reader, buffer io.Writer) error {

	fmt.Println("Inside Tally")

	op := parseInput(reader)

	io.WriteString(buffer, op)

	return nil
	
}