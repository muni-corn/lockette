package main

import (
    "fmt"
    "os"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "github.com/harrisonthorne/lockette/brightness"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("lockette needs an image to work")
        fmt.Println("use lockette like this:")
        fmt.Println("\tlockette /path/to/file.jpg")
        os.Exit(1);
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("sorry, lockette couldn't open your image file")
        os.Exit(1);
    }

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("sorry, lockette couldn't decode your image file")
        os.Exit(1);
    }

    classification, _ := brightness.GetImageBrightness(img, brightness.Center)

    fmt.Print(classification)
}
