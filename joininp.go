package main

import (
    "fmt"
    "os"
    "math"
    "strconv"
    "github.com/yofu/st/stlib"
)

func main() {
    var err error
    var sfn, tfn string
    if len(os.Args)>= 2 {
        sfn = os.Args[1]
    } else {
        sfn = st.Input("Input File Name: ")
        sfn = st.Ce(sfn, ".inp")
    }
    if !st.FileExists(sfn) {
        st.Input(fmt.Sprintf("File doesn't exist: %s\n", sfn))
        os.Exit(1)
    }
    if len(os.Args)>= 3 {
        tfn = os.Args[2]
    } else {
        tfn = st.Input("Input File Name: ")
        tfn = st.Ce(tfn, ".inp")
    }
    if !st.FileExists(tfn) {
        st.Input(fmt.Sprintf("File doesn't exist: %s\n", tfn))
        os.Exit(1)
    }
    coord := make([]float64, 3)
    for i, j := range []string{"X", "Y", "Z"} {
        tmp := st.Input(fmt.Sprintf("Coord %s: ", j))
        val, err := strconv.ParseFloat(tmp, 64)
        if err != nil {
            coord[i] = 0.0
        } else {
            coord[i] = val
        }
    }
    tmp := st.Input("Rotate angle[deg]: ")
    val, err := strconv.ParseFloat(tmp, 64)
    var angle float64
    if err != nil {
        angle = 0.0
    } else {
        angle = val
    }
    fmt.Printf("Add %s at Coord (%.3f, %.3f, %.3f), Rotate Angle = %.3f\n", tfn, coord[0], coord[1], coord[2], angle)
    s := st.NewFrame()
    fmt.Printf("Reading %s...\n", sfn)
    err = s.ReadInp(sfn, []float64{0.0, 0.0, 0.0}, 0.0)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Printf("Reading %s...\n", tfn)
    angle *= math.Pi/180.0
    err = s.ReadInp(tfn, coord, angle)
    s.SetFocus()
    s.WriteInp("hogtxt.inp")
    st.Input("Succeeded")
}
