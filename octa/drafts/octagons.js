const SVG = function()
{
	const SVG = 'http://www.w3.org/2000/svg';

	const create = (elem)=> {
		return document.createElementNS(SVG, elem);
	}

	this.svg = (width, height, obj)=> {
	    var svg = create('svg');
		svg.setAttribute("xmlns", "http://www.w3.org/2000/svg");
		svg.setAttribute("width", width);
		svg.setAttribute("height", height);
	    for (prop in obj)
	        svg.setAttribute(prop, obj[prop]);
	    return svg;
	}

	this.g = (obj)=> {
	    var g = create('g');
	    for (prop in obj)
	        g.setAttribute(prop, obj[prop]);
	    return g;
	}

	this.line = (obj)=> {
	    var e = create('line');
	    for (prop in obj)
	        e.setAttribute(prop, obj[prop]);
	    return e;
	}

	this.path = (d, transform, obj)=> {
	    var e = create('path');
	    if (d)
	    	e.setAttribute("d", d);
	    if (transform)
	    	e.setAttribute("transform", transform);
	    if (obj) {
		    for (prop in obj)
		        e.setAttribute(prop, obj[prop]);
	    }
	    return e;
	}

	this.circle = (obj)=> {
	    var e = create('circle');
	    for (prop in obj)
	        e.setAttribute(prop, obj[prop]);
	    return e;
	}

	this.polygon = (points)=> {
		const e = create('polygon');
		e.setAttribute('points', points.map(p => `${p.x},${p.y}`).join(" "))
		return e;
	}

	const self = this;

	/*this.point = function(g, x, y) {
		this.circle = function(r) {
			const circle = self.circle({
				cx: x,
				cy: y,
				r:  r
			});
			g.appendChild(circle);
			return circle;
		}
	}*/
}


const Octagon = function()
{
	this.a1 = (a,b,c)=> ( a*a + b*b - c*c ) / (2*a);

	this.a2 = (a,b,c)=> ( a*a + c*c - b*b ) / (2*a);

	this.ay = (a,b,c)=> {
		const cb = c*c - b*b;
  		return Math.sqrt( a*a * (2*b*b + 2*c*c - a*a) - cb*cb ) / (2*a);
	}

	this.radical = (a,b,c,d)=> {
		const cb = c*c - b*b
		return (a*a + cb)                                      // 2*a*a2(a,b,c) 
			+ 2*a*d                                            // 2*a*d
			- Math.sqrt(a*a * (2*b*b + 2*c*c - a*a) - cb*cb ); // 2*a*ay(a,b,c)
	}

	this.integer = (a,b,c,d,e)=> {
		const cb = c*c - b*b
		return (a*a + cb)*(a*a + cb) / (4*a*a)              // a2(a,b,c)*a2(a,b,c) 
			+ (a*a*(2*b*b + 2*c*c - a*a) - cb*cb) / (4*a*a) // ay(a,b,c)*ay(a,b,c) 
	        + d*d + 4*e*e 
			+ (a*a + cb)*d / a;                             // 2*a2(a,b,c)*d
	}
}


const Build = function(id, points, colors, svg)
{
	const lines = (support)=> {
		const line = (p1, p2, stroke)=> {
			return Svg.line({ 
				x1:      p1.x,
				y1:      p1.y,
				x2:      p2.x,
				y2:      p2.y,
				opacity: 0.6,
				stroke:  stroke
			});
		}
		const l = [ 
			line(points.s0, points.s1, colors.a)  // e_a_d
		];
		if (support) {
			l.push(line(points.ab, points.bc, colors.b)); // b
			l.push(line(points.bc, points.ac, colors.c)); // c
			l.push(line(points.bc, points.ef, colors.f)); // f
		}
		return l;
	}

	const dist = (a, b)=> {
		return Math.sqrt((a.x - b.x)**2 + (a.y - b.y)**2);
	}

	this.data = (id)=> {
		var d = `<table>`;
		let cs = [ "a", "b", "c", "f" ]
		cs.forEach((c, i) => {
			d += `<tr>`;
			d += `<td style="background-color:` + colors[c] + `">&nbsp;&nbsp;&nbsp;</td>`;
			if (i == 0) {
				d += `<td>${dist(points.s0, points.s1)} (s0 - s1) border</td>`;
			} else if (i == 1) {
				d += `<td>${dist(points.ab, points.bc)} (ab - bc) inner</td>`;
			} else if (i == 2) {
				d += `<td>${dist(points.bc, points.ac)} (bc - ac) inner</td>`;
			} else if (i == 3) {
				d += `<td>${dist(points.bc, points.ef)} (bc - ef) inner</td>`;
			}
			d += `</tr>`;
		})
		d += `</table>`
		d += `<div>` + JSON.stringify(points) + `</div>`;
		document.getElementById(id).innerHTML = d;
	}

	this.octagon = (t)=> {
		const octand = (before, support)=> {
			const next = Svg.g({
				transform: `translate(${t.x} ${t.y}) rotate(45)`
			})
			lines(support).forEach(line => {
				next.appendChild(line);
			})
			before.appendChild(next);
			return next;
		}
		lines(true).forEach(line => { 
			g1.appendChild(line);
		})
		const g2 = octand(g1, true);
		const g3 = octand(g2, true);
		const g4 = octand(g3, true);
		const g5 = octand(g4, true);
		const g6 = octand(g5);
		const g7 = octand(g6);
		const g8 = octand(g7);
		return this;
	}

	const s = Svg.svg(svg.width, svg.height);
	document.getElementById(id).appendChild(s);

	const g1 = Svg.g({ 
		transform:`translate(${svg.tx} ${svg.ty}) scale(${svg.scale} -${svg.scale})`,
		"stroke-width":`${svg.strokeWidth}px`,
		"stroke-linecap": "round"
	});
	s.appendChild(g1);
}


