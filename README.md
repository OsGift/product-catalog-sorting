
# 📦 Product Catalog Sorting

> **A robust, extensible Go application for sorting product catalogs using flexible, runtime-selectable strategies.**  
> Designed for scalability, performance, and maintainability following industry-standard software engineering principles.

---

## 🚀 Table of Contents

- [💡 Overview](#💡-overview)
- [🏗 Project Structure](#🏗-project-structure)
- [🧩 Architecture & Design Decisions](#🧩-architecture--design-decisions)
- [✨ Key Features](#✨-key-features)
- [🛠 Installation & Running](#🛠-installation--running)
- [🔄 Extensibility Guide](#🔄-extensibility-guide)
- [⚡ Performance Considerations](#⚡-performance-considerations)
- [🧪 Testing](#🧪-testing)

---

## 💡 Overview

This Go application sorts product catalogs based on multiple strategies (e.g., price, sales/view ratio).  
The **Strategy Design Pattern** enables **runtime switching** between sorting algorithms, allowing **A/B testing** of different sortings without altering existing code.

### ⚙️ Engineering Highlights:
- 💎 **Clean Code**: Idiomatic Go, well-organized modules.
- 🚪 **Extensible**: Add new sorters without modifying existing code (**Open/Closed Principle**).
- 🏃 **Performant**: Leveraging Go’s efficient `sort.Slice` with `O(n log n)` complexity.
- 🧱 **Scalable Architecture**: Supports runtime strategy switching with a flexible context.

---

## 🏗 Project Structure

```
product-catalog-sorting/
├── README.md              # 📖 Documentation
├── go.mod                 # ⚡ Go module definition
├── main.go                # 🏃 Application entry point
│
├── models/                # 🧱 Domain models
│   └── product.go
│
├── sorters/               # 🧩 Sorting strategies
│   ├── price_sorter.go
│   ├── ratio_sorter.go
│   └── sorter.go
│
└── tests/                 # 🧪 Unit tests
    ├── price_sorter_test.go
    ├── product_test.go
    └── ratio_sorter_test.go
```

---

## 🧩 Architecture & Design Decisions

### 1️⃣ **Strategy Design Pattern (with Runtime Flexibility)**
- Sorting algorithms are encapsulated in independent **strategy structs**, implementing a common `Sorter` interface.
- A `SortContext` struct dynamically switches strategies at runtime, **without code modification**.

✅ **Why this matters:**  
- Allows teams to **add custom sorters** independently.  
- Adheres to **Open/Closed Principle**: **Open for extension, closed for modification**.  
- Promotes **single responsibility** per sorting strategy.

---

### 2️⃣ **Separation of Concerns**
- **`models/`**: Defines the domain entity (`Product`).  
- **`sorters/`**: Contains sorting logic (price and ratio strategies).  
- **`tests/`**: Houses **isolated unit tests** for each component.  
- **`main.go`**: Handles application flow and strategy selection.

✅ **Why this matters:**  
- Minimal coupling for **easy maintenance and testing**.  
- Supports **modular growth** as the project scales.

---

### 3️⃣ **Performance Optimized**
- **O(n log n)** sorting complexity via `sort.Slice`.  
- **Zero reflection** or generic complexity—**idiomatic Go** performance.  
- Handles **large datasets** efficiently.

---

## ✨ Key Features

- 🏷 **Sort by Price**
- 📊 **Sort by Sales/View Ratio**
- 🔄 **Dynamic Runtime Strategy Switching**
- 🧩 **Plug-and-Play Extensibility**: Add new sorters without altering existing logic.
- 🏃 **Production-Ready Structure**

---

## 🛠 Installation & Running

1️⃣ **Clone the repository:**
```bash
git clone https://github.com/yourusername/product-catalog-sorting.git
cd product-catalog-sorting
```

2️⃣ **Run the application:**
```bash
go run main.go
```

3️⃣ **Sample Output:**
```
🔽 Sorted by Price:
3: Coffee Table   | $10.00 | Ratio: 0.0521
1: Alabaster Table| $12.99 | Ratio: 0.0438
2: Zebra Table    | $44.49 | Ratio: 0.0918

🔽 Sorted by Sales/View Ratio:
2: Zebra Table    | $44.49 | Ratio: 0.0918
3: Coffee Table   | $10.00 | Ratio: 0.0521
1: Alabaster Table| $12.99 | Ratio: 0.0438
```

---

## 🔄 Extensibility Guide

💡 **Adding a New Sorting Strategy:**  
To add, for example, **sorting by product name**:

1️⃣ **Create a new sorter:**  
```go
type NameSorter struct{}

func (n NameSorter) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
}
```

2️⃣ **Use the new strategy at runtime:**
```go
context.SetStrategy(NameSorter{})
context.Execute(products)
```

✨ **No existing code modifications required.**

---

## ⚡ Performance Considerations

- 🔥 Leveraging **`sort.Slice`** ensures **O(n log n)** complexity.  
- 🏃 Zero unnecessary memory allocations.  
- 🧵 Optimized for concurrent access (if extended in future iterations).

---

## 🧪 Testing

✅ **Unit Tests:**  
- Full coverage for all sorting strategies and the product model.  
- Edge cases: empty slices, nil slices, zero views.  
- **Table-driven tests** for extensibility.

### 🏃 **Run all tests:**
```bash
go test ./tests/...
```

### ✅ **Sample Test Output:**
```
ok  	product-catalog-sorting/tests		0.002s
```
