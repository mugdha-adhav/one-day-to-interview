# Nuances in GoLang

### Basic Constructs
* The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. Try avoinding sized or unsigned integer * unless specific need.
* Short variable assignment `:=` will not work outside function body.
* If initializer for a variable is present, then type can be ommited.
* Provision for named return statements, and return statement can be naked.
* `%T` prints type of the variable and `%v` for value of variable.
* Explicit conversion of variables using function.
* Numeric constants are high-precision values.

### Looping
* Go has only one looping construct, the `for` loop.
* The init and post statements are optional in a for loop are optional. (Semicolons not required)
* The most efficient infinite loop is empty for loop.
* The `range` form of the for loop iterates over a slice or map.
* In range, ff you only want the index, you can omit the second variable.

### Conditions
* `if` statement can start with a short statement to execute before the condition.

### Switch
* Runs the selected case only, not all the cases that follow. Thus no need of break.
* Cases need not be constants.
* Values involved need not be integers.
* Switch statement itself can have initializer.
* Switch without a condition is the same as switch true.

### Defer
* Executed at the end of the function.
* Calls arguments are evaluated immediately.
* Pushed onto stack and executed in LIFO order.

### Pointers
* Go has no pointer arithmatic.

### Structs
* Pointer to a struct can be directly accessed without `*`
* Assign particular values in a struct by specifying name explicitely.

### Arrays & Slices
* Arrays cannot be resized
* A slice does not store any data, it just describes a section of an underlying array.
* Changing the elements of a slice modifies the corresponding elements of its underlying array.
* Direct slice creation initially creates an array, then builds a slice that references it.
* In a slice, the default is zero for the low bound and the length of the slice for the high bound.
* The length of a slice `len` is the number of elements it contains.
* The capacity of a slice `cap` is the number of elements in the underlying array, counting from the first element in the slice.
* The zero value of a slice is nil.
* A nil slice has a length and capacity of 0 and has no underlying array.
* `make` can be used to dynamically sized zeroed array and returns a slice.
* A new array is allocated if slice `append` call exceeds the capacity of existing array.

### Maps
* The zero value of a map is nil. A nil map has no keys, nor can keys be added.
* Map element retrieval can return single or double value, where second value is bool to check if key is present.

### Functions
* Function values may be used as function arguments and return values.
