package dao

import (
	"golang.org/x/net/context"
	"github.com/coreos/etcd/client"
	"strings"
	"sort"
)

type PageOptions struct {
	Root     string
	PageSize int
	LastDir  string
	LastKey  string
}

func Paginate(ctx context.Context, api client.KeysAPI, opts *PageOptions) (*client.Response, error) {
	ev, err := api.Get(ctx, opts.Root, &client.GetOptions{Sort :true})
	if (err != nil && strings.Contains(err.Error(), "Key not found")) {
		return genEmptyPage(opts), nil
	}

	if (err != nil) {
		return nil, err
	}

	var dirs []string
	nodes := []*client.Node{}

	for _, d := range ev.Node.Nodes {
		if (d == nil || d.Dir == false) {
			continue
		}
		subDir := strings.TrimLeft(d.Key, opts.Root)
		dirs = append(dirs, subDir)
	}

	//not found
	if (len(dirs) == 0) {
		return genEmptyPage(opts), nil
	}

	nextDir := findNextSubDir(dirs, opts.LastDir, true)
	for (len(nodes) < opts.PageSize && len(nextDir) > 0) {
		opts.LastDir = nextDir
		nodes = queryNextDir(ctx, api, nodes, opts)
		nextDir = findNextSubDir(dirs, nextDir, false)
	}

	ret := genEmptyPage(opts)
	ret.Node.Nodes = nodes

	return ret, nil
}

func queryNextDir(ctx context.Context, api client.KeysAPI, nodesArray []*client.Node, opts *PageOptions) ([]*client.Node) {
	key := opts.Root + "/" + opts.LastDir
	ev, _ := api.Get(ctx, key, &client.GetOptions{Sort :true, Recursive:false})
	//fmt.Println("opts.Root:", opts.Root, ",opts.LastDir:", opts.LastDir, ",ev:", ev)
	sort.Sort(sort.Reverse(ev.Node.Nodes))
	for _, n := range ev.Node.Nodes {
		if (len(nodesArray) < opts.PageSize) {
			if len(opts.LastKey) == 0 || (len(opts.LastKey) > 0 && strings.Compare(n.Key, opts.LastKey) == -1) {
				nodesArray = append(nodesArray, n)
			}
		}
	}
	return nodesArray

}

func reverseInts(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func findNextSubDir(dirs []string, lastDir string, first bool) (string) {
	sort.Strings(dirs)
	dirs = reverseInts(dirs)

	if (lastDir == "") {
		return dirs[0]
	}

	for _, k := range dirs {
		if strings.Compare(k, lastDir) == -1 {
			return k
		}
		if (strings.Compare(k, lastDir) == 0 && first) {
			return k
		}
	}
	return ""
}

func genEmptyPage(opts *PageOptions) *client.Response {
	resp := &client.Response{
		Action :"page",
		Node :&client.Node{
			Key:opts.Root,
			Dir:true,
			Nodes:make(client.Nodes, 0),
		},
	}
	return resp
}
