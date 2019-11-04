package strcon

import "testing"

func TestSwapCase(t *testing.T) {
    s := SwapCase("goPher")
    if s != "GOpHER" {
        t.Error("Converted string is: " + s)
    }
}
