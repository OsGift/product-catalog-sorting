📦 Product Catalog Sorting
An extensible, efficient, and production-ready product catalog sorting solution in Go.

🚀 Overview
This project implements a highly extensible, efficient, and scalable product catalog sorting solution using Go. The solution leverages the Strategy Design Pattern with a runtime-configurable strategy context, ensuring that teams can add new sorting behaviors without modifying existing code, adhering strictly to the Open/Closed Principle (OCP) of SOLID design.

🎯 Key Features
⚡ Flexible Sorting: Sort products by:
🏷 Price (ascending).
📈 Sales per View Ratio (descending).
🧩 Easily Extensible: Add new sorters without touching core code.
💡 Dynamic Runtime Switching: Switch sorting strategies at runtime.
🧱 Clean Architecture: Well-structured with clear separation of concerns.
🚀 Optimized Performance: Uses Go’s highly efficient sort.Slice (O(n log n) complexity).


💡 Why Strategy Pattern?
Open/Closed Principle: New sorting strategies can be added without modifying existing ones.
Encapsulation: Sorting algorithms are encapsulated in separate structs.
Flexibility: Sorting behavior can be switched at runtime.

🏃 Running the Code
git clone https://github.com/OsGift/product-catalog-sorting.git
cd product-catalog-sorting
go run main.go