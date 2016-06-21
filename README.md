# Generalized Suffix Tree
[![Build Status](https://api.travis-ci.org/dispareil/Generalized-Suffix-tree.svg?branch=master)](https://travis-ci.org/dispareil/Generalized-Suffix-tree)

# Summary

A package for creating "online generalized suffix trees" for fast retrieval.
It creates a tree based index out of a set of strings and then provides methods for
fast retrieval.

# How it works

Main methods are 'Add' and 'Search':

1. `Add` : adds the given string in the generalized suffix tree with a given key, this key is used for identifying the string uniquely
2. `Search` : searches for a input string in the generalized suffix tree and returns all the strings which contains this input string

In particular, after `Add(K, V)`, `Search(H)` will return a set containing `V` for any string `H` that is substring of `K`.

# Complexity

Time Complexity is *O(m)*


# Pending Items

1. `Remove` method
2. Test cases
