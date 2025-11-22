package youdao

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestQuery(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	t.Run("word", func(t *testing.T) {
		resp, err := Query(context.Background(), "simple")
		if err != nil {
			t.Fatal(err)
		}
		bytes, _ := json.Marshal(resp)
		t.Log(string(bytes))
	})
	t.Run("phrase", func(t *testing.T) {
		resp, err := Query(context.Background(), "ready to go")
		if err != nil {
			t.Fatal(err)
		}
		bytes, _ := json.Marshal(resp)
		t.Log(string(bytes))
	})
	t.Run("Chinese", func(t *testing.T) {
		resp, err := Query(context.Background(), "简单的")
		if err != nil {
			t.Fatal(err)
		}
		bytes, _ := json.Marshal(resp)
		t.Log(string(bytes))
	})
}
