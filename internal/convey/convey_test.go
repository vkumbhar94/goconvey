package convey

import (
	"errors"
	"fmt"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestConvey(t *testing.T) {
	Convey("Convey Examples", t, func() {
		Convey("With Empty Context", func() {

		})
		Convey("With Context", func(c C) {
			c.Println("received context")
		})

		Convey("assertions", func() {
			Convey("nil or not nil", func() {
				Convey("should be nil", func() {
					var actual interface{}
					So(actual, ShouldBeNil)
				})
				Convey("should not be nil", func() {
					arr := make([]struct{}, 0)
					So(arr, ShouldNotBeNil)
				})
			})
			Convey("delta equality - nearness", func() {
				Convey("should almost equal", func() {
					// third param delta allowed
					So(5, ShouldAlmostEqual, 10, 6)
				})
				Convey("should not almost equal", func() {
					So(5, ShouldNotAlmostEqual, 10, 2)
				})
			})

			Convey("values falling within bounds", func() {
				Convey("should be between", func() {
					// two bounds - lower and upper
					So(5, ShouldBeBetween, 10, 3)
					// if order is given in reverse still it manages to assert
					So(5, ShouldBeBetween, 4, 6)
				})
				Convey("should not be between", func() {
					// two bounds - lower and upper
					So(50, ShouldNotBeBetween, 10, 3)
					// if order is given in reverse still it manages to assert
					So(55, ShouldNotBeBetween, 4, 6)
				})
			})
			Convey("given function should Panic", func() {
				So(func() {
					panic("panicked with msg")
				}, ShouldPanicWith, "panicked with msg")
			})
			Convey("checking errors is - wrapped errors", func() {
				constErr := errors.New("my error")

				actualErr := fmt.Errorf("failed with : %w", constErr)
				SoMsg(fmt.Sprintf("%s is not of type %s", actualErr, constErr), actualErr, ShouldWrap, constErr)
			})

		})
		Convey("teardown", func() {
			Convey("first test", func(c C) {
				c.Println("running first test")
			})
			Convey("second test", func(c C) {
				c.Println("running second test")
			})
			Reset(func() {
				Println("resetting")
			})
		})
		Convey("changing failure mode", func() {
			SetDefaultFailureMode(FailureContinues)
			defer SetDefaultFailureMode(FailureHalts)
			Convey("failing test", func() {
				SoMsg("1 is not equal 2", 1, ShouldEqual, 2)
			})
			SkipConvey("skip failing test for a while", func() {
				SoMsg("1 is not equal 2", 1, ShouldEqual, 2)
			})

			Convey("skip some assertions for a while", func() {
				So(1, ShouldEqual, 1)
				SkipSo(1, ShouldEqual, 2)
			})
		})
	})
}

type Abc struct {
}

func (a *Abc) Play() {

}

type IPlay interface{ Play() }
