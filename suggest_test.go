package youdao

import (
	"github.com/bytedance/mockey"
	"github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSuggest(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	mockey.PatchConvey("sim", t, func() {
		result, err := Suggest("sim", 10)
		t.Logf("result is %v, err is %v", result, err)
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(result), convey.ShouldEqual, 10)
	})
}
