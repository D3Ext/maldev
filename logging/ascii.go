package logging

/*

References:
https://github.com/common-nighthawk/go-figure
https://github.com/moul/banner
https://manytools.org/hacker-tools/ascii-banner/

*/

import (
  figure "github.com/common-nighthawk/go-figure"
  "fmt"
)

func GetBanner(str_to_convert string) (string) { // Create ascii banner and return it as string
  var new_banner_str string
  new_banner := figure.NewFigure(str_to_convert, "", true)

	for _, banner_row := range new_banner.Slicify() {
    new_banner_str += fmt.Sprintf(banner_row + "\n")
  }

  return new_banner_str
}

func PrintBanner(str_to_convert string){ // Directly create banner and print to stdout
  new_banner := figure.NewFigure(str_to_convert, "", true)
  new_banner.Print()
}


