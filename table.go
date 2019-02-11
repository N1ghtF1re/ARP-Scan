package main

import (
	"fmt"
	"io"
)

// CONSTANTS
const colSize = 20
const rowStr = "|%-20s|%-20s|%-20s|\n"

func atMiddle(s string) string{
	return fmt.Sprintf("%[1]*s", -colSize, fmt.Sprintf("%[1]*s", (colSize + len(s))/2, s))
}

func drawSplitter(writer io.Writer)  {
	for i := 0; i < colSize * 3 + 4; i++ {
		_, _ = fmt.Fprint(writer, "-")

	}
	_, _ = fmt.Fprintln(writer, "")
}

func drawHeader(writer io.Writer) {
	drawSplitter(writer)
	_, _ = fmt.Fprintf(writer, rowStr, atMiddle("IP"), atMiddle("MAC"), atMiddle("Name"))
	drawSplitter(writer)
}

func drawRow(writer io.Writer, node Node) {
	_, _ = fmt.Fprintf(writer, rowStr, atMiddle(node.ip), atMiddle(node.mac), atMiddle(node.name))
}