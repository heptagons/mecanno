package nest

import (
	"fmt"
)

// aCmp return three valid options:
//	 0  if q numeric value equals r numeric value
//  +1  if q > r
//	-1  if q < r
func (qs *A32s) aCmp(q, r *A32) (int, error) {
	if q == nil || r == nil {
		return 0, nil // two nils are equal 0 = 0 ???
	}
	qMax, max, min := q.Deeper(r)

	cmp := func(mx, mn Z) int {
		if mx == mn {
			return 0
		} else if mx > mn {
			if qMax { return +1 }; return -1
		} else {
			if qMax { return -1 }; return +1
		}
	}

	w, U, u, B, b := max.LCM(min)
	BU := B*U
	bu := b*u
	BU_bu := BU - bu
	_ = w // simplified denominator is not used since we compare only numerators
	switch len(max.num) {
	case 1:
		// both number have deep level = 1
		//  B     b     BU     bu    
		// --- > --- : ---- > ----- : BU > bu
		//  A     a      w      w
		return cmp(BU, bu), nil

	case 3:
		C, D := max.cd()
		CU := C*U
		switch len(min.num) {
		case 1: // max = B/A min = b/a
			if B == 0 { // max = (C√D)/A, min = b/a
				// C√D    b    CU√D    bu
				// --- > --- : ---- > ---- : CU√D > bu : (CU)²D > (bu)²
				//  A     a      w      w
				return cmp(CU*CU*D, bu*bu), nil
			}
			// B + C√D    b    BU + CU√D    bu
			// ------- > --- : --------- > ---- : BU + CU√D > bu
			//    A       a        w         w
			// BU + CU√D > bu
			// CU√D > bu - BU
			return cmp(CU*CU*D, BU_bu*BU_bu), nil

		case 3: // min = (b+c√d)/a
			c, d := min.cd()
			cu := c*u
			if d == D { // Example: 1 + 2√3 > 4 + 5√3
				// B+C√D   b+c√D   BU+CU√D   bu+cu√D
				// ----- > ----- : ------- > ------- : BU + CU√D > bu + cu√D
				//   A       a        w         w
				// BU - bu > (cu - CU)√D
				// (BU - bu)² > (cu - CU)²D
				cu_CU := cu - CU
				return cmp(BU_bu*BU_bu, cu_CU*cu_CU*D), nil
			}
			if B == 0 {
				if b == 0 { // Example: 2√3 > 4√5
					// C√D   c√d   CU√D   cu√d
					// --- > --- : ---- > ---- : CU√D > cu√d
					//  A     a      w      w
					// (CU)²D > (cu)²d
					return cmp(CU*CU*D, cu*cu*d), nil

				} else { // Example: 2√3 > 1 + 4√5 
					// C√D   b+c√d   CU√D   bu+cu√d
					// --- > ----- : ---- > ------- : CU√D > bu + cu√d
					//  A       a      w        w
					// CU√D > bu + cu√d
					// x + CU√D = bu + cu√d
				}
			}
			if b == 0 { // example 1+2√3 > 4√5
				// B+C√D   c√d   BU+CU√D   cu√d
				// ----- > --- : ------- > ---- : BU+CU√D > cu√d
				//   A      a        w      w
				// BU+CU√D > cu√d


			}
fmt.Printf("33 max=%s min=%s\n", max, min)				
			// B + C√D    b + c√d   BU + CU√D    bu + cu√d
			// ------- > -------- : --------- > ---------- : BU + CU√D > bu + cu√d
			//    A         a           w           w
			// BU + CU√D > bu + cu√d
			// BU - bu > cu√d - CU√D
			// (BU - bu)² > (cu)²d + (CU)²D - 2cuCU√dD




		}

	}
	return 0, fmt.Errorf("Can't comp pair %s and %s", q, r)
}
