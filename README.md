# Breaking Diffie Hellman

This problem will implement two discrete logarithm programs that can break Diffie-Hellman at various key strengths. The input to the program is decimal formated and provided as (p, g, h). The 2 programs attempts to find an integer x such that pow(g,x) = h.

## Brute Force

A brute-force algorithm that simply tries every possibly x. On input a file containing decimal-formatted ( p, g, h ), prints x to standard output. The program can be executed as follows:

    dl-brute <filename for inputs>

## Baby Step Giant Step Algorithm

An efficient algorithm that is a meet-in-the-middle algorithm for computing the discrete logarithm or order of an element in a finite abelian group. On input a file containing decimal-formatted ( p, g, h ), prints x to standard output.The program can be executed as follows:

    dl-efficient <filename for inputs>

The program was performed on 20 - 40 bit numbers and it was successful. Pollard-rho algorithm can also be implemeneted for effeciently breaking Diffie-Hellman.