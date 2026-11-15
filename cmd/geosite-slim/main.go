package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jzksnsjswkw/geosite-cn-gfw/internal/router"
	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: geosite-slim <input.dat> <output.dat> <CODE1> [CODE2 ...]")
		os.Exit(1)
	}
	in, out := os.Args[1], os.Args[2]
	want := map[string]bool{}
	for _, c := range os.Args[3:] {
		want[strings.ToUpper(c)] = true
	}

	data, err := os.ReadFile(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", in, err)
		os.Exit(1)
	}

	var list router.GeoSiteList
	if err := proto.Unmarshal(data, &list); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "input: %d categories\n", len(list.Entry))

	var kept router.GeoSiteList
	total := 0
	for _, e := range list.Entry {
		code := e.GetCountryCode()
		if !want[strings.ToUpper(code)] {
			continue
		}
		kept.Entry = append(kept.Entry, e)
		total += len(e.GetDomain())
		fmt.Fprintf(os.Stderr, "  keep: %-10s %6d domains\n", code, len(e.GetDomain()))
	}

	if len(kept.Entry) == 0 {
		fmt.Fprintln(os.Stderr, "no categories matched")
		os.Exit(1)
	}

	buf, err := proto.Marshal(&kept)
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(out, buf, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write %s: %v\n", out, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "output: %d categories, %d domains, %d bytes\n",
		len(kept.Entry), total, len(buf))
}
