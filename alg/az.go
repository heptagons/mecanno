package alg

type AZ32 struct {
	o int32
	i []*AZ32
}

type AZ32s struct { // factory
	
}

func (a *AZ32s) F0(b int32) *AZ32 {
	return &AZ32{
		o: b,
	}
}

func (a *AZ32s) F1(c, d int32) *AZ32 {
	return &AZ32 {
		o: c,
		i: []*AZ32{
			a.F0(d),
		},
	}
}

func (a *AZ32s) F2(e, f, g, h int32) *AZ32 {
	return &AZ32 {
		o: e,
		i: []*AZ32{
			a.F0(f),
			a.F1(g, h),
		},
	}
}

func (a *AZ32s) F3(i, j, k, l, m, n, o, p int32) *AZ32 {
	return &AZ32 {
		o: i,
		i: []*AZ32{
			a.F0(j),
			a.F1(k, l),
			a.F2(m, n, o, p),
		},
	}
}

func (a *AZ32) String() string {
	return ""
}

