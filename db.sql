CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    car_name CHAR(50) NOT NULL,
    day_rate DOUBLE PRECISION NOT NULL,
    month_rate DOUBLE PRECISION NOT NULL,
    image CHAR(256) NOT NULL
);

-- Tabel orders
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    car_id INT NOT NULL REFERENCES cars(id),
    order_date DATE NOT NULL,
    pickup_date DATE NOT NULL,
    dropoff_date DATE NOT NULL,
    pickup_location CHAR(50) NOT NULL,
    dropoff_location CHAR(50) NOT NULL
);
