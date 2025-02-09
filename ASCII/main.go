package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/qeesung/image2ascii/convert"
)

func main() {
	// Define CLI flags
	text := flag.String("text", "", "Text to convert to ASCII art")
	image := flag.String("image", "", "Path to image file to convert to ASCII")
	width := flag.Int("width", 70, "Width of ASCII output (default 80)")
	height := flag.Int("height", 30, "Height of ASCII output (default 40)")
	color := flag.Bool("color", false, "Enable colored ASCII output")
	foregroundColor := flag.Int("fg", 15, "Foreground color (0-255, default 15 - white)")
	backgroundColor := flag.Int("bg", 0, "Background color (0-255, default 0 - black)")
	flag.Parse()

	// Handle text-to-ASCII conversion
	if *text != "" {
		ascii := figure.NewFigure(*text, "standard", true)
		ascii.Print()
		return
	}

	// Handle image-to-ASCII conversion with size adjustment and color
	if *image != "" {
		converter := convert.NewImageConverter()
		options := convert.Options{
			FixedWidth:  *width,
			FixedHeight: *height,
			Colored:     *color, // Set to true if you want color output
		}

		// Generate ASCII art from the image
		asciiArt := converter.ImageFile2ASCIIString(*image, &options)

		// Apply foreground and background color if flags are set
		coloredArt := applyColors(asciiArt, *foregroundColor, *backgroundColor)
		fmt.Println(coloredArt)
		return
	}

	// If no flags were provided, prompt the user for text input
	fmt.Print("Enter text to convert to ASCII Art: ")
	reader := bufio.NewReader(os.Stdin)
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSpace(inputText)

	// Generate ASCII Art for the user input
	ascii := figure.NewFigure(inputText, "standard", true)
	ascii.Print()
}

// Function to apply foreground and background colors to ASCII art
func applyColors(asciiArt string, fgColor, bgColor int) string {
	// ANSI escape code for foreground color
	fgCode := fmt.Sprintf("\033[38;5;%dm", fgColor)
	// ANSI escape code for background color
	bgCode := fmt.Sprintf("\033[48;5;%dm", bgColor)
	// Reset color after the art
	resetCode := "\033[0m"

	// Apply the colors to the ASCII art
	return fgCode + bgCode + asciiArt + resetCode
}
