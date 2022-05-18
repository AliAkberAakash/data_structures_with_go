#!/bin/bash 

# define base path
BASE_PATH=github.com/AliAkberAakash/data_structures_with_go

# test linked list
go test $BASE_PATH/linkedlist/singlylist

# test binary search tree
go test $BASE_PATH/binarysearchtree/recurcive

# test hash map
go test $BASE_PATH/hashmap

# test trie
go test $BASE_PATH/trie