package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"fmt"
	"os"
	"flag"
)

type Layer_t struct {
	Layer     string   `yaml:"layer"`
	ImgBase64 string   `yaml:"imgBase64,omitempty"`
	XImage    *int     `yaml:"x_image,omitempty"`
	YImage    *int     `yaml:"y_image,omitempty"`
	Color     string   `yaml:"color,omitempty"`
	X         []int16  `yaml:"x_coords,omitempty"`
	Y         []int16  `yaml:"y_coords,omitempty"`
}



type Sprite_t struct {
	Animations []struct {
		Animation string      `yaml:"animation"`
		Frames    []struct {
			Frame  string     `yaml:"frame"`
			Delay  int        `yaml:"delay"`
			Layers []Layer_t  `yaml:"layers"`
		}                     `yaml:"frames"`
	}                         `yaml:"animations"`
}

func main() {

	var x int
	var y int
	var f float64
	var anim string
	var fram string
	var layr string
	var file string
	
	flag.IntVar(&x,       "x",     0, "x pad")
	flag.IntVar(&y,       "y",     0, "y pad")
	flag.Float64Var(&f,   "f",     0, "scale")
	flag.StringVar(&anim, "anim", "", "filter by  animation, if none, perform on all animations")
	flag.StringVar(&fram, "fram", "", "filter by frame name, if none, perform on all frames")
	flag.StringVar(&layr, "layr", "", "filter by layer name, if none, perform on all layer")
	flag.StringVar(&file, "file", "", "file to modify")
	flag.Parse()
	
	if file == "" {
		fmt.Printf("E: --file\n")
		flag.Usage()
		os.Exit(1)
	}
	
	if x == 0 && y == 0 && f == 0.0 {
		fmt.Printf("E: --x, --y, or --scale must be specified as a non-zero value\n")
		flag.Usage()
		os.Exit(1)
	}
	
	filename := filepath.Base(file)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("E: %v\n", err)
		os.Exit(1)
	}

	var s Sprite_t

	err = yaml.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Printf("E: %v\n", err)
		os.Exit(1)
	}
	
	printable_anim := "'" + anim + "'"
	if printable_anim == "''" { printable_anim = "ALL" }
	printable_fram := "'" + fram + "'"
	if printable_fram == "''" { printable_fram = "ALL" }
	printable_layr := "'" + layr + "'"
	if printable_layr == "''" { printable_layr = "ALL" }
	
	fmt.Printf("I: Shift %s animation(s) on %s frame(s) in %s layer(s) '%s' by %dx and %dy and scale %fx\n", printable_anim, printable_fram, printable_layr, file, x, y, f)
	
	
	for i := 0; i < len(s.Animations); i++ {
		
		if anim != "" { if s.Animations[i].Animation != anim { continue } }
		
		for j := 0; j < len(s.Animations[i].Frames); j++ {
			
			if fram != "" { if s.Animations[i].Frames[j].Frame != fram { continue } }
			
			for k := 0; k < len(s.Animations[i].Frames[j].Layers); k++ {
				
				if layr != "" { if s.Animations[i].Frames[j].Layers[k].Layer != layr { continue } }
				
				for l := 0; l < len(s.Animations[i].Frames[j].Layers[k].X); l++ {
					
					if ( x != 0 ) {
						s.Animations[i].Frames[j].Layers[k].X[l] += int16(x)
					}
					if ( y != 0 ) {
						s.Animations[i].Frames[j].Layers[k].Y[l] += int16(y)
					}
					if ( f != 0 ) {
						s.Animations[i].Frames[j].Layers[k].X[l] = int16(float64(s.Animations[i].Frames[j].Layers[k].X[l]) * f)
						s.Animations[i].Frames[j].Layers[k].Y[l] = int16(float64(s.Animations[i].Frames[j].Layers[k].Y[l]) * f)
					}
					
					s.Animations[i].Frames[j].Layers[k].XImage = nil
					s.Animations[i].Frames[j].Layers[k].YImage = nil
				}
			}
		}
	}
	
	out, err := yaml.Marshal(s)
	if err != nil {
		fmt.Printf("E: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(file, out, 0644)
	if err != nil {
		fmt.Println("error: %v", err)
	}

}
