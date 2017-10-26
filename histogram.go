// Command histogram reads non-negative float64 numbers on stdin and prints
// ASCII histogram of their distribution
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	if err := run(os.Stdout, os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(w io.Writer, r io.Reader) error {
	sc := bufio.NewScanner(r)
	var nums []float64
	var cnt int
	var max, sum float64
	min := math.Inf(1)
	for sc.Scan() {
		f, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil || f < 0 {
			continue
		}
		nums = append(nums, f)
		sum += f
		cnt++
		if f < min {
			min = f
		}
		if f > max {
			max = f
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	if len(nums) == 0 {
		return nil
	}
	sort.Float64s(nums)
	mean := sum / float64(len(nums))
	var med float64
	switch l := len(nums); {
	case l%2 == 0:
		med = (nums[l/2] + nums[l/2-1]) / 2
	default:
		med = nums[l/2]
	}
	var dev float64
	bkts := make(map[uint64]int)
	{
		var ss float64
		for _, f := range nums {
			ss += math.Pow(f-mean, 2)
			bkts[1<<uint64(math.Log2(f))] += 1
		}
		dev = math.Sqrt(ss / float64(len(nums)))
	}
	keys := make([]uint64, 0, len(bkts))
	for k := range bkts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	tw := tabwriter.NewWriter(w, 0, 8, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(tw, "min:%.2f mean:%.2f median:%.2f max:%.2f stddev:%.2f cnt:%v\n",
		min, mean, med, max, dev, cnt)
	fmt.Fprint(tw, "bkt\t"+strings.Repeat("-", 50)+"\tcnt\t\n")
	for _, k := range keys {
		hlen := bkts[k] * 100 / cnt / 2
		if hlen == 0 {
			hlen = 1
		}
		fmt.Fprintf(tw, "%d\t%s\t%d\t\n", k, strings.Repeat("*", hlen), bkts[k])
	}
	return tw.Flush()
}
