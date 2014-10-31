/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package json

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/api/testapi"
	"github.com/openshift/origin/pkg/runtime"
	"github.com/openshift/origin/pkg/watch"
)

func TestDecoder(t *testing.T) {
	table := []watch.EventType{watch.Added, watch.Deleted, watch.Modified, watch.Error}

	for _, eventType := range table {
		out, in := io.Pipe()
		decoder := NewDecoder(out, testapi.Codec())

		expect := &api.Pod{ObjectMeta: api.ObjectMeta{Name: "foo"}}
		encoder := json.NewEncoder(in)
		go func() {
			data, err := testapi.Codec().Encode(expect)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if err := encoder.Encode(&watchEvent{eventType, runtime.RawExtension{json.RawMessage(data)}}); err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			in.Close()
		}()

		done := make(chan struct{})
		go func() {
			action, got, err := decoder.Decode()
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if e, a := eventType, action; e != a {
				t.Errorf("Expected %v, got %v", e, a)
			}
			if e, a := expect, got; !reflect.DeepEqual(e, a) {
				t.Errorf("Expected %v, got %v", e, a)
			}
			t.Logf("Exited read")
			close(done)
		}()
		<-done

		done = make(chan struct{})
		go func() {
			_, _, err := decoder.Decode()
			if err == nil {
				t.Errorf("Unexpected nil error")
			}
			close(done)
		}()
		<-done

		decoder.Close()
	}
}

func TestDecoder_SourceClose(t *testing.T) {
	out, in := io.Pipe()
	decoder := NewDecoder(out, testapi.Codec())

	done := make(chan struct{})

	go func() {
		_, _, err := decoder.Decode()
		if err == nil {
			t.Errorf("Unexpected nil error")
		}
		close(done)
	}()

	in.Close()

	select {
	case <-done:
		break
	case <-time.After(10 * time.Second):
		t.Error("Timeout")
	}
}
