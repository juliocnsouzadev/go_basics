## **🚀 Go 1.21: The Future is Now! 🚀**

Hey there, code enthusiasts! 🎉 If you've been riding the Go train (or should I say, the Go rocket? 🚀), you're in for a treat. Go 1.21 just dropped, and it's packed with some seriously cool features. Let's dive into the ocean of Go and see what treasures we can find! 🌊

### **Some Cool Features in Go 1.21:**
- **New `math/bits` Functions:** Go 1.21 added `Add`, `Sub`, and `Mul` functions to the `math/bits` package. These functions perform addition, subtraction, and multiplication on pairs of unsigned integers with carry or borrow. Math wizards, rejoice! 🧙‍♂️
  
- **Enhanced `net` Package:** The `net` package now supports new methods for IP and IPNet, making IP address manipulation even easier. No more IP address headaches! 🎉

- **Improved Error Handling:** Go 1.21 introduced a new package called `errorsx`, which provides additional functions for inspecting errors. Because who likes cryptic error messages, right? 🤷‍♂️

- **And so much more!** From enhancements in the `runtime` package to new features in the `os` package, Go 1.21 is packed with goodies. Check out the [official documentation](https://tip.golang.org/doc/go1.21) for a deep dive.

### 🚀 **Hold Onto Your Keyboards, Coders! 🎸**

Introducing the *slices* package - the superhero of Go! 🦸‍♂️ Not just your average package, this bad boy is flexing with generic functions that vibe with slices of *any* element type. No more "one-size-fits-all" - it's time for "one-size-rocks-them-all!" 🎉

Whether you're slicing through arrays like a ninja 🥷 or just dabbling in the world of Go, the *slices* package is your new coding BFF! Dive in and let the slicing magic begin! 🌌

**The New `slices` Package: Binary Search Like a Pro!**
Ever felt like searching through slices was a bit... meh? Well, Go 1.21 heard your cries and introduced the brand new `slices` package. And guess what? It comes with a `BinarySearch` function! 🎉

Let's take a look at a example to see this bad boy in action:

```go
type Cat struct {
	Name string
	Age  int
}

func CompareCatFunc() func(someCat, otherCat Cat) int {
	return func(someCat, otherCat Cat) int {
		if cmpName := cmp.Compare(someCat.Name, otherCat.Name); cmpName != 0 {
			return cmpName
		}
		return cmp.Compare(someCat.Age, otherCat.Age)
	}
}

func BinarySearchFunc(list []Cat, wanted Cat) (int, bool) {
	sort.Slice(list, func(i, j int) bool {
		return CompareCatFunc()(list[i], list[j]) < 0
	})
	if index, ok := slices.BinarySearchFunc(list, wanted, CompareCatFunc()); ok {
		return index, true
	}
	return -1, false
}

func BinarySearch(list []int, wanted int) (int, bool) {
	sort.Ints(list)
	if index, ok := slices.BinarySearch(list, wanted); ok {
		return index, true
	}
	return -1, false
}

```

🎤 **Mic Drop Moment for Go Coders! 🎶**

Introducing *code* - Go's latest package that's more lit than a concert's spotlight! 🎵

1. **BinarySearch**: Ever tried finding that one song in a playlist? 🎧 This function is like your DJ for integers! Give it a sorted slice of numbers and your favorite integer, and watch it drop the beat at the right index. If your number ain't on the list, no worries - it'll give you a "-1" and a "better luck next time" (aka false). And guess what? It even sorts your tunes (I mean, integers) with `sort.Ints` before hitting the decks with `slices.BinarySearch`. 🎚️

2. **BinarySearchFunc**: Now, for the Cat lovers 🐱, this function is your feline matchmaker! Hand it a bunch of Cat structs and your purrfect Cat, and it'll find your feline's spot in the crowd. If your Cat's playing hide and seek? You'll get a "-1" and a "not today" (aka false). But here's the twist - it's got a custom DJ, the `CompareCatFunc`, spinning records based on Cat's Name and then Age. And with `sort.Slice` and `slices.BinarySearchFunc`, it's a party you won't forget!

So, whether you're jamming with integers or grooving with Cats, *code* is the package that'll keep the party going! 🎉

🚀 **Testing in Style** 🎸

Dive into the tests for `BinarySearch` and `BinarySearchFunc` from our main stage - the main package. Using Go's rockstar testing package, we're setting the stage on fire with some epic test cases!

```go
func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		list          []int
		wanted        int
		expectedIndex int
		expectedOk    bool
	}{
		{[]int{}, 1, -1, false},
		{[]int{1}, 1, 0, true},
		{[]int{1}, 2, -1, false},
		{[]int{1, 2, 3, 4, 5}, 3, 2, true},
		{[]int{1, 2, 3, 4, 5}, 6, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearch(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearch(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchFunc(t *testing.T) {
	testCases := []struct {
		list          []Cat
		wanted        Cat
		expectedIndex int
		expectedOk    bool
	}{
		{[]Cat{}, Cat{Name: "Fluffy", Age: 2}, -1, false},
		{[]Cat{{Name: "Fluffy", Age: 2}}, Cat{Name: "Fluffy", Age: 2}, 0, true},
		{[]Cat{{Name: "Fluffy", Age: 2}}, Cat{Name: "Mittens", Age: 3}, -1, false},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}}, Cat{Name: "Mittens", Age: 3}, 1, true},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}}, Cat{Name: "Whiskers", Age: 1}, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearchFunc(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearchFunc(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchUnsorted(t *testing.T) {
	testCases := []struct {
		list          []int
		wanted        int
		expectedIndex int
		expectedOk    bool
	}{
		{[]int{3, 2, 1}, 2, 1, true},
		{[]int{1, 3, 2}, 2, 1, true},
		{[]int{1, 2, 3}, 4, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearch(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearch(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchFuncUnsorted(t *testing.T) {
	testCases := []struct {
		list          []Cat
		wanted        Cat
		expectedIndex int
		expectedOk    bool
	}{
		{[]Cat{{Name: "Mittens", Age: 2}, {Name: "Fluffy", Age: 3}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Mittens", Age: 2}, 1, true},
		{[]Cat{{Name: "Fluffy", Age: 3}, {Name: "Fluffy", Age: 2}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Fluffy", Age: 2}, 0, true},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Tiger", Age: 1}, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearchFunc(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearchFunc(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}
```

🎤 **TestBinarySearch**: Jamming with integer slices, we're checking if our DJ, `BinarySearch`, can find the right beat. From solo tracks to full-blown albums, we're ensuring every integer finds its spotlight. If there's a mismatch, `t.Errorf` drops the mic with an error tune!

🐱 **TestBinarySearchFunc**: Cat lovers, this one's for you! Testing with a bunch of Cat structs, we're seeing if our feline DJ, `BinarySearchFunc`, can find the purrfect match. From solo cats to a clowder, every Cat's moment in the limelight is checked. Any mismatch? `t.Errorf` meows an error!

🔀 **TestBinarySearchUnsorted & TestBinarySearchFuncUnsorted**: Ever tried dancing to an unsorted playlist? These tests check if our DJs can handle the chaos of unsorted integers and Cats. Spoiler: They expect a "-1" and a "nope" (aka false) since binary search loves order!

So, whether it's integers or Cats, we're ensuring the party never stops with some solid testing! 🎉

---


### **Wrapping Up:**
Go 1.21 is not just an update; it's a leap into the future of programming. With its new features and improvements, coding in Go feels like a breeze. So, whether you're a seasoned Go developer or just starting out, Go 1.21 has something for everyone. Dive in, explore, and let the Go rocket take you to new coding heights! 🚀
