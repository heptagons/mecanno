# Meccano pentagons

To identify a pentagon we use two angles _A_ and _B_.
Some identities are solved for sqrt(5) values and will be used next.

| &nbsp;&nbsp;&nbsp;&nbsp;<img src="eqs/1.svg" width="225px"> |
|-------------|
| [equations 1](eqs/1.tex) |

## Pentagon of type 1

A pentagon of type is defined in next figure:

<img src="pentagon-type-1.svg">


We identify in this type, three rods (or sections of rods) 
**a**, **b** and **c** at fixed angles with integer sizes as for any meccanno figure.

We want to find a fourth rod **d** with integer size too to make the pentagon.

We start by looking the formulas to relate the rods:

| <img src="eqs/2.svg" width="500px"> |
|-------------|
| [equations 2](eqs/2.tex) |

We define two variables **m** and **n**. **m** is the sum of all terms multipled by **sqrt(5)**s while **n** is the sum of all the terms not multipled by **sqrt(5)**:

| <img src="eqs/3.svg" width="400px"> |
|-------------|
| [equations 3](eqs/3.tex) |

Simplifying, we get a value of rod **d<sup>2</sup>** in fuction of the rest of rods:

| <img src="eqs/4.svg" width="300px"> |
|-------------|
| [equations 4](eqs/4.tex) |

Now, we want rod **d<sup>2</sup>** to be as simple as possible so is good idea to set **m = 0**
which requires **ac = (a + c)b**. 

This way rod **d** is a simple function:

| <img src="eqs/5.svg" width="230px"> |
|-------------|
| [equations 5](eqs/5.tex) |

Withe equations 5, a program can iterate over the integer values
of **a**, **b** and **c** to discover **d** as integer too.

Next javascript program was run and found a single solution *`a = 12, b = 3, c = 4, d = 11`* after 5000 iterations. Scaled solutions are discared as are repetitions.
At this point

```
function meccano_pentagons_1(sols)
{
  this.find = (max)=> {
    for (let a=1; a < max; a++)
      for (let b=1; b <= max; b++)
        for (let c=0; c <= a; c++)
          if (a*c == (a + c)*b)
            mZero(a, b, c)
  }
  const mZero = (a, b, c)=> {
    const d = Math.sqrt(a*a + b*b + c*c - a*c)
    if (d > 0 && d % 1 === 0)
      dInteger(a, b, c, d)
  }
  const dInteger = (a, b, c, d) => {
    for (let i=0; i < sols.length; i++) {
      const s = sols[i]
      if (a % s.a == 0) {
        const f = a / s.a
        const bS = (b % s.b == 0) && b / s.b == f
        const cS = (c % s.c == 0) && c / s.c == f
        const dS = (d % s.d == 0) && d / s.d == f
        if (bS && cS && dS)
          return // scaled solution already
      }
    }
    sols.push({ a:a, b:b, c:c, d:d }) // solution!
  }
}

```



| The first solution of pentagon type 1 |
|---------------|
| <img src="pentagon-12a.svg"> |


## Pentagons type 2

