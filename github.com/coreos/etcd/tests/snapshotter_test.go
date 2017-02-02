// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"os"
	"path"
	"testing"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/store"
	"fmt"
	"encoding/json"
)

var testSnap = &raftpb.Snapshot{
	Data: []byte("some snapshot"),
	Metadata: raftpb.SnapshotMetadata{
		ConfState: raftpb.ConfState{
			Nodes: []uint64{1, 2, 3},
		},
		Index: 1,
		Term:  1,
	},
}

func TestLoadNewestSnap(t *testing.T) {
	dir := path.Join(os.TempDir(), "snapshot")
	err := os.Mkdir(dir, 0700)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	testSnap, err = snap.Read(`C:\etcd-v2.0.5\default.etcd\member\snap\000000000000000e-000000000000ea66.snap`)
	if err != nil {
		fmt.Println("err:", err)
		t.Fatal(err)
		return
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("testSnap.Metadata.Index", ":", testSnap.Metadata.Index)
	fmt.Println("testSnap.Metadata.ConfState", ":", testSnap.Metadata.ConfState)
	fmt.Println("testSnap.Metadata.ConfState.Nodes", ":", testSnap.Metadata.ConfState.Nodes)
	fmt.Println("testSnap.Metadata.Term", ":", testSnap.Metadata.Term)
	fmt.Println("-----------------------------------------------")

	testStore(t, testSnap.Data)
}

func testStore(t *testing.T, data []byte) {
	st := store.New("/0", "/1")
	err := st.Recovery(data)
	if err != nil {
		t.Fatal(err)
		return
	}
	if ev, err := st.Get("/", true, false); err == nil {
		ret,_ :=json.Marshal(*ev)
		fmt.Println("event", "=", string(ret))
	}

}