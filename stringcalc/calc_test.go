package stringcalc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddition(t *testing.T) {
	rand.Seed(time.Now().Unix())
	c := new(StringCalculator)
	Convey("given some arithmetic string", t, func() {
		Convey("should handle  addition", func() {
			So(c.Process("1 + 1"), ShouldEqual, 2)
			So(c.Process("1 + 999"), ShouldEqual, 1000)
			So(c.Process("5 + 4 + 3 + 2 + 1"), ShouldEqual, 15)
		})
		Convey("should handle subtraction", func() {
			So(c.Process("2 - 1"), ShouldEqual, 1)
			So(c.Process("100 - 50"), ShouldEqual, 50)
		})
		Convey("should handle multiplication", func() {
			So(c.Process("1 * 1"), ShouldEqual, 1)
			So(c.Process("0 * 1000"), ShouldEqual, 0)
			So(c.Process("9 * 9"), ShouldEqual, 81)
		})
		Convey("should handle division", func() {
			So(c.Process("1 / 1"), ShouldEqual, 1)
			So(c.Process("100 / 10"), ShouldEqual, 10)
			So(c.Process("81 / 9"), ShouldEqual, 9)
		})
		Convey("should handle combined operations", func() {
			So(c.Process("10 + 5 * 10"), ShouldEqual, 60)
			So(c.Process(" 9 / 3 - 1"), ShouldEqual, 2)
			So(c.Process("10 * 2 + 200 / 10"), ShouldEqual, 40)
			w := rand.Intn(100)
			x := rand.Intn(100)
			y := rand.Intn(100)
			z := rand.Intn(100)
			expected := (w * x) + (y / z)
			randomCalculation := fmt.Sprintf("%d * %d + %d / %d", w, x, y, z)
			So(c.Process(randomCalculation), ShouldEqual, expected)
		})
		Convey("should reject bad input", func() {
			So(c.Process("ABC - 50"), ShouldEqual, -1)
			So(c.Process("0 / 0"), ShouldEqual, -1)
			So(c.Process("-1 * 1"), ShouldEqual, -1)
		})
	})
}
