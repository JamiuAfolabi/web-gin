package dbrepo

import (
	"database/sql"
	"jamiuafolabi/web-service-gin/models"
	"log"
	"reflect"
)

// Generic function that executes SQL queries and maps results to any struct type
func GetResult[T any](db *sql.DB, query string) ([]T, error) {
	rows, err := db.Query(query)
	if err != nil {
		log.Println("query error:", err)
		return nil, err
	}
	defer rows.Close()

	var results []T

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t T
		dest := make([]interface{}, len(cols))

		val := reflect.ValueOf(&t).Elem()
		for i := range cols {
			if i < val.NumField() {
				dest[i] = val.Field(i).Addr().Interface()
			} else {
				// Extra column in query, not enough fields in struct
				var dummy interface{}
				dest[i] = &dummy
			}
		}

		if err := rows.Scan(dest...); err != nil {
			log.Println("scan error:", err)
			return nil, err
		}

		results = append(results, t)
	}

	if err = rows.Err(); err != nil {
		log.Println("rows error:", err)
		return nil, err
	}
	return results, nil
}

// Retrieves basic product information
func (m *DBRepo) GetAllProductInformation() ([]models.Product, error) {
	return GetResult[models.Product](m.DB, "SELECT id, name FROM product")
}

// Total sales by category
func (m *DBRepo) TotalSalesByCategory() ([]models.TotalSalesByCategory, error) {
	return GetResult[models.TotalSalesByCategory](m.DB, `SELECT 
								c.name AS category,
								CAST(SUM(st.total_amount) AS DOUBLE) AS total_sales
							FROM sales_transaction st
							JOIN product p ON st.product_id = p.id
							JOIN category c ON p.category_id = c.id
							GROUP BY c.name
							ORDER BY total_sales DESC`)

}

// Revenue trends by month
func (m *DBRepo) RevenueTrendMonthly() ([]models.RevenueByMonth, error) {
	return GetResult[models.RevenueByMonth](m.DB, `SELECT 
    DATE_TRUNC('month', transaction_date) AS month,
    CAST(SUM(total_amount) AS DOUBLE) AS revenue
FROM sales_transaction
GROUP BY month
ORDER BY month;`)

}

// Top 5 performing products

func (m *DBRepo) Top5PerformingProduct() ([]models.Top5Revenue, error) {
	return GetResult[models.Top5Revenue](m.DB, `SELECT 
    p.name AS product,
    CAST(SUM(st.total_amount) AS DOUBLE) AS total_revenue
FROM sales_transaction st
JOIN product p ON st.product_id = p.id
GROUP BY p.name
ORDER BY total_revenue DESC
LIMIT 5;`)

}

// Sales performance metrics
func (m *DBRepo) SalesPerformanceMetric() ([]models.SalesPerformanceMetrics, error) {
	return GetResult[models.SalesPerformanceMetrics](m.DB, `SELECT
    COUNT(*) AS total_orders,
    CAST(SUM(quantity) AS DOUBLE) AS total_items_sold,
    CAST(SUM(total_amount) AS DOUBLE) AS total_revenue,
    CAST(AVG(total_amount) AS DOUBLE) AS avg_order_value
FROM sales_transaction`)

}

// Category-wise sales distribution
func (m *DBRepo) SalesDistributionByCategory() ([]models.CategorySalesDistribution, error) {
	return GetResult[models.CategorySalesDistribution](m.DB, `SELECT 
    c.name AS category,
    CAST(SUM(st.total_amount) AS DOUBLE) AS category_sales,
    ROUND(CAST(SUM(st.total_amount) * 100.0 / SUM(SUM(st.total_amount)) OVER () AS DOUBLE), 2) AS percentage
FROM sales_transaction st
JOIN product p ON st.product_id = p.id
JOIN category c ON p.category_id = c.id
GROUP BY c.name
ORDER BY category_sales DESC`)

}

// Product performance within categories
func (m *DBRepo) ProductPerformanceByCategory() ([]models.ProductPerformanceByCategory, error) {
	return GetResult[models.ProductPerformanceByCategory](m.DB, `SELECT 
    c.name AS category,
    p.name AS product,
    CAST(SUM(st.total_amount) AS DOUBLE) AS revenue
FROM sales_transaction st
JOIN product p ON st.product_id = p.id
JOIN category c ON p.category_id = c.id
GROUP BY c.name, p.name
ORDER BY c.name, revenue DESC;`)

}

// Top 20 Customer purchase patterns
func (m *DBRepo) Top20CustomerPurchase() ([]models.CustomerPurchasePattern, error) {
	return GetResult[models.CustomerPurchasePattern](m.DB, `SELECT 
    c.name,
    c.email,
    COUNT(st.id) AS orders,
    CAST(SUM(st.total_amount) AS DOUBLE) AS total_spent
FROM sales_transaction st
JOIN customer c ON st.customer_id = c.id
GROUP BY c.name, c.email
ORDER BY orders DESC
LIMIT 20`)

}
