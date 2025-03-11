- `LRU cache` only the most recently used data is stored, to prevent long and frequent and expensive database calls again and again.

- `True LRU`
  - when a value arrives search for it , remove it and add new value to the front.
  - order of elements is maintained.
  - deletion happens at the tail, addition at head

- Parts of project
  - A linkedList using a `struct` Node.
  - A `queue` for FIFO.
  - A `unordered_map / hashmap` for storing K-V pairs.

- To insert
  - `newnode->left = head` of queue and `newnode->right = oldnode->left` of queue.