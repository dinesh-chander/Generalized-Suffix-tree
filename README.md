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

## License

This Generalized Suffix Tree is released under the Apache License 2.0

   Copyright 2012 Dinesh Chander

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
