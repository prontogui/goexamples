package main

/*
List<dl.Primitive> makeRow(String letter, int fromNo, int toNo,
    {String letterEmbodiment = '', String numberEmbodiment = ''}) {
  var row = List<dl.Primitive>.empty(growable: true);

  row.add(dl.Text(content: letter, embodiment: letterEmbodiment));
  for (int n = fromNo; n <= toNo; n++) {
    row.add(dl.Text(content: '$n', embodiment: numberEmbodiment));
  }
  return row;
}

dl.Primitive makeBingo() {
  var numEmb =
      'color:#FF000000, fontSize:5, fontWeight:bold, borderAll:2, paddingAll:5, width:40, height:40, horizontalTextAlignment:right, verticalTextAlignment:bottom';
  var modelRow = [
    dl.Text(
        embodiment:
            'color:#FFFFFFFF, fontSize:5, fontWeight:bold, borderAll:2, paddingAll:2, width:40, height:40'),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb),
    dl.Text(embodiment: numEmb)
  ];
  var rows = [
    makeRow('B', 1, 15,
        letterEmbodiment: 'backgroundColor:#FF0000FF',
        numberEmbodiment: 'backgroundColor:#00000000'),
    makeRow('I', 16, 30, letterEmbodiment: 'backgroundColor:#FFFF0000'),
    makeRow('N', 31, 45, letterEmbodiment: 'backgroundColor:#FF00FF00'),
    makeRow('G', 46, 60, letterEmbodiment: 'backgroundColor:#FFFFFF00'),
    makeRow('O', 61, 75, letterEmbodiment: 'backgroundColor:#FFFF00FF')
  ];
  return dl.Table(modelRow: modelRow, rows: rows);
}

*/

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

// makeRow creates a row for the bingo table.
func makeRow(letter string, fromNo, toNo int, letterEmbodiment, numberEmbodiment string) []pg.Primitive {
	row := []pg.Primitive{}

	row = append(row, pg.TextWith{
		Content:    letter,
		Embodiment: letterEmbodiment,
	}.Make())

	for n := fromNo; n <= toNo; n++ {
		row = append(row, pg.TextWith{
			Content:    fmt.Sprintf("%d", n),
			Embodiment: numberEmbodiment,
		}.Make())
	}
	return row
}

// makeBingo creates the bingo table.
func makeBingo() *pg.Table {
	numEmb := "color:#FF000000, fontSize:16, fontWeight:bold, borderAll:2, paddingAll:5, width:40, height:40, horizontalTextAlignment:center, verticalTextAlignment:middle"
	modelRow := []pg.Primitive{pg.TextWith{Embodiment: "color:#FFFFFFFF, fontSize:20, fontWeight:bold, borderAll:2, paddingAll:2, width:40, height:40"}.Make()}
	for i := 0; i < 15; i++ {
		modelRow = append(modelRow, pg.TextWith{Embodiment: numEmb}.Make())
	}
	rows := [][]pg.Primitive{
		makeRow("B", 1, 15, "backgroundColor:#FF0000FF", "backgroundColor:#00000000"),
		makeRow("I", 16, 30, "backgroundColor:#FFFF0000", ""),
		makeRow("N", 31, 45, "backgroundColor:#FF00FF00", ""),
		makeRow("G", 46, 60, "backgroundColor:#FFFFFF00", ""),
		makeRow("O", 61, 75, "backgroundColor:#FFFF00FF", ""),
	}
	return pg.TableWith{
		ModelRow: modelRow,
		Rows:     rows,
	}.Make()
}

func bingoDemo(pgui pg.ProntoGUI) {
	bingo := makeBingo()
	pgui.SetGUI(bingo)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
