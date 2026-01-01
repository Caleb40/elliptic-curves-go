# Elliptic Curves over Finite Fields (Go)

A small, from scratch implementation of elliptic curves over finite fields in Go.
The goal of this project is to make the core ideas behind elliptic curve cryptography visible, not hidden behind
libraries.

This code focuses on *why* elliptic curves work in cryptography by explicitly implementing the mathematics they rely on.

## Motivation

Elliptic curve cryptography is often introduced as a black box: you call a library, pass a private key, and get a public
key. What’s usually missing is the structure that makes this possible.

At the heart of ECC are **finite fields** and the algebraic properties they provide. Without them, elliptic curves do
not form a usable group, and cryptographic assumptions like the discrete logarithm problem fall apart.

This repository exists to expose that structure.

## Finite Fields (𝔽ₚ)

A finite field 𝔽ₚ is the set:

$$
\mathbb{F}_p = {0, 1, 2, \dots, p - 1}
$$

where *p* is a prime number, and all arithmetic is performed modulo *p*.

Key properties:

* Addition, subtraction, multiplication, and division are closed
* Every non-zero element has a multiplicative inverse
* There are no rounding errors or approximations
* Arithmetic behaves predictably and deterministically

Division works because for any non-zero ( a \in \mathbb{F}_p ), there exists an inverse ( a^{-1} ) such that:

$$
a \cdot a^{-1} \equiv 1 \pmod{p}
$$

This is what makes slope calculations on elliptic curves well-defined.

## Elliptic Curves over 𝔽ₚ

An elliptic curve over a prime field is typically defined as:

$$
a \cdot a^{-1} \equiv 1 \pmod{p}
$$

Unlike curves over the real numbers, curves over finite fields are **discrete**. There are only finitely many valid
points, and all operations stay within the field.

The set of points on the curve, together with a special point at infinity, forms an **abelian group** under point
addition.

This implementation demonstrates:

* Point addition and point doubling
* How modular inverses replace division
* Why the group law holds only in a finite field
* How repeated addition leads to scalar multiplication

## What This Project Implements

* Finite field arithmetic (modular add, multiply, inverse)
* Elliptic curve point validation
* Point addition and doubling
* Scalar multiplication via repeated group operations

## What This Project Is Not

* Not constant-time
* Not resistant to side-channel attacks
* Not parameter-safe
* Not suitable for real cryptographic use
