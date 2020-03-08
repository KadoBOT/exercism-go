package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Team struct {
	name string
	mp, w, d, l, p int
}

func (t *Team) Win() {
	t.w++
	t.mp++
	t.p += 3
}

func (t *Team) Tie() {
	t.d++
	t.mp++
	t.p++
}

func (t *Team) Lose() {
	t.l++
	t.mp++
}

func (t *Team) String() string {
	return fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", t.name, t.mp, t.w, t.d, t.l, t.p)
}

func rankByPoints(table map[string]*Team) PairList{
	pl := make(PairList, len(table))
	i := 0
	for k, v := range table {
		pl[i] = Pair{k, v.p}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key string
	Points int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Points < p[j].Points || (p[i].Points == p[j].Points && p[i].Key > p[j].Key) }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func Tally(r io.Reader, w io.Writer) error {
	table := make(map[string]*Team)

	p, _ := ioutil.ReadAll(r)
	s := strings.Split(string(p), "\n")

	createTeam := func(t string) *Team {
		if _, ok := table[t]; !ok {
			table[t] = &Team{ name:t }
		}

		return table[t]
	}

	for i, val := range s {
		g := strings.Split(val, ";")
		if s[i] == "" || s[i][0] == '#' {
			continue
		}

		if len(g) == 3 {
			t1 := createTeam(g[0])
			t2 := createTeam(g[1])
			r := g[2]

			switch r {
			case "win":
				t1.Win()
				t2.Lose()
			case "loss":
				t2.Win()
				t1.Lose()
			case "draw":
				t1.Tie()
				t2.Tie()
			default:
				return fmt.Errorf("invalid result: %s", r)
			}
		} else {
			return fmt.Errorf("invalid format: %s", s[i])
		}
	}

	rank := rankByPoints(table)

	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, t := range rank {
		team := table[t.Key]
		w.Write([]byte(team.String()))
	}

	return nil
}
