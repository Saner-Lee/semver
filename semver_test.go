package semver

import "testing"

func TestNew(t *testing.T) {
	v, err := New("1.2.3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)

	v.MakeMajor()
	t.Logf("val : %s, excepted : %s", v, "2.0.0")

	v.MakeMinor()
	t.Logf("val : %s, excepted : %s", v, "2.1.0")

	v.MakePatch()
	t.Logf("val : %s, excepted : %s", v, "2.1.1")
}

func TestCompare(t *testing.T) {
	a, _ := New("1.12.3")
	b, _ := New("2.0.1")
	c, _ := New("1.12.3")
	d, _ := New("1.11.5")

	t.Logf("val : %v, excepted : %v", a.Lt(b), true)
	t.Logf("val : %v, excepted : %v", a.Eq(c), true)
	t.Logf("val : %v, excepted : %v", a.Gt(d), true)

	t.Log(a.Lt(b), true)
	t.Log(a.Eq(c), true)
	t.Log(a.Gt(d), true)

	t.Log(a.InDoubleCloseRange(d, b), true)
}
