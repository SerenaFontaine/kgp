package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"
	"time"

	_ "image/png" // Register PNG decoder

	"github.com/SerenaFontaine/kgp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Clear screen and position cursor
	clearScreen()

	fmt.Println("Kitty Graphics Protocol - Go Bindings Demo")
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("Press Enter to proceed through each demo...")
	fmt.Println()

	// Demo 1: Basic formats - PNG
	clearScreen()
	fmt.Println("Demo 1: Image Formats - PNG format")
	redSquare := createColoredSquare(100, 100, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	pngData1, _ := kgp.ImageToPNG(redSquare)
	cmd := kgp.NewTransmitDisplay().
		ImageID(1).
		Format(kgp.FormatPNG).
		TransmitDirect(pngData1).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(cmd.Encode())
	fmt.Println("Red square displayed using PNG format")
	fmt.Println()
	waitForEnter(reader)

	// Demo 2: Basic formats - RGBA
	clearScreen()
	fmt.Println("Demo 2: Image Formats - RGBA format (32-bit)")
	blueData := kgp.SolidColorImage(100, 100, 0, 0, 255, 255)
	cmd2 := kgp.NewTransmitDisplay().
		ImageID(2).
		Format(kgp.FormatRGBA).
		Dimensions(100, 100).
		TransmitDirect(blueData).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(cmd2.Encode())
	fmt.Println("Blue square displayed using raw RGBA format")
	fmt.Println()
	waitForEnter(reader)

	// Demo 3: Basic formats - RGB
	clearScreen()
	fmt.Println("Demo 3: Image Formats - RGB format (24-bit)")
	greenSquare := createColoredSquare(100, 100, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	rgbData := kgp.ImageToRGB(greenSquare)
	cmd3 := kgp.NewTransmitDisplay().
		ImageID(3).
		Format(kgp.FormatRGB).
		Dimensions(100, 100).
		TransmitDirect(rgbData).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(cmd3.Encode())
	fmt.Println("Green square displayed using raw RGB format (no alpha channel)")
	fmt.Println()
	waitForEnter(reader)

	// Demo 4: Compression
	clearScreen()
	fmt.Println("Demo 4: Compression - ZLIB compressed RGBA data")
	largeOrangeData := kgp.SolidColorImage(200, 200, 255, 165, 0, 255)
	compressed, _ := kgp.CompressZlib(largeOrangeData)
	cmd4 := kgp.NewTransmitDisplay().
		ImageID(4).
		Format(kgp.FormatRGBA).
		Dimensions(200, 200).
		Compress().
		TransmitDirect(compressed).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(cmd4.Encode())
	fmt.Printf("Orange square (200x200) with ZLIB compression\n")
	fmt.Printf("Uncompressed: %d bytes, Compressed: %d bytes\n", len(largeOrangeData), len(compressed))
	fmt.Println()
	waitForEnter(reader)

	// Demo 5: Load and display an external PNG file
	clearScreen()
	fmt.Println("Demo 5: Loading external PNG file (sf.png)")
	if pngFile, err := os.Open("sf.png"); err == nil {
		defer pngFile.Close()
		if img, _, err := image.Decode(pngFile); err == nil {
			pngData5, _ := kgp.ImageToPNG(img)
			cmd5 := kgp.NewTransmitDisplay().
				ImageID(10).
				Format(kgp.FormatPNG).
				TransmitDirect(pngData5).
				ResponseSuppression(kgp.ResponseErrorsOnly).
				Build()
			fmt.Print(cmd5.Encode())
		} else {
			fmt.Printf("Error decoding sf.png: %v\n", err)
		}
	} else {
		fmt.Printf("sf.png not found (place a PNG file named 'sf.png' in the demo directory to see this demo)\n")
	}
	fmt.Println()
	waitForEnter(reader)

	// Demo 6: Cropping and scaling
	clearScreen()
	fmt.Println("Demo 6: Cropping and Scaling - Display part of an image")
	largeSquare := createGradientSquare(200, 200)
	largePNG, _ := kgp.ImageToPNG(largeSquare)
	cmd6 := kgp.NewTransmitDisplay().
		ImageID(20).
		Format(kgp.FormatPNG).
		TransmitDirect(largePNG).
		SourceRect(50, 50, 100, 100).
		DisplaySize(10, 10).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(cmd6.Encode())
	fmt.Println("Cropped to center 100x100 region of a 200x200 gradient")
	fmt.Println()
	waitForEnter(reader)

	// Demo 7: Multiple placements
	clearScreen()
	fmt.Println("Demo 7: Multiple Placements - One image, multiple displays")
	cyanSquare := createColoredSquare(80, 80, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	cyanPNG, _ := kgp.ImageToPNG(cyanSquare)

	// Transmit once with an ID
	transmit7 := kgp.NewTransmit().
		ImageID(100).
		Format(kgp.FormatPNG).
		TransmitDirect(cyanPNG).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(transmit7.Encode())

	// Create two placements
	place1 := kgp.NewPut(100).
		DisplaySize(10, 5).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(place1.Encode())

	place2 := kgp.NewPut(100).
		DisplaySize(5, 3).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(place2.Encode())
	fmt.Println()
	waitForEnter(reader)

	// Demo 8: Multiple placements with spacing
	clearScreen()
	fmt.Println("Demo 8: Multiple Placements with Different Sizes")
	fmt.Println()
	fmt.Println("Showing same image at different sizes:")
	fmt.Println()

	// Create a single gradient square
	gradientSquare := createGradientSquare(100, 100)
	gradientPNG, _ := kgp.ImageToPNG(gradientSquare)

	// Transmit once
	transmitGrad := kgp.NewTransmit().
		ImageID(160).
		Format(kgp.FormatPNG).
		TransmitDirect(gradientPNG).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(transmitGrad.Encode())

	// Place same image three times at different sizes, cursor will position them
	placeGrad1 := kgp.NewPut(160).
		DisplaySize(5, 5).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(placeGrad1.Encode())
	fmt.Print("  ")  // Add spacing

	placeGrad2 := kgp.NewPut(160).
		DisplaySize(8, 8).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(placeGrad2.Encode())
	fmt.Print("  ")  // Add spacing

	placeGrad3 := kgp.NewPut(160).
		DisplaySize(12, 12).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(placeGrad3.Encode())

	fmt.Println()
	fmt.Println()
	fmt.Println("Small (5x5 cells), Medium (8x8 cells), Large (12x12 cells)")
	fmt.Println()
	waitForEnter(reader)

	// Demo 9: Z-index layering
	clearScreen()
	fmt.Println("Demo 9: Z-Index Layering - Images above/below text")
	bgSquare := createColoredSquare(200, 100, color.RGBA{R: 255, G: 165, B: 0, A: 128})
	bgPNG, _ := kgp.ImageToPNG(bgSquare)
	bgCmdWithZ := kgp.NewTransmitDisplay().
		ImageID(200).
		Format(kgp.FormatPNG).
		TransmitDirect(bgPNG).
		ZIndex(-1).
		CursorMovement(false).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(bgCmdWithZ.Encode())
	fmt.Println("This orange square should appear behind the text!")
	fmt.Println()
	waitForEnter(reader)

	// Demo 10: Query support
	clearScreen()
	fmt.Println("Demo 10: Query Support - Check terminal capabilities")
	queryCmd := kgp.QuerySupport()
	fmt.Print(queryCmd.Encode())
	fmt.Println("(Check if your terminal responds with OK)")
	fmt.Println()
	waitForEnter(reader)

	// Demo 11: Animation - color-changing square moving left and right
	clearScreen()
	fmt.Println("Demo 11: Animation - Color-changing square with movement")
	fmt.Println("Creating animated sequence with 8 positions...")

	// Create base frame (red square) - transmit without displaying
	width, height := 80, 80
	redFrame := createColoredSquare(width, height, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	redPNG, _ := kgp.ImageToPNG(redFrame)

	baseAnimCmd := kgp.NewTransmit().
		ImageID(300).
		Format(kgp.FormatPNG).
		TransmitDirect(redPNG).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(baseAnimCmd.Encode())

	// Create additional frames with different colors as RGBA
	orangeData := kgp.SolidColorImage(width, height, 255, 165, 0, 255)
	yellowData := kgp.SolidColorImage(width, height, 255, 255, 0, 255)
	greenData := kgp.SolidColorImage(width, height, 0, 255, 0, 255)

	// Add frames with composition - use RGBA format with explicit dimensions
	frame1 := kgp.NewFrame(300).
		Format(kgp.FormatRGBA).
		Dimensions(width, height).
		FrameData(orangeData).
		Gap(100).
		Composition(kgp.CompositionReplace).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(frame1.Encode())

	frame2 := kgp.NewFrame(300).
		Format(kgp.FormatRGBA).
		Dimensions(width, height).
		FrameData(yellowData).
		Gap(100).
		Composition(kgp.CompositionReplace).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(frame2.Encode())

	frame3 := kgp.NewFrame(300).
		Format(kgp.FormatRGBA).
		Dimensions(width, height).
		FrameData(greenData).
		Gap(100).
		Composition(kgp.CompositionReplace).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(frame3.Encode())

	// Create a single placement and start the animation
	placement := kgp.NewPut(300).
		PlacementID(1).
		CellOffset(0, 0).
		CursorMovement(false).
		ResponseSuppression(kgp.ResponseErrorsOnly).
		Build()
	fmt.Print(placement.Encode())

	// Start the animation with infinite looping
	playCmd := kgp.PlayAnimationLoop(300)
	fmt.Print(playCmd.Encode())

	fmt.Println()
	fmt.Println("Animation playing with colors: Red -> Orange -> Yellow -> Green")
	fmt.Println()

	// Simulate movement by deleting and recreating the placement at different positions
	for i := 1; i < 8; i++ {
		time.Sleep(400 * time.Millisecond)

		// Delete the current placement
		delCmd := kgp.NewDelete(kgp.DeleteByPlacementID).
			ImageID(300).
			PlacementID(1).
			ResponseSuppression(kgp.ResponseErrorsOnly).
			Build()
		fmt.Print(delCmd.Encode())

		// Create new placement at next position (reusing placement ID 1)
		newPlacement := kgp.NewPut(300).
			PlacementID(1).
			CellOffset(i*5, 0).
			CursorMovement(false).
			ResponseSuppression(kgp.ResponseErrorsOnly).
			Build()
		fmt.Print(newPlacement.Encode())
	}

	// Stop the animation
	stopCmd := kgp.StopAnimation(300)
	fmt.Print(stopCmd.Encode())

	fmt.Println()
	fmt.Println("Animation sequence complete!")
	fmt.Println()
	fmt.Println()
	fmt.Println("Demo complete!")
	fmt.Println()
	fmt.Println("For more examples, see examples_test.go")
	fmt.Println("For full documentation, see README.md")
	waitForEnter(reader)
}

// waitForEnter waits for the user to press Enter
func waitForEnter(reader *bufio.Reader) {
	fmt.Print("Press Enter to continue...")
	reader.ReadString('\n')
	fmt.Println()
}

// clearScreen clears the terminal screen and moves cursor to top
func clearScreen() {
	// ANSI escape sequence to clear screen and move cursor to home
	fmt.Print("\x1b[2J\x1b[H")
	// Add some padding from the top
	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}

// createColoredSquare creates a solid colored square image
func createColoredSquare(width, height int, col color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, col)
		}
	}
	return img
}

// createGradientSquare creates a gradient square for cropping demo
func createGradientSquare(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Create a red-to-blue horizontal gradient
			r := uint8(255 * x / width)
			b := uint8(255 * (width - x) / width)
			img.Set(x, y, color.RGBA{R: r, G: 128, B: b, A: 255})
		}
	}
	return img
}
