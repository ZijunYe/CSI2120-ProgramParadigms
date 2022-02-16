#Search Trees & Backtracking 

## Search trees represent queries
- question = root of the tree 
- Nodes = decisions and show which goals still need to be satisfied 
- transition = from one node to next (unification between a goal and a fact or rule)
- edges = a step in the proof(unifications

**Nodes in search tree**
- Goals are ordered from left to right 
- leaf nodes contain one or several goals are failure node(left most goal caused the failure)
- empty leaf nodes are success nodes


- Prolog builds the search tree from the question as a root node
- an empty leaf node =  proof or solution 
- an non-empty leaf = failure

## Backtracking 
- if there reached nodes that no any solution, prolog go back to the maker place to see if there have other answer 
	- the order based on the body of clause, and order of predicates 
- *How to determine which predicate is recursion*: same name, predicate is repeated 