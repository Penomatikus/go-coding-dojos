# Collection inspired slice

The implementation is based on the idea of implementing Java or JavaScript collection functions in Go.

The particularity of this implementation is:
- The underlying slice is not modified.
- Besides .New(), the average amount of memory allocated per operation in bytes (B) is 0
- Besides .New(), the average number of allocations per operation is 0

This is because, the implementation works with a caching system based on the indices of the underling slice. 
The cache it's self is reused for every collection function.  

### Example

| Slice / Indices  | 10 | 12 | 100 | -5 | 96 | 82 | 99 | 1 | 6 | B/op | allocs/op |
|------------------|----|----|-----|----|----|----|----|---|---|-----:|----------:|
| New()            | 0  | 1  | 2   | 3  | 4  | 5  | 6  | 7 | 8 |  96  |    1      |
| Filter(i%2 == 0) | 0  | 1  | 2   | _  | 4  | 5  | _  | 7 | 8 |   0  |    0      |
| Filter(i > 10)   | _  | 1  | 2   | _  | 4  | 5  | 6  | _ | _ |   0  |    0      |
| Take(4)          | _  | 1  | 2   | _  | 4  | 5  | _  | _ | _ |   0  |    0      |
| Skip(3)          | _  | _  | _   | _  | _  | 5  | _  | _ | _ |   0  |    0      |
| Last(i == 99)    | _  | _  | _   | _  | _  | _  | _  | _ | _ |   0  |    0      |
| First(i == -1)   | _  | _  | _   | _  | _  | _  | _  | _ | _ |   0  |    0      |

- Calling `.Collect()` after `.Skip(3)`, will take whats left in the cache, which is here "5" and use it as index on the slice passed to `.New()` resulting in "82".  
- Calling `.CollectIndices()` after `.Skip(3)`, will return "[5]".

### Error handling
There are no errors.

As you can see, `.Last(i == 99)` will result in an empty cache, the output of `.Collect()` would be an empty int slice:  
_This state is called `dry`._

Each collecting function on a dried collection will result in a "pass", meaning "nothing will happen".  
So, calling `.First(i == -1)` after `.Last(i == 99)` will not result in an error and wont hit the performance hard, since nothing will be done. 

### Implemented functions
The implementation features the following methods:

- `Filter`: The collection can be filtered with a filter function.
- `Take`: The first n elements are used for subsequent functions.
- `Skip`: The first n elements are not used for subsequent functions.
- `First`: The first element found that matches the filter function is returned.
- `Last`: The last element found that matches the filter function is returned.
- `Collect`: All elements determined based on the preceding collection functions are returned.
- `CollectIndices`: All indices of the underlying slices determined based on the preceding collection functions are returned.

### TODO
- implement a `ForEach`
- implement a `Map`
  - The map will use a nother cache (a map) to ensure the slice itself will not be mutated
  - Map will definitely impact B/op and allocs/op slightly
- implement a `ReverseFilter` starting the filter process from the last to the first
  - An implementation was done but I gave up at 80% because I had no time left. It was working besides some edge-cases.   


