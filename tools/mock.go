package dao

import (
	"golang.org/x/net/context"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/store"
	"github.com/pkg/errors"
)

type mockKeysAPI struct {
	st store.Store
}

var mapi = &mockKeysAPI{
	st : store.New(),
}

func GetMockEtcd() *mockKeysAPI {
	return mapi
}
func (this *mockKeysAPI ) Delete(ctx context.Context, key string, opts *client.DeleteOptions) (*client.Response, error) {
	ev, err := this.st.Delete(key, opts.Dir, opts.Recursive)
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), nil
}

// Get retrieves a set of Nodes from etcd
func (this *mockKeysAPI ) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	ev, err := this.st.Get(key, opts.Recursive, opts.Sort)
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), nil
	//return resp, nil
}
func (this *mockKeysAPI ) wrapRespone(ev *store.Event) (*client.Response) {
	resp := &client.Response{
		Action :ev.Action,
		Node :&client.Node{
			Key:ev.Node.Key,
			Dir:ev.Node.Dir,
			CreatedIndex:ev.Node.CreatedIndex,
			ModifiedIndex:ev.Node.ModifiedIndex,
			Nodes:make(client.Nodes, 0),
		},
	}

	if ev.Node.Value != nil {
		resp.Node.Value = *ev.Node.Value
	}
	if (len(ev.Node.Nodes) > 0) {
		resp.Node.Nodes = this.copyNodes(ev.Node.Nodes)
	}
	return resp
}

func (this *mockKeysAPI ) copyNodes(src store.NodeExterns) (client.Nodes) {
	dest := make(client.Nodes, len(src))
	i := 0
	for _, n := range src {
		dest[i] = &client.Node{
			Key:n.Key,
			Dir:n.Dir,
			CreatedIndex:n.CreatedIndex,
			ModifiedIndex:n.ModifiedIndex,
		}

		if (n.Value != nil) {
			dest[i].Value = *n.Value
		}
		i++
	}
	return dest
}

func (this *mockKeysAPI ) Set(ctx context.Context, key, value string, opts *client.SetOptions) (*client.Response, error) {
	ev, err := this.st.Set(key, opts.Dir, value, store.TTLOptionSet{})
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), nil
}

func (this *mockKeysAPI ) Create(ctx context.Context, key, value string) (*client.Response, error) {
	ev, err := this.st.Create(key, false, value, false, store.TTLOptionSet{})
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), nil
}
func (this *mockKeysAPI ) CreateInOrder(ctx context.Context, dir, value string, opts *client.CreateInOrderOptions) (*client.Response, error) {
	ev, err := this.st.Create(dir, false, value, false, store.TTLOptionSet{})
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), errors.New("not implement")
}
func (this *mockKeysAPI ) Update(ctx context.Context, key, value string) (*client.Response, error) {
	ev, err := this.st.Update(key, value, store.TTLOptionSet{})
	if (err != nil) {
		return nil, err
	}
	return this.wrapRespone(ev), nil
}
func (this *mockKeysAPI ) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	return nil
}
