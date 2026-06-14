# Breaking Diffie Hellman

This problem will implement two discrete logarithm programs that can break Diffie-Hellman at various key strengths. The input to the program is decimal formated and provided as $(p, g, h)$. The 2 programs attempts to find an integer $x$ such that $g^x$ $mod$ $p= h$.

## Brute Force

A brute-force algorithm that simply tries every possibly $x$. On input a file containing decimal-formatted $( p, g, h )$, prints $x$ to standard output. The program can be executed as follows:

    dl-brute <filename for inputs>

## Baby Step Giant Step Algorithm

An efficient algorithm that is a meet-in-the-middle algorithm for computing the discrete logarithm or order of an element in a finite abelian group. On input a file containing decimal-formatted $( p, g, h )$, prints $x$ to standard output.The program can be executed as follows:

    dl-efficient <filename for inputs>

The program was performed on 20 - 40 bit numbers and it was successful. Pollard-rho algorithm can also be implemeneted for effeciently breaking Diffie-Hellman.

### Baby-Step Giant-Step Working Based on the Code

**Goal:** find `x` such that `G^x ≡ H (mod P)`.
**Input file contents:** `( 47, 5, 10 )`  →  `P = 47`, `G = 5`, `H = 10`.
**Answer we'll recover:** `x = 19` (since `5^19 mod 47 = 10`).

The algorithm rewrites the unknown as `x = q·M − r` and finds `q`, `r`
separately by meeting in the middle through a hash table.

---

#### Step 0 — Parse parameters (`getParameters`)

```
P = 47   G = 5   H = 10
```

#### Step 1 — Step size & iteration count (`getM` + `main`)

```
M     = floor(sqrt(P)) = floor(sqrt(47)) = 6
limit = floor(P / M) + 1 = floor(47 / 6) + 1 = 7 + 1 = 8
```

Both loops below run for index `0 .. limit-1` (i.e. `0 .. 7`).

#### Step 2 — Build the baby table (`buildBabyTable`)

Stores `baby[ H·G^r mod P ] = r`. The value starts at `H` and is advanced
by a single multiply per step (no per-step exponentiation):

```go
value := new(big.Int).Mod(H, P) // r = 0  → value = H
for r.SetInt64(0); r.Cmp(limit) < 0; r.Add(r, one) {
    baby[value.String()] = r.String()
    value.Mul(value, G); value.Mod(value, P)   // value ← value·G mod P
}
```

| `r` | `value` = H·G^r mod P  (key stored) | entry added | then `value ← value·G mod 47` |
|----:|:-----------------------------------:|:-----------:|:------------------------------|
| 0   | 10 | `baby[10] = 0` | 10·5 = 50 → 3  |
| 1   | 3  | `baby[3]  = 1` | 3·5  = 15       |
| 2   | 15 | `baby[15] = 2` | 15·5 = 75 → 28  |
| 3   | 28 | `baby[28] = 3` | 28·5 = 140 → 46 |
| 4   | 46 | `baby[46] = 4` | 46·5 = 230 → 42 |
| 5   | 42 | `baby[42] = 5` | 42·5 = 210 → 22 |
| 6   | 22 | `baby[22] = 6` | 22·5 = 110 → 16 |
| 7   | 16 | `baby[16] = 7` | (loop ends)     |

**Resulting baby table (key → r):**

```
{ 10→0, 3→1, 15→2, 28→3, 46→4, 42→5, 22→6, 16→7 }
```

#### Step 3 — Giant search (`babyStepGiantStep`)

Stride `gm = G^M mod P = 5^6 mod 47 = 21`. `giant` starts at `1` and is
multiplied by `gm` each step, so at iteration `q` it equals `(G^M)^q`:

```go
gm    := new(big.Int).Exp(G, M, P) // G^M = 21
giant := big.NewInt(1)             // (G^M)^0 = 1
for q.SetInt64(0); q.Cmp(limit) < 0; q.Add(q, one) {
    if rStr, ok := baby[giant.String()]; ok {   // is giant a key?
        // giant == H·G^r  ⟹  x = q·M − r
        x := new(big.Int).Mul(q, M); x.Sub(x, r)
        return x, true
    }
    giant.Mul(giant, gm); giant.Mod(giant, P)    // giant ← giant·G^M mod P
}
```

| `q` | `giant` = (G^M)^q mod 47 | key in baby table? | action |
|----:|:------------------------:|:-------------------|:-------|
| 0 | 1  | no            | `giant ← 1·21  = 21` |
| 1 | 21 | no            | `giant ← 21·21 = 18` |
| 2 | 18 | no            | `giant ← 18·21 = 2`  |
| 3 | 2  | no            | `giant ← 2·21  = 42` |
| 4 | 42 | **yes → r = 5** | recover `x` (below), `return` |

Note the `giant` column is exactly `(G^M)^q` — the loop keeps the invariant
`giant == (G^M)^q (mod P)` by advancing one multiply per step, never
recomputing from `q`.

#### Step 4 — Recover the exponent

Match at `q = 4`, `r = 5`:

```
x = q·M − r = 4·6 − 5 = 24 − 5 = 19
```

**Why it works:** the hit means `G^(q·M) ≡ H·G^r (mod P)`, so
`H ≡ G^(q·M − r) = G^x`.

**Check:** `5^19 mod 47 = 10 = H`  ✓

#### Program output

```
19
```

---

## Cost summary

| | brute force | this BSGS |
|---|---|---|
| Group operations | `O(P)` | `O(√P)` — here ~`2·8` vs up to `47` |
| Memory | `O(1)` | `O(√P)` for the baby table |